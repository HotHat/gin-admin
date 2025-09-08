package entity

import (
	"time"

	"github.com/LyricTian/gin-admin/v10/internal/config"
)

const (
	RoleStatusEnabled  = "enabled"  // Enabled
	RoleStatusDisabled = "disabled" // Disabled

	RoleResultTypeSelect = "select" // Select
)

// Role management for RBAC
type Role struct {
	ID          int       `json:"id" gorm:"size:20;primarykey;"` // Unique ID
	Code        string    `json:"code" gorm:"size:32;index;"`    // Code of role (unique)
	Name        string    `json:"name" gorm:"size:128;index"`    // Display name of role
	Description string    `json:"description" gorm:"size:1024"`  // Details about role
	Sequence    int       `json:"sequence" gorm:"index"`         // Sequence for sorting
	Status      string    `json:"status" gorm:"size:20;index"`   // Status of role (disabled, enabled)
	CreatedAt   time.Time `json:"created_at" gorm:"index;"`      // Create time
	UpdatedAt   time.Time `json:"updated_at" gorm:"index;"`      // Update time
	Menus       RoleMenus `json:"menus" gorm:"-"`                // Role menu list
}

func (a *Role) TableName() string {
	return config.C.FormatTableName("role")
}

// Defining the slice of `Role` struct.
type Roles []*Role
