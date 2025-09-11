package service

import (
	"context"
	"fmt"
	"time"

	"github.com/HotHat/gin-admin/v10/internal/config"
	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/entity"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/repo"
	"github.com/HotHat/gin-admin/v10/pkg/cachex"
	"github.com/HotHat/gin-admin/v10/pkg/errors"
	"github.com/HotHat/gin-admin/v10/pkg/util"
)

// RoleService management for RBAC
type RoleService struct {
	Cache        cachex.Cacher
	Trans        *util.Trans
	RoleRepo     *repo.RoleRepo
	RoleMenuRepo *repo.RoleMenuRepo
	UserRoleRepo *repo.UserRoleRepo
}

// Query roles from the data access object based on the provided parameters and options.
func (a *RoleService) Query(ctx context.Context, params dto.RoleQueryParam) (*dto.RoleQueryResult, error) {
	params.Pagination = true

	var selectFields []string
	if params.ResultType == entity.RoleResultTypeSelect {
		params.Pagination = false
		selectFields = []string{"id", "name"}
	}

	result, err := a.RoleRepo.Query(ctx, params, dto.RoleQueryOptions{
		QueryOptions: util.QueryOptions{
			OrderFields: []util.OrderByParam{
				{Field: "sequence", Direction: util.DESC},
				{Field: "created_at", Direction: util.DESC},
			},
			SelectFields: selectFields,
		},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Get the specified role from the data access object.
func (a *RoleService) Get(ctx context.Context, id comm.ID) (*entity.Role, error) {
	role, err := a.RoleRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if role == nil {
		return nil, errors.NotFound("", "RoleService not found")
	}

	roleMenuResult, err := a.RoleMenuRepo.Query(ctx, dto.RoleMenuQueryParam{
		RoleID: id,
	})
	if err != nil {
		return nil, err
	}
	role.Menus = roleMenuResult.Data

	return role, nil
}

// Create a new role in the data access object.
func (a *RoleService) Create(ctx context.Context, formItem *dto.RoleForm) (*entity.Role, error) {
	if exists, err := a.RoleRepo.ExistsCode(ctx, formItem.Code); err != nil {
		return nil, err
	} else if exists {
		return nil, errors.BadRequest("", "RoleService code already exists")
	}

	role := &entity.Role{
		//ID:        util.NewXID(),
		CreatedAt: time.Now(),
	}
	if err := formItem.FillTo(role); err != nil {
		return nil, err
	}

	err := a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.RoleRepo.Create(ctx, role); err != nil {
			return err
		}

		for _, roleMenu := range formItem.Menus {
			//roleMenu.ID = util.NewXID()
			roleMenu.RoleID = role.ID
			roleMenu.CreatedAt = time.Now()
			if err := a.RoleMenuRepo.Create(ctx, roleMenu); err != nil {
				return err
			}
		}
		return a.syncToCasbin(ctx)
	})
	if err != nil {
		return nil, err
	}
	role.Menus = formItem.Menus

	return role, nil
}

// Update the specified role in the data access object.
func (a *RoleService) Update(ctx context.Context, id comm.ID, formItem *dto.RoleForm) error {
	role, err := a.RoleRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if role == nil {
		return errors.NotFound("", "RoleService not found")
	} else if role.Code != formItem.Code {
		if exists, err := a.RoleRepo.ExistsCode(ctx, formItem.Code); err != nil {
			return err
		} else if exists {
			return errors.BadRequest("", "RoleService code already exists")
		}
	}

	if err := formItem.FillTo(role); err != nil {
		return err
	}
	role.UpdatedAt = time.Now()

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.RoleRepo.Update(ctx, role); err != nil {
			return err
		}
		if err := a.RoleMenuRepo.DeleteByRoleID(ctx, id); err != nil {
			return err
		}
		for _, roleMenu := range formItem.Menus {
			//if roleMenu.ID == "" {
			//	roleMenu.ID = util.NewXID()
			//}
			roleMenu.RoleID = role.ID
			if roleMenu.CreatedAt.IsZero() {
				roleMenu.CreatedAt = time.Now()
			}
			roleMenu.UpdatedAt = time.Now()
			if err := a.RoleMenuRepo.Create(ctx, roleMenu); err != nil {
				return err
			}
		}
		return a.syncToCasbin(ctx)
	})
}

// Delete the specified role from the data access object.
func (a *RoleService) Delete(ctx context.Context, id comm.ID) error {
	exists, err := a.RoleRepo.Exists(ctx, id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "RoleService not found")
	}

	return a.Trans.Exec(ctx, func(ctx context.Context) error {
		if err := a.RoleRepo.Delete(ctx, id); err != nil {
			return err
		}
		if err := a.RoleMenuRepo.DeleteByRoleID(ctx, id); err != nil {
			return err
		}
		if err := a.UserRoleRepo.DeleteByRoleID(ctx, id); err != nil {
			return err
		}

		return a.syncToCasbin(ctx)
	})
}

func (a *RoleService) syncToCasbin(ctx context.Context) error {
	return a.Cache.Set(ctx, config.CacheNSForRole, config.CacheKeyForSyncToCasbin, fmt.Sprintf("%d", time.Now().Unix()))
}
