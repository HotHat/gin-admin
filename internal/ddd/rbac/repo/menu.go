package repo

import (
	"context"

	"github.com/HotHat/gin-admin/v10/internal/config"
	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/entity"
	"github.com/HotHat/gin-admin/v10/pkg/errors"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"gorm.io/gorm"
)

// Get menu storage instance

func GetMenuDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(entity.Menu))
}

// MenuRepo management for RBAC
type MenuRepo struct {
	DB *gorm.DB
}

func (a *MenuRepo) TableName() string {
	return config.C.FormatTableName("menu")
}

// Query menus from the database based on the provided parameters and options.
func (a *MenuRepo) Query(ctx context.Context, params dto.MenuQueryParam, opts ...dto.MenuQueryOptions) (*dto.MenuQueryResult, error) {
	var opt dto.MenuQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetMenuDB(ctx, a.DB)

	if v := params.InIDs; len(v) > 0 {
		db = db.Where("id IN ?", v)
	}
	if v := params.LikeName; len(v) > 0 {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.Status; v != -1 {
		db = db.Where("status = ?", v)
	}
	if v := params.ParentID; v > 0 {
		db = db.Where("parent_id = ?", v)
	}
	if v := params.ParentPathPrefix; len(v) > 0 {
		db = db.Where("parent_path LIKE ?", v+"%")
	}
	if v := params.UserID; v > 0 {
		userRoleQuery := GetUserRoleDB(ctx, a.DB).Where("user_id = ?", v).Select("role_id")
		roleMenuQuery := GetRoleMenuDB(ctx, a.DB).Where("role_id IN (?)", userRoleQuery).Select("menu_id")
		db = db.Where("id IN (?)", roleMenuQuery)
	}
	if v := params.RoleID; v > 0 {
		roleMenuQuery := GetRoleMenuDB(ctx, a.DB).Where("role_id = ?", v).Select("menu_id")
		db = db.Where("id IN (?)", roleMenuQuery)
	}

	var list entity.Menus
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &dto.MenuQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified menu from the database.
func (a *MenuRepo) Get(ctx context.Context, id comm.ID, opts ...dto.MenuQueryOptions) (*entity.Menu, error) {
	var opt dto.MenuQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(entity.Menu)
	ok, err := util.FindOne(ctx, GetMenuDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

func (a *MenuRepo) GetByCodeAndParentID(ctx context.Context, code string, parentID comm.ID, opts ...dto.MenuQueryOptions) (*entity.Menu, error) {
	var opt dto.MenuQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(entity.Menu)
	ok, err := util.FindOne(ctx, GetMenuDB(ctx, a.DB).Where("code=? AND parent_id=?", code, parentID), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// GetByNameAndParentID get the specified menu from the database.
func (a *MenuRepo) GetByNameAndParentID(ctx context.Context, name string, parentID comm.ID, opts ...dto.MenuQueryOptions) (*entity.Menu, error) {
	var opt dto.MenuQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(entity.Menu)
	ok, err := util.FindOne(ctx, GetMenuDB(ctx, a.DB).Where("name=? AND parent_id=?", name, parentID), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists Checks if the specified menu exists in the database.
func (a *MenuRepo) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := util.Exists(ctx, GetMenuDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// ExistsCodeByParentID Checks if a menu with the specified `code` exists under the specified `parentID` in the database.
func (a *MenuRepo) ExistsCodeByParentID(ctx context.Context, code string, parentID comm.ID) (bool, error) {
	ok, err := util.Exists(ctx, GetMenuDB(ctx, a.DB).Where("code=? AND parent_id=?", code, parentID))
	return ok, errors.WithStack(err)
}

// ExistsNameByParentID Checks if a menu with the specified `name` exists under the specified `parentID` in the database.
func (a *MenuRepo) ExistsNameByParentID(ctx context.Context, name string, parentID comm.ID) (bool, error) {
	ok, err := util.Exists(ctx, GetMenuDB(ctx, a.DB).Where("name=? AND parent_id=?", name, parentID))
	return ok, errors.WithStack(err)
}

// Create a new menu.
func (a *MenuRepo) Create(ctx context.Context, item *entity.Menu) error {
	result := GetMenuDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified menu in the database.
func (a *MenuRepo) Update(ctx context.Context, item *entity.Menu) error {
	result := GetMenuDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified menu from the database.
func (a *MenuRepo) Delete(ctx context.Context, id comm.ID) error {
	result := GetMenuDB(ctx, a.DB).Where("id=?", id).Delete(new(entity.Menu))
	return errors.WithStack(result.Error)
}

// UpdateParentPath Updates the parent path of the specified menu.
func (a *MenuRepo) UpdateParentPath(ctx context.Context, id comm.ID, parentPath string) error {
	result := GetMenuDB(ctx, a.DB).Where("id=?", id).Update("parent_path", parentPath)
	return errors.WithStack(result.Error)
}

// UpdateStatusByParentPath Updates the status of all menus whose parent path starts with the provided parent path.
func (a *MenuRepo) UpdateStatusByParentPath(ctx context.Context, parentPath string, status int) error {
	result := GetMenuDB(ctx, a.DB).Where("parent_path like ?", parentPath+"%").Update("status", status)
	return errors.WithStack(result.Error)
}
