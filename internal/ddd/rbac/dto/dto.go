package dto

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/entity"
	"github.com/HotHat/gin-admin/v10/pkg/crypto/hash"
	"github.com/HotHat/gin-admin/v10/pkg/errors"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"github.com/go-playground/validator/v10"
)

type LoginForm struct {
	Username    string `json:"username" binding:"required"`     // Login name
	Password    string `json:"password" binding:"required"`     // Login password (md5 hash)
	CaptchaID   string `json:"captcha_id" binding:"required"`   // Captcha verify id
	CaptchaCode string `json:"captcha_code" binding:"required"` // Captcha verify code
}

func (a *LoginForm) Trim() *LoginForm {
	a.Username = strings.TrimSpace(a.Username)
	a.CaptchaCode = strings.TrimSpace(a.CaptchaCode)
	return a
}

// MenuForm Defining the data structure for creating a `Menu` struct.
type MenuForm struct {
	Code        string               `json:"code" binding:"required,max=32"`                   // Code of menu (unique for each level)
	Name        string               `json:"name" binding:"required,max=128"`                  // Display name of menu
	Description string               `json:"description"`                                      // Details about menu
	Sequence    int                  `json:"sequence"`                                         // Sequence for sorting (Order by desc)
	Type        string               `json:"type" binding:"required,oneof=page button"`        // Type of menu (page, button)
	Path        string               `json:"path"`                                             // Access path of menu
	Properties  string               `json:"properties"`                                       // Properties of menu (JSON)
	Status      uint                 `json:"status" binding:"required,oneof=disabled enabled"` // Status of menu (enabled, disabled)
	ParentID    comm.ID              `json:"parent_id"`                                        // Parent ID (From Menu.ID)
	Resources   entity.MenuResources `json:"resources"`                                        // Resources of menu
}

// Validate A validation function for the `MenuForm` struct.
func (a *MenuForm) Validate() error {
	if v := a.Properties; v != "" {
		if !json.Valid([]byte(v)) {
			return errors.BadRequest("", "invalid properties")
		}
	}
	return nil
}

func (a *MenuForm) FillTo(menu *entity.Menu) error {
	menu.Code = a.Code
	menu.Name = a.Name
	menu.Description = a.Description
	menu.Sequence = a.Sequence
	menu.Type = a.Type
	menu.Path = a.Path
	menu.Properties = a.Properties
	menu.Status = a.Status
	menu.ParentID = a.ParentID
	return nil
}

type MenuQueryParam struct {
	util.PaginationParam
	CodePath         string   `form:"code"`             // Code path (like xxx.xxx.xxx)
	LikeName         string   `form:"name"`             // Display name of menu
	IncludeResources bool     `form:"includeResources"` // Include resources
	InIDs            []string `form:"-"`                // Include menu IDs
	Status           uint     `form:"-"`                // Status of menu (disabled, enabled)
	ParentID         comm.ID  `form:"-"`                // Parent ID (From Menu.ID)
	ParentPathPrefix string   `form:"-"`                // Parent path (split by .)
	UserID           comm.ID  `form:"-"`                // User ID
	RoleID           comm.ID  `form:"-"`                // Role ID
}

// MenuQueryOptions Defining the query options for the `Menu` struct.
type MenuQueryOptions struct {
	util.QueryOptions
}

// MenuQueryResult Defining the query result for the `Menu` struct.
type MenuQueryResult struct {
	Data       entity.Menus
	PageResult *util.PaginationResult
}

// MenuResourceQueryParam Defining the query parameters for the `MenuResource` struct.
type MenuResourceQueryParam struct {
	util.PaginationParam
	MenuID  string   `form:"-"` // From Menu.ID
	MenuIDs []string `form:"-"` // From Menu.ID
}

// MenuResourceQueryOptions Defining the query options for the `MenuResource` struct.
type MenuResourceQueryOptions struct {
	util.QueryOptions
}

// MenuResourceQueryResult Defining the query result for the `MenuResource` struct.
type MenuResourceQueryResult struct {
	Data       entity.MenuResources
	PageResult *util.PaginationResult
}

// MenuResourceForm Defining the data structure for creating a `MenuResource` struct.
type MenuResourceForm struct {
}

// Validate A validation function for the `MenuResourceForm` struct.
func (a *MenuResourceForm) Validate() error {
	return nil
}

func (a *MenuResourceForm) FillTo(menuResource *entity.MenuResource) error {
	return nil
}

// RoleQueryParam Defining the query parameters for the `Role` struct.
type RoleQueryParam struct {
	util.PaginationParam
	LikeName    string     `form:"name"`                                       // Display name of role
	Status      string     `form:"status" binding:"oneof=disabled enabled ''"` // Status of role (disabled, enabled)
	ResultType  string     `form:"resultType"`                                 // Result type (options: select)
	InIDs       []string   `form:"-"`                                          // ID list
	GtUpdatedAt *time.Time `form:"-"`                                          // Update time is greater than
}

// RoleQueryOptions Defining the query options for the `Role` struct.
type RoleQueryOptions struct {
	util.QueryOptions
}

// RoleQueryResult Defining the query result for the `Role` struct.
type RoleQueryResult struct {
	Data       entity.Roles
	PageResult *util.PaginationResult
}

