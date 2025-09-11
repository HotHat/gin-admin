package service

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/HotHat/gin-admin/v10/internal/config"
	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/entity"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/repo"
	"github.com/HotHat/gin-admin/v10/pkg/cachex"
	"github.com/HotHat/gin-admin/v10/pkg/crypto/hash"
	resp "github.com/HotHat/gin-admin/v10/pkg/errors"
	"github.com/HotHat/gin-admin/v10/pkg/jwtx"
	"github.com/HotHat/gin-admin/v10/pkg/logging"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"github.com/LyricTian/captcha"
	"github.com/LyricTian/captcha/store"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type AuthService struct {
	Cache        cachex.Cacher
	Auth         jwtx.Auther
	UserRepo     *repo.UserRepo
	UserRoleRepo *repo.UserRoleRepo
	UserService  *UserService
}

func (a *AuthService) ParseUserID(c *gin.Context) (comm.ID, error) {
	rootIDStr := config.C.General.Root.ID
	rootID, err := comm.StrToID(rootIDStr)
	if err != nil {
		return 0, err
	}

	if config.C.Middleware.Auth.Disable {
		return rootID, nil
	}

	invalidToken := resp.Unauthorized(config.ErrInvalidTokenID, "Invalid access token")
	token := util.GetToken(c)
	if token == "" {
		return 0, invalidToken
	}

	ctx := c.Request.Context()
	ctx = util.NewUserToken(ctx, token)

	userIDStr, err := a.Auth.ParseSubject(ctx, token)
	if err != nil {
		return 0, err
	}
	userID, err := comm.StrToID(userIDStr)
	if err != nil {
		if errors.Is(err, jwtx.ErrInvalidToken) {
			return 0, invalidToken
		}
		return 0, err
	} else if userID == rootID {
		c.Request = c.Request.WithContext(util.NewIsRootUser(ctx))
		return userID, nil
	}

	userCacheVal, ok, err := a.Cache.Get(ctx, config.CacheNSForUser, userIDStr)
	if err != nil {
		return 0, err
	} else if ok {
		userCache := util.ParseUserCache(userCacheVal)
		c.Request = c.Request.WithContext(util.NewUserCache(ctx, userCache))
		return userID, nil
	}

	// Check user status, if not activated, force to logout
	user, err := a.UserRepo.Get(ctx, userID, dto.UserQueryOptions{
		QueryOptions: util.QueryOptions{SelectFields: []string{"status"}},
	})
	if err != nil {
		return 0, err
	} else if user == nil || user.Status != entity.UserStatusActivated {
		return 0, invalidToken
	}

	roleIDs, err := a.UserService.GetRoleIDs(ctx, userID)
	if err != nil {
		return 0, err
	}

	userCache := util.UserCache{
		RoleIDs: comm.IDArrToStr(roleIDs),
	}
	err = a.Cache.Set(ctx, config.CacheNSForUser, userIDStr, userCache.String())
	if err != nil {
		return 0, err
	}

	c.Request = c.Request.WithContext(util.NewUserCache(ctx, userCache))
	return userID, nil
}

func (a *AuthService) setStore(c context.Context) {
	redisStore := store.NewRedisStore(&redis.Options{
		Addr:     config.C.Util.Captcha.Redis.Addr,
		Username: config.C.Util.Captcha.Redis.Username,
		Password: config.C.Util.Captcha.Redis.Password,
		DB:       config.C.Util.Captcha.Redis.DB,
	}, 10*time.Minute, nil, config.C.Util.Captcha.Redis.KeyPrefix)

	captcha.SetCustomStore(redisStore)
}

// GetCaptcha This function generates a new captcha ID and returns it as a `schema.Captcha` struct. The length of
// the captcha is determined by the `config.C.Util.Captcha.Length` configuration value.
func (a *AuthService) GetCaptcha(ctx context.Context) (*dto.Captcha, error) {
	a.setStore(ctx)

	return &dto.Captcha{
		CaptchaID: captcha.NewLen(config.C.Util.Captcha.Length),
	}, nil
}

// ResponseCaptcha Response captcha image
func (a *AuthService) ResponseCaptcha(ctx context.Context, w http.ResponseWriter, id string, reload bool) error {
	a.setStore(ctx)

	if reload && !captcha.Reload(id) {
		return resp.NotFound("", "Captcha id not found")
	}

	err := captcha.WriteImage(w, id, config.C.Util.Captcha.Width, config.C.Util.Captcha.Height)
	if err != nil {
		if errors.Is(err, captcha.ErrNotFound) {
			return resp.NotFound("", "Captcha id not found")
		}
		return err
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "image/png")
	return nil
}

