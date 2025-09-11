package repo

import (
	"context"

	"github.com/HotHat/gin-admin/v10/internal/config"
	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/entity"
	"github.com/HotHat/gin-admin/v10/internal/mods/rbac/schema"
	"github.com/HotHat/gin-admin/v10/pkg/errors"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"gorm.io/gorm"
)

// GetMenuResourceDB Get menu resource storage instance
func GetMenuResourceDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(entity.MenuResource))
}

// MenuRepo resource management for RBAC
type MenuResourceRepo struct {
	DB *gorm.DB
}

func (a *MenuResourceRepo) TableName() string {
	return config.C.FormatTableName("menu_resource")
}

// Query menu resources from the database based on the provided parameters and options.
func (a *MenuResourceRepo) Query(ctx context.Context, params dto.MenuResourceQueryParam, opts ...dto.MenuResourceQueryOptions) (*dto.MenuResourceQueryResult, error) {
	var opt dto.MenuResourceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetMenuResourceDB(ctx, a.DB)
	if v := params.MenuID; len(v) > 0 {
		db = db.Where("menu_id = ?", v)
	}
	if v := params.MenuIDs; len(v) > 0 {
		db = db.Where("menu_id IN ?", v)
	}

	var list entity.MenuResources
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &dto.MenuResourceQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified menu resource from the database.
func (a *MenuResourceRepo) Get(ctx context.Context, id string, opts ...dto.MenuResourceQueryOptions) (*entity.MenuResource, error) {
	var opt dto.MenuResourceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(entity.MenuResource)
	ok, err := util.FindOne(ctx, GetMenuResourceDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exist checks if the specified menu resource exists in the database.
func (a *MenuResourceRepo) Exists(ctx context.Context, id comm.ID) (bool, error) {
	ok, err := util.Exists(ctx, GetMenuResourceDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// ExistsMethodPathByMenuID checks if the specified menu resource exists in the database.
func (a *MenuResourceRepo) ExistsMethodPathByMenuID(ctx context.Context, method, path string, menuID comm.ID) (bool, error) {
	ok, err := util.Exists(ctx, GetMenuResourceDB(ctx, a.DB).Where("method=? AND path=? AND menu_id=?", method, path, menuID))
	return ok, errors.WithStack(err)
}

// Create a new menu resource.
func (a *MenuResourceRepo) Create(ctx context.Context, item *entity.MenuResource) error {
	result := GetMenuResourceDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified menu resource in the database.
func (a *MenuResourceRepo) Update(ctx context.Context, item *entity.MenuResource) error {
	result := GetMenuResourceDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified menu resource from the database.
func (a *MenuResourceRepo) Delete(ctx context.Context, id comm.ID) error {
	result := GetMenuResourceDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.MenuResource))
	return errors.WithStack(result.Error)
}

// DeleteByMenuID Deletes the menu resource by menu id.
func (a *MenuResourceRepo) DeleteByMenuID(ctx context.Context, menuID comm.ID) error {
	result := GetMenuResourceDB(ctx, a.DB).Where("menu_id=?", menuID).Delete(new(entity.MenuResource))
	return errors.WithStack(result.Error)
}
