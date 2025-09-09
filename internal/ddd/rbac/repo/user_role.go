package repo

import (
	"context"
	"fmt"

	"github.com/LyricTian/gin-admin/v10/internal/config"
	"github.com/LyricTian/gin-admin/v10/internal/ddd/comm"
	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/entity"
	"github.com/LyricTian/gin-admin/v10/internal/mods/rbac/schema"
	"github.com/LyricTian/gin-admin/v10/pkg/errors"
	"github.com/LyricTian/gin-admin/v10/pkg/util"
	"gorm.io/gorm"
)

// Get user role storage instance
func GetUserRoleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(entity.UserRole))
}

// UserRepo roles for RBAC
type UserRoleRepo struct {
	DB *gorm.DB
}

func (a *UserRoleRepo) TableName() string {
	return config.C.FormatTableName("user_role")
}

// Query user roles from the database based on the provided parameters and options.
func (a *UserRoleRepo) Query(ctx context.Context, params dto.UserRoleQueryParam, opts ...dto.UserRoleQueryOptions) (*dto.UserRoleQueryResult, error) {
	var opt dto.UserRoleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := a.DB.Table(fmt.Sprintf("%s AS a", new(schema.UserRole).TableName()))
	if opt.JoinRole {
		db = db.Joins(fmt.Sprintf("left join %s b on a.role_id=b.id", new(schema.Role).TableName()))
		db = db.Select("a.*,b.name as role_name")
	}

	if v := params.InUserIDs; len(v) > 0 {
		db = db.Where("a.user_id IN (?)", v)
	}
	if v := params.UserID; v > 0 {
		db = db.Where("a.user_id = ?", v)
	}
	if v := params.RoleID; v > 0 {
		db = db.Where("a.role_id = ?", v)
	}

	var list entity.UserRoles
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &dto.UserRoleQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified user role from the database.
func (a *UserRoleRepo) Get(ctx context.Context, id int, opts ...dto.UserRoleQueryOptions) (*entity.UserRole, error) {
	var opt dto.UserRoleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(entity.UserRole)
	ok, err := util.FindOne(ctx, GetUserRoleDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists Exist checks if the specified user role exists in the database.
func (a *UserRoleRepo) Exists(ctx context.Context, id int) (bool, error) {
	ok, err := util.Exists(ctx, GetUserRoleDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new user role.
func (a *UserRoleRepo) Create(ctx context.Context, item *entity.UserRole) error {
	result := GetUserRoleDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified user role in the database.
func (a *UserRoleRepo) Update(ctx context.Context, item *entity.UserRole) error {
	result := GetUserRoleDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified user role from the database.
func (a *UserRoleRepo) Delete(ctx context.Context, id comm.ID) error {
	result := GetUserRoleDB(ctx, a.DB).Where("id=?", id).Delete(new(entity.UserRole))
	return errors.WithStack(result.Error)
}

func (a *UserRoleRepo) DeleteByUserID(ctx context.Context, userID comm.ID) error {
	result := GetUserRoleDB(ctx, a.DB).Where("user_id=?", userID).Delete(new(entity.UserRole))
	return errors.WithStack(result.Error)
}

func (a *UserRoleRepo) DeleteByRoleID(ctx context.Context, roleID comm.ID) error {
	result := GetUserRoleDB(ctx, a.DB).Where("role_id=?", roleID).Delete(new(entity.UserRole))
	return errors.WithStack(result.Error)
}
