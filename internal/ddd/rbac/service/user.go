package service

import (
	"context"
	"sort"
	"time"

	"github.com/HotHat/gin-admin/v10/internal/config"
	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/entity"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/repo"
	"github.com/HotHat/gin-admin/v10/pkg/cachex"
	"github.com/HotHat/gin-admin/v10/pkg/crypto/hash"
	"github.com/HotHat/gin-admin/v10/pkg/errors"
	"github.com/HotHat/gin-admin/v10/pkg/util"
)

// UserService management for RBAC
type UserService struct {
	Cache        cachex.Cacher
	Trans        *util.Trans
	UserRepo     *repo.UserRepo
	UserRoleRepo *repo.UserRoleRepo
	MenuRepo     *repo.MenuRepo
}

// Query users from the data access object based on the provided parameters and options.
func (a *UserService) Query(ctx context.Context, params dto.UserQueryParam) (*dto.UserQueryResult, error) {
	params.Pagination = true

	result, err := a.UserRepo.Query(ctx, params, dto.UserQueryOptions{
		QueryOptions: util.QueryOptions{
			OrderFields: []util.OrderByParam{
				{Field: "created_at", Direction: util.DESC},
			},
			OmitFields: []string{"password"},
		},
	})
	if err != nil {
		return nil, err
	}

	if userIDs := result.Data.ToIDs(); len(userIDs) > 0 {
		userRoleResult, err := a.UserRoleRepo.Query(ctx, dto.UserRoleQueryParam{
			InUserIDs: userIDs,
		}, dto.UserRoleQueryOptions{
			JoinRole: true,
		})
		if err != nil {
			return nil, err
		}
		userRolesMap := userRoleResult.Data.ToUserIDMap()
		for _, user := range result.Data {
			user.Roles = userRolesMap[user.ID]
		}
	}

	return result, nil
}

// Get the specified user from the data access object.
func (a *UserService) Get(ctx context.Context, id comm.ID) (*entity.User, error) {
	user, err := a.UserRepo.Get(ctx, id, dto.UserQueryOptions{
		QueryOptions: util.QueryOptions{
			OmitFields: []string{"password"},
		},
	})
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, errors.NotFound("", "UserService not found")
	}

	userRoleResult, err := a.UserRoleRepo.Query(ctx, dto.UserRoleQueryParam{
		UserID: id,
	})
	if err != nil {
		return nil, err
	}
	user.Roles = userRoleResult.Data

	return user, nil
}

// Create a new user in the data access object.
func (a *UserService) Create(ctx context.Context, formItem *dto.UserForm) (*entity.User, error) {
	existsUsername, err := a.UserRepo.ExistsUsername(ctx, formItem.Username)
	if err != nil {
		return nil, err
	} else if existsUsername {
		return nil, errors.BadRequest("", "Username already exists")
	}

	user := &entity.User{
		//ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}

	if formItem.Password == "" {
		formItem.Password = config.C.General.DefaultLoginPwd
	}

	if err := formItem.FillTo(user); err != nil {
		return nil, err
	}

	err = a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.UserRepo.Create(ctx, user); err != nil {
			return err
		}

		for _, userRole := range formItem.Roles {
			//userRole.ID = util.NewXID()
			userRole.UserID = user.ID
			userRole.CreatedAt = time.Now()
			if err := a.UserRoleRepo.Create(ctx, userRole); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	user.Roles = formItem.Roles

	return user, nil
}

// Update the specified user in the data access object.
func (a *UserService) Update(ctx context.Context, id comm.ID, formItem *dto.UserForm) error {
	user, err := a.UserRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if user == nil {
		return errors.NotFound("", "UserService not found")
	} else if user.Username != formItem.Username {
		existsUsername, err := a.UserRepo.ExistsUsername(ctx, formItem.Username)
		if err != nil {
			return err
		} else if existsUsername {
			return errors.BadRequest("", "Username already exists")
		}
	}

	if err := formItem.FillTo(user); err != nil {
		return err
	}
	user.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.UserRepo.Update(ctx, user); err != nil {
			return err
		}

		if err := a.UserRoleRepo.DeleteByUserID(ctx, id); err != nil {
			return err
		}
		for _, userRole := range formItem.Roles {
			//if userRole.ID == 0 {
			//userRole.ID = util.NewXID()
			//}
			userRole.UserID = user.ID
			if userRole.CreatedAt.IsZero() {
				userRole.CreatedAt = time.Now()
			}
			userRole.UpdatedAt = time.Now()
			if err := a.UserRoleRepo.Create(ctx, userRole); err != nil {
				return err
			}
		}

		return a.Cache.Delete(ctx, config.CacheNSForUser, comm.IDToStr(id))
	})
}

// Delete the specified user from the data access object.
func (a *UserService) Delete(ctx context.Context, id comm.ID) error {
	exists, err := a.UserRepo.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "UserService not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.UserRepo.Delete(ctx, id); err != nil {
			return err
		}
		if err := a.UserRoleRepo.DeleteByUserID(ctx, id); err != nil {
			return err
		}
		return a.Cache.Delete(ctx, config.CacheNSForUser, comm.IDToStr(id))
	})
}

func (a *UserService) ResetPassword(ctx context.Context, id comm.ID) error {
	exists, err := a.UserRepo.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "UserService not found")
	}

	hashPass, err := hash.GeneratePassword(config.C.General.DefaultLoginPwd)
	if err != nil {
		return errors.BadRequest("", "Failed to generate hash password: %s", err.Error())
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.UserRepo.UpdatePasswordByID(ctx, id, hashPass); err != nil {
			return err
		}
		return nil
	})
}