func (a *AuthService) genUserToken(ctx context.Context, userID string) (*dto.LoginToken, error) {
	token, err := a.Auth.GenerateToken(ctx, userID)
	if err != nil {
		return nil, err
	}

	tokenBuf, err := token.EncodeToJSON()
	if err != nil {
		return nil, err
	}
	logging.Context(ctx).Info("Generate user token", zap.Any("token", string(tokenBuf)))

	return &dto.LoginToken{
		AccessToken: token.GetAccessToken(),
		TokenType:   token.GetTokenType(),
		ExpiresAt:   token.GetExpiresAt(),
	}, nil
}

func (a *AuthService) Login(ctx context.Context, formItem *dto.LoginForm) (*dto.LoginToken, error) {
	// verify captcha
	if !captcha.VerifyString(formItem.CaptchaID, formItem.CaptchaCode) {
		return nil, resp.BadRequest(config.ErrInvalidCaptchaID, "Incorrect captcha")
	}

	ctx = logging.NewTag(ctx, logging.TagKeyLogin)

	// login by root
	if formItem.Username == config.C.General.Root.Username {
		if formItem.Password != config.C.General.Root.Password {
			return nil, resp.BadRequest(config.ErrInvalidUsernameOrPassword, "Incorrect username or password")
		}

		userID := config.C.General.Root.ID
		ctx = logging.NewUserID(ctx, userID)
		logging.Context(ctx).Info("AuthService by root")
		return a.genUserToken(ctx, userID)
	}

	// get user info
	user, err := a.UserRepo.GetByUsername(ctx, formItem.Username, dto.UserQueryOptions{
		QueryOptions: util.QueryOptions{
			SelectFields: []string{"id", "password", "status"},
		},
	})
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, resp.BadRequest(config.ErrInvalidUsernameOrPassword, "Incorrect username or password")
	} else if user.Status != entity.UserStatusActivated {
		return nil, resp.BadRequest("", "UserService status is not activated, please contact the administrator")
	}

	// check password
	if err := hash.CompareHashAndPassword(user.Password, formItem.Password); err != nil {
		return nil, resp.BadRequest(config.ErrInvalidUsernameOrPassword, "Incorrect username or password")
	}

	userID := user.ID
	userIDStr := comm.IDToStr(userID)
	ctx = logging.NewUserID(ctx, userIDStr)

	// set user cache with role ids
	roleIDs, err := a.UserService.GetRoleIDs(ctx, userID)
	if err != nil {
		return nil, err
	}
	roleIDArr := comm.IDArrToStr(roleIDs)
	userCache := util.UserCache{RoleIDs: roleIDArr}
	err = a.Cache.Set(ctx, config.CacheNSForUser, userIDStr, userCache.String(),
		time.Duration(config.C.Dictionary.UserCacheExp)*time.Hour)
	if err != nil {
		logging.Context(ctx).Error("Failed to set cache", zap.Error(err))
	}
	logging.Context(ctx).Info("AuthService success", zap.String("username", formItem.Username))

	// generate token
	return a.genUserToken(ctx, userIDStr)
}

func (a *AuthService) RefreshToken(ctx context.Context) (*dto.LoginToken, error) {
	userIDStr := util.FromUserID(ctx)
	userID, err := comm.StrToID(userIDStr)
	if err != nil {
		return nil, err
	}

	user, err := a.UserRepo.Get(ctx, userID, dto.UserQueryOptions{
		QueryOptions: util.QueryOptions{
			SelectFields: []string{"status"},
		},
	})
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, resp.BadRequest("", "Incorrect user")
	} else if user.Status != entity.UserStatusActivated {
		return nil, resp.BadRequest("", "UserService status is not activated, please contact the administrator")
	}

	return a.genUserToken(ctx, userIDStr)
}

func (a *AuthService) Logout(ctx context.Context) error {
	userToken := util.FromUserToken(ctx)
	if userToken == "" {
		return nil
	}

	ctx = logging.NewTag(ctx, logging.TagKeyLogout)
	if err := a.Auth.DestroyToken(ctx, userToken); err != nil {
		return err
	}

	userID := util.FromUserID(ctx)
	err := a.Cache.Delete(ctx, config.CacheNSForUser, userID)
	if err != nil {
		logging.Context(ctx).Error("Failed to delete user cache", zap.Error(err))
	}
	logging.Context(ctx).Info("Logout success")

	return nil
}
