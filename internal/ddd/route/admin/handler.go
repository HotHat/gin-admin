package admin

import (
	"context"
	"fmt"

	"github.com/HotHat/gin-admin/v10/internal/config"
	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/service"
	"github.com/HotHat/gin-admin/v10/pkg/middleware"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AuthService *service.AuthService
	Casbinx     *Casbinx
}

func InitAdminHandler(ctx context.Context, authService *service.AuthService, cabin *Casbinx) (*AdminHandler, func(), error) {
	handler := &AdminHandler{
		AuthService: authService,
		Casbinx:     cabin,
	}

	err := handler.Register(ctx)
	if err != nil {
		return nil, nil, err
	}

	return handler, func() {
		err := handler.Release(ctx)
		if err != nil {
			return
		}
	}, nil
}

func (a *AdminHandler) Register(ctx context.Context) error {
	return a.Casbinx.Load(ctx)
}

func (a *AdminHandler) Release(ctx context.Context) error {
	return a.Casbinx.Release(ctx)
}

func (a *AdminHandler) GetHandlers() []gin.HandlerFunc {
	allowedPrefixes := config.C.Middleware.Auth.AllowedPrefixes
	return []gin.HandlerFunc{
		// must before Casbin
		middleware.AuthWithConfig(middleware.AuthConfig{
			AllowedPathPrefixes: allowedPrefixes,
			SkippedPathPrefixes: config.C.Middleware.Auth.SkippedPathPrefixes,
			ParseUserID:         a.parseUserId,
			RootID:              config.C.General.Root.ID,
		}),
		middleware.CasbinWithConfig(middleware.CasbinConfig{
			AllowedPathPrefixes: allowedPrefixes,
			SkippedPathPrefixes: config.C.Middleware.Casbin.SkippedPathPrefixes,
			Skipper: func(c *gin.Context) bool {
				if config.C.Middleware.Casbin.Disable ||
					util.FromIsRootUser(c.Request.Context()) {
					return true
				}
				return false
			},
			GetEnforcer: func(c *gin.Context) *casbin.Enforcer {
				return a.Casbinx.GetEnforcer()
			},
			GetSubjects: func(c *gin.Context) []string {
				ids := util.FromUserCache(c.Request.Context()).RoleIDs
				fmt.Println("user role ids", ids)
				return ids
			},
		}),
	}
}

func (a *AdminHandler) parseUserId(c *gin.Context) (string, error) {
	userID, err := a.AuthService.ParseUserID(c)
	if err != nil {
		return "", err
	}
	userIDStr := comm.IDToStr(userID)
	return userIDStr, nil
}