func (a *UserService) GetRoleIDs(ctx context.Context, id comm.ID) ([]comm.ID, error) {
	userRoleResult, err := a.UserRoleRepo.Query(ctx, dto.UserRoleQueryParam{
		UserID: id,
	}, dto.UserRoleQueryOptions{
		QueryOptions: util.QueryOptions{
			SelectFields: []string{"role_id"},
		},
	})
	if err != nil {
		return nil, err
	}
	return userRoleResult.Data.ToRoleIDs(), nil
}

// GetUserInfo Get user info
func (a *UserService) GetUserInfo(ctx context.Context) (*dto.UserInfo, error) {
	if util.FromIsRootUser(ctx) {
		userID, err := comm.StrToID(config.C.General.Root.ID)
		if err != nil {
			return nil, err
		}
		permissions, err := a.MenuRepo.GetRolePermissions(ctx, 0)
		if err != nil {
			return nil, err
		}
		return &dto.UserInfo{
			ID:          userID,
			Username:    config.C.General.Root.Username,
			Name:        config.C.General.Root.Name,
			Permissions: permissions,
		}, nil
	}

	userIDStr := util.FromUserID(ctx)
	userID, err := comm.StrToID(userIDStr)
	if err != nil {
		return nil, err
	}
	user, err := a.UserRepo.Get(ctx, userID, dto.UserQueryOptions{
		QueryOptions: util.QueryOptions{
			OmitFields: []string{"password"},
		},
	})
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, errors.NotFound("", "UserService not found")
	}

	userRoleResult, err := a.UserRoleRepo.Query(ctx, dto.UserRoleQueryParam{
		UserID: userID,
	}, dto.UserRoleQueryOptions{
		JoinRole: true,
	})
	if err != nil {
		return nil, err
	}
	user.Roles = userRoleResult.Data

	var roleID comm.ID
	if len(user.Roles) > 0 {
		roleID = user.Roles[0].RoleID
	}

	permissions, err := a.MenuRepo.GetRolePermissions(ctx, roleID)
	if err != nil {
		return nil, err
	}

	return &dto.UserInfo{
		ID:          userID,
		Name:        user.Name,
		Username:    user.Username,
		Phone:       user.Phone,
		Email:       user.Email,
		Permissions: permissions,
	}, nil
}

// UpdatePassword Change login password
func (a *UserService) UpdatePassword(ctx context.Context, updateItem *dto.UpdateLoginPassword) error {
	if util.FromIsRootUser(ctx) {
		return errors.BadRequest("", "Root user cannot change password")
	}

	userIDStr := util.FromUserID(ctx)
	userID, err := comm.StrToID(userIDStr)
	user, err := a.UserRepo.Get(ctx, userID, dto.UserQueryOptions{
		QueryOptions: util.QueryOptions{
			SelectFields: []string{"password"},
		},
	})
	if err != nil {
		return err
	} else if user == nil {
		return errors.NotFound("", "UserService not found")
	}

	// check old password
	if err := hash.CompareHashAndPassword(user.Password, updateItem.OldPassword); err != nil {
		return errors.BadRequest("", "Incorrect old password")
	}

	// update password
	newPassword, err := hash.GeneratePassword(updateItem.NewPassword)
	if err != nil {
		return err
	}
	return a.UserRepo.UpdatePasswordByID(ctx, userID, newPassword)
}

// QueryMenus Query menus based on user permissions
func (a *UserService) QueryMenus(ctx context.Context) (entity.Menus, error) {
	menuQueryParams := dto.MenuQueryParam{
		Status: entity.MenuStatusEnabled,
	}

	isRoot := util.FromIsRootUser(ctx)
	if !isRoot {
		idStr := util.FromUserID(ctx)
		id, _ := comm.StrToID(idStr)
		menuQueryParams.UserID = id
	}
	menuResult, err := a.MenuRepo.Query(ctx, menuQueryParams, dto.MenuQueryOptions{
		QueryOptions: util.QueryOptions{
			OrderFields: entity.MenusOrderParams,
		},
	})
	if err != nil {
		return nil, err
	} else if isRoot {
		return menuResult.Data.ToTree(), nil
	}

	// fill parent menus
	if parentIDs := menuResult.Data.SplitParentIDs(); len(parentIDs) > 0 {
		var missMenusIDs []comm.ID
		menuIDMapper := menuResult.Data.ToMap()
		for _, parentID := range parentIDs {
			if _, ok := menuIDMapper[parentID]; !ok {
				missMenusIDs = append(missMenusIDs, parentID)
			}
		}
		if len(missMenusIDs) > 0 {
			parentResult, err := a.MenuRepo.Query(ctx, dto.MenuQueryParam{
				InIDs: comm.IDArrToStr(missMenusIDs),
			})
			if err != nil {
				return nil, err
			}
			menuResult.Data = append(menuResult.Data, parentResult.Data...)
			sort.Sort(menuResult.Data)
		}
	}

	return menuResult.Data.ToTree(), nil
}

// UpdateUser Update current user info
func (a *UserService) UpdateUser(ctx context.Context, updateItem *dto.UpdateCurrentUser) error {
	if util.FromIsRootUser(ctx) {
		return errors.BadRequest("", "Root user cannot update")
	}

	userIDStr := util.FromUserID(ctx)
	userID, err := comm.StrToID(userIDStr)
	if err != nil {
		return err
	}

	user, err := a.UserRepo.Get(ctx, userID)
	if err != nil {
		return err
	} else if user == nil {
		return errors.NotFound("", "UserService not found")
	}

	user.Name = updateItem.Name
	user.Phone = updateItem.Phone
	user.Email = updateItem.Email
	user.Remark = updateItem.Remark
	return a.UserRepo.Update(ctx, user, "name", "phone", "email", "remark")
}