// RoleForm Defining the data structure for creating a `Role` struct.
type RoleForm struct {
	Code        string           `json:"code" binding:"required,max=32"`                   // Code of role (unique)
	Name        string           `json:"name" binding:"required,max=128"`                  // Display name of role
	Description string           `json:"description"`                                      // Details about role
	Sequence    int              `json:"sequence"`                                         // Sequence for sorting
	Status      string           `json:"status" binding:"required,oneof=disabled enabled"` // Status of role (enabled, disabled)
	Menus       entity.RoleMenus `json:"menus"`                                            // Role menu list
}

// Validate A validation function for the `RoleForm` struct.
func (a *RoleForm) Validate() error {
	return nil
}

func (a *RoleForm) FillTo(role *entity.Role) error {
	role.Code = a.Code
	role.Name = a.Name
	role.Description = a.Description
	role.Sequence = a.Sequence
	role.Status = a.Status
	return nil
}

// UserQueryParam Defining the query parameters for the `User` struct.
type UserQueryParam struct {
	util.PaginationParam
	LikeUsername string `form:"username"`                                    // Username for login
	LikeName     string `form:"name"`                                        // Name of user
	Status       int    `form:"status" binding:"oneof=activated freezed ''"` // Status of user (activated, freezed)
}

// UserQueryOptions Defining the query options for the `User` struct.
type UserQueryOptions struct {
	util.QueryOptions
}

// UserQueryResult Defining the query result for the `User` struct.
type UserQueryResult struct {
	Data       entity.Users
	PageResult *util.PaginationResult
}

// UserForm Defining the data structure for creating a `User` struct.
type UserForm struct {
	Username string           `json:"username" binding:"required,max=64"` // Username for login
	Name     string           `json:"name" binding:"required,max=64"`     // Name of user
	Password string           `json:"password" binding:"max=64"`          // Password for login (md5 hash)
	Phone    string           `json:"phone" binding:"max=32"`             // Phone number of user
	Email    string           `json:"email" binding:"max=128"`            // Email of user
	Remark   string           `json:"remark" binding:"max=1024"`          // Remark of user
	Status   int              `json:"status" binding:"required"`          // Status of user (activated, freezed)
	Roles    entity.UserRoles `json:"roles" binding:"required"`           // Roles of user
}

// A validation function for the `UserForm` struct.
func (a *UserForm) Validate() error {
	if a.Email != "" && validator.New().Var(a.Email, "email") != nil {
		return errors.BadRequest("", "Invalid email address")
	}
	return nil
}

// Convert `UserForm` to `User` object.
func (a *UserForm) FillTo(user *entity.User) error {
	user.Username = a.Username
	user.Name = a.Name
	user.Phone = a.Phone
	user.Email = a.Email
	user.Remark = a.Remark
	user.Status = a.Status

	if pass := a.Password; pass != "" {
		hashPass, err := hash.GeneratePassword(pass)
		if err != nil {
			return errors.BadRequest("", "Failed to generate hash password: %s", err.Error())
		}
		user.Password = hashPass
	}

	return nil
}

// UserRoleQueryParam Defining the query parameters for the `UserRole` struct.
type UserRoleQueryParam struct {
	util.PaginationParam
	InUserIDs []comm.ID `form:"-"` // From User.ID
	UserID    comm.ID   `form:"-"` // From User.ID
	RoleID    comm.ID   `form:"-"` // From Role.ID
}

// UserRoleQueryOptions Defining the query options for the `UserRole` struct.
type UserRoleQueryOptions struct {
	util.QueryOptions
	JoinRole bool // Join role table
}

// UserRoleQueryResult Defining the query result for the `UserRole` struct.
type UserRoleQueryResult struct {
	Data       entity.UserRoles
	PageResult *util.PaginationResult
}

// UserRoleForm Defining the data structure for creating a `UserRole` struct.
type UserRoleForm struct {
}

// Validate A validation function for the `UserRoleForm` struct.
func (a *UserRoleForm) Validate() error {
	return nil
}

func (a *UserRoleForm) FillTo(userRole *entity.UserRole) error {
	return nil
}

type RoleMenuQueryParam struct {
	util.PaginationParam
	RoleID comm.ID `form:"-"` // From Role.ID
}

// RoleMenuQueryOptions Defining the query options for the `RoleMenu` struct.
type RoleMenuQueryOptions struct {
	util.QueryOptions
}

// RoleMenuQueryResult Defining the query result for the `RoleMenu` struct.
type RoleMenuQueryResult struct {
	Data       entity.RoleMenus
	PageResult *util.PaginationResult
}

// RoleMenuForm Defining the data structure for creating a `RoleMenu` struct.
type RoleMenuForm struct {
}

// Validate A validation function for the `RoleMenuForm` struct.
func (a *RoleMenuForm) Validate() error {
	return nil
}

func (a *RoleMenuForm) FillTo(roleMenu *entity.RoleMenu) error {
	return nil
}

// LoggerQueryParam Defining the query parameters for the `Logger` struct.
type LoggerQueryParam struct {
	util.PaginationParam
	Level        string `form:"level"`     // Log level
	TraceID      string `form:"traceID"`   // Trace ID
	LikeUserName string `form:"userName"`  // User Name
	Tag          string `form:"tag"`       // Log tag
	LikeMessage  string `form:"message"`   // Log message
	StartTime    string `form:"startTime"` // Start time
	EndTime      string `form:"endTime"`   // End time
}

// Defining the query options for the `Logger` struct.
type LoggerQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `Logger` struct.
type LoggerQueryResult struct {
	Data       entity.Loggers
	PageResult *util.PaginationResult
}
