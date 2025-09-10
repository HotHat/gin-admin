package entity

import (
	"time"

	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
)

const (
	UserStatusActivated = 1
	UserStatusFreezed   = 0
)

// User management for RBAC
type User struct {
	ID        comm.ID   `json:"id" gorm:"primarykey;"`         // Unique ID
	Username  string    `json:"username" gorm:"size:64;index"` // Username for login
	Name      string    `json:"name" gorm:"size:64;index"`     // Name of user
	Password  string    `json:"-" gorm:"size:64;"`             // Password for login (encrypted)
	Phone     string    `json:"phone" gorm:"size:32;"`         // Phone number of user
	Email     string    `json:"email" gorm:"size:128;"`        // Email of user
	Remark    string    `json:"remark" gorm:"size:1024;"`      // Remark of user
	Status    int       `json:"status" gorm:"index"`           // Status of user (activated, freezed)
	CreatedAt time.Time `json:"created_at" gorm:"index;"`      // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`      // Update time
	Roles     UserRoles `json:"roles" gorm:"-"`                // Roles of user
}

// Defining the slice of `User` struct.
type Users []*User

func (a Users) ToIDs() []comm.ID {
	var ids []comm.ID
	for _, item := range a {
		ids = append(ids, item.ID)
	}
	return ids
}
