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

// Get user storage instance
func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDB(ctx, defDB).Model(new(entity.User))
}

// UserRepo management for RBAC
type UserRepo struct {
	DB *gorm.DB
}

func (a *UserRepo) TableName() string {
	return config.C.FormatTableName("user")
}

// Query users from the database based on the provided parameters and options.
func (a *UserRepo) Query(ctx context.Context, params dto.UserQueryParam, opts ...dto.UserQueryOptions) (*dto.UserQueryResult, error) {
	var opt dto.UserQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetUserDB(ctx, a.DB)
	if v := params.LikeUsername; len(v) > 0 {
		db = db.Where("username LIKE ?", "%"+v+"%")
	}
	if v := params.LikeName; len(v) > 0 {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.Status; v > 0 {
		db = db.Where("status = ?", v)
	}

	var list dto.UserQueryItemResults
	pageResult, err := util.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &dto.UserQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified user from the database.
func (a *UserRepo) Get(ctx context.Context, id comm.ID, opts ...dto.UserQueryOptions) (*entity.User, error) {
	var opt dto.UserQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(entity.User)
	ok, err := util.FindOne(ctx, GetUserDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

func (a *UserRepo) GetByUsername(ctx context.Context, username string, opts ...dto.UserQueryOptions) (*entity.User, error) {
	var opt dto.UserQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(entity.User)
	ok, err := util.FindOne(ctx, GetUserDB(ctx, a.DB).Where("username=?", username), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exists Exist checks if the specified user exists in the database.
func (a *UserRepo) Exists(ctx context.Context, id comm.ID) (bool, error) {
	ok, err := util.Exists(ctx, GetUserDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

func (a *UserRepo) ExistsUsername(ctx context.Context, username string) (bool, error) {
	ok, err := util.Exists(ctx, GetUserDB(ctx, a.DB).Where("username=?", username))
	return ok, errors.WithStack(err)
}

// Create a new user.
func (a *UserRepo) Create(ctx context.Context, item *entity.User) error {
	result := GetUserDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified user in the database.
func (a *UserRepo) Update(ctx context.Context, item *entity.User, selectFields ...string) error {
	db := GetUserDB(ctx, a.DB).Where("id=?", item.ID)
	if len(selectFields) > 0 {
		db = db.Select(selectFields)
	} else {
		db = db.Select("*").Omit("created_at")
	}
	result := db.Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified user from the database.
func (a *UserRepo) Delete(ctx context.Context, id comm.ID) error {
	result := GetUserDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.User))
	return errors.WithStack(result.Error)
}

func (a *UserRepo) UpdatePasswordByID(ctx context.Context, id comm.ID, password string) error {
	result := GetUserDB(ctx, a.DB).Where("id=?", id).Select("password").Updates(entity.User{Password: password})
	return errors.WithStack(result.Error)
}
