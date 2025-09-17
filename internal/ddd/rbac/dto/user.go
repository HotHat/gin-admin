package dto

import (
	"time"

	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/pkg/util"
)

// UserQueryParam Defining the query parameters for the `User` struct.
type UserQueryParam struct {
	util.PaginationParam
	LikeUsername string `form:"username"` // Username for login
	LikeName     string `form:"name"`     // Name of user
	Status       int    `form:"status"`   // Status of user (activated, freezed)
}

// UserQueryOptions Defining the query options for the `User` struct.
type UserQueryOptions struct {
	util.QueryOptions
}

// UserQueryResult Defining the query result for the `User` struct.
type UserQueryResult struct {
	Data       UserQueryItemResults
	PageResult *util.PaginationResult
}

type UserRoleItem struct {
	ID   comm.ID `json:"id" binding:"required"`
	Name string  `json:"name" binding:"required,max=64"`
}

type UserQueryItemResult struct {
	ID        comm.ID        `json:"id" binding:"required,max=64"`       // id for login
	Username  string         `json:"username" binding:"required,max=64"` // Username for login
	Name      string         `json:"name" binding:"required,max=64"`     // Name of user
	Phone     string         `json:"phone" binding:"max=32"`             // Phone number of user
	Email     string         `json:"email" binding:"max=128"`            // Email of user
	Remark    string         `json:"remark" binding:"max=1024"`          // Remark of user
	Status    int            `json:"status" binding:"required"`          // Status of user (activated, freezed)
	Roles     []UserRoleItem `json:"roles" binding:"required"`           // Roles of user
	CreatedAt time.Time      `json:"created_at" gorm:"index;"`           // Create time
	UpdatedAt time.Time      `json:"updated_at" gorm:"index;"`           // Update time
}

type UserQueryItemResults []*UserQueryItemResult

func (a UserQueryItemResults) ToIDs() []comm.ID {
	var ids []comm.ID
	for _, item := range a {
		ids = append(ids, item.ID)
	}
	return ids
}
