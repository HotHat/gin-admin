package entity

import (
	"time"

	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
)

// Role permissions for RBAC
type RoleMenu struct {
	ID        comm.ID   `json:"id" gorm:"primarykey"`     // Unique ID
	RoleID    comm.ID   `json:"role_id" gorm:"index"`     // From Role.ID
	MenuID    comm.ID   `json:"menu_id" gorm:"index"`     // From Menu.ID
	CreatedAt time.Time `json:"created_at" gorm:"index;"` // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"` // Update time
}

// Defining the slice of `RoleMenu` struct.
type RoleMenus []*RoleMenu
