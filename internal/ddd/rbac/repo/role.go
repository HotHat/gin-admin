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

// GetRoleDB Get role storage instance
func GetRoleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(entity.Role))
}

// RoleRepo management for RBAC
type RoleRepo struct {
	DB *gorm.DB
}

func (a *RoleRepo) TableName() string {
	return config.C.FormatTableName("role")
}

// Query roles from the database based on the provided parameters and options.
func (a *RoleRepo) Query(ctx context.Context, params dto.RoleQueryParam, opts ...dto.RoleQueryOptions) (*dto.RoleQueryResult, error) {
	var opt dto.RoleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetRoleDB(ctx, a.DB)
	if v := params.InIDs; len(v) > 0 {
		db = db.Where("id IN (?)", v)
	}
	if v := params.LikeName; len(v) > 0 {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.Status; len(v) > 0 {
		db = db.Where("status = ?", v)
	}
	if v := params.GtUpdatedAt; v != nil {
		db = db.Where("updated_at > ?", v)
	}

	var list entity.Roles
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &dto.RoleQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified role from the database.
func (a *RoleRepo) Get(ctx context.Context, id comm.ID, opts ...dto.RoleQueryOptions) (*entity.Role, error) {
	var opt dto.RoleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(entity.Role)
	ok, err := util.FindOne(ctx, GetRoleDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exist checks if the specified role exists in the database.
func (a *RoleRepo) Exists(ctx context.Context, id comm.ID) (bool, error) {
	ok, err := util.Exists(ctx, GetRoleDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

func (a *RoleRepo) ExistsCode(ctx context.Context, code string) (bool, error) {
	ok, err := util.Exists(ctx, GetRoleDB(ctx, a.DB).Where("code=?", code))
	return ok, errors.WithStack(err)
}

// Create a new role.
func (a *RoleRepo) Create(ctx context.Context, item *entity.Role) error {
	result := GetRoleDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified role in the database.
func (a *RoleRepo) Update(ctx context.Context, item *entity.Role) error {
	result := GetRoleDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified role from the database.
func (a *RoleRepo) Delete(ctx context.Context, id comm.ID) error {
	result := GetRoleDB(ctx, a.DB).Where("id=?", id).Delete(new(entity.Role))
	return errors.WithStack(result.Error)
}
