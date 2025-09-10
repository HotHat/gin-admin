package repo

import (
	"context"

	"github.com/HotHat/gin-admin/v10/internal/config"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/entity"
	"github.com/HotHat/gin-admin/v10/pkg/errors"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"gorm.io/gorm"
)

// Get role menu storage instance
func GetRoleMenuDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(entity.RoleMenu))
}

// RoleRepo permissions for RBAC
type RoleMenuRepo struct {
	DB *gorm.DB
}

func (a *RoleMenuRepo) TableName() string {
	return config.C.FormatTableName("role_menu")
}

// Query role menus from the database based on the provided parameters and options.
func (a *RoleMenuRepo) Query(ctx context.Context, params dto.RoleMenuQueryParam, opts ...dto.RoleMenuQueryOptions) (*dto.RoleMenuQueryResult, error) {
	var opt dto.RoleMenuQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetRoleMenuDB(ctx, a.DB)
	if v := params.RoleID; v > 0 {
		db = db.Where("role_id = ?", v)
	}

	var list entity.RoleMenus
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &dto.RoleMenuQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified role menu from the database.
func (a *RoleMenuRepo) Get(ctx context.Context, id string, opts ...dto.RoleMenuQueryOptions) (*entity.RoleMenu, error) {
	var opt dto.RoleMenuQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(entity.RoleMenu)
	ok, err := util.FindOne(ctx, GetRoleMenuDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exist checks if the specified role menu exists in the database.
func (a *RoleMenuRepo) Exists(ctx context.Context, id int) (bool, error) {
	ok, err := util.Exists(ctx, GetRoleMenuDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new role menu.
func (a *RoleMenuRepo) Create(ctx context.Context, item *entity.RoleMenu) error {
	result := GetRoleMenuDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified role menu in the database.
func (a *RoleMenuRepo) Update(ctx context.Context, item *entity.RoleMenu) error {
	result := GetRoleMenuDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified role menu from the database.
func (a *RoleMenuRepo) Delete(ctx context.Context, id int) error {
	result := GetRoleMenuDB(ctx, a.DB).Where("id=?", id).Delete(new(entity.RoleMenu))
	return errors.WithStack(result.Error)
}

// Deletes role menus by role id.
func (a *RoleMenuRepo) DeleteByRoleID(ctx context.Context, roleID int) error {
	result := GetRoleMenuDB(ctx, a.DB).Where("role_id=?", roleID).Delete(new(entity.RoleMenu))
	return errors.WithStack(result.Error)
}

// Deletes role menus by menu id.
func (a *RoleMenuRepo) DeleteByMenuID(ctx context.Context, menuID int) error {
	result := GetRoleMenuDB(ctx, a.DB).Where("menu_id=?", menuID).Delete(new(entity.RoleMenu))
	return errors.WithStack(result.Error)
}
