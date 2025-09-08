package entity

import (
	"time"

	"github.com/LyricTian/gin-admin/v10/internal/config"
)

// User roles for RBAC
type UserRole struct {
	ID        int       `json:"id" gorm:"primarykey"`                   // Unique ID
	UserID    string    `json:"user_id" gorm:"size:20;index"`           // From User.ID
	RoleID    string    `json:"role_id" gorm:"size:20;index"`           // From Role.ID
	CreatedAt time.Time `json:"created_at" gorm:"index;"`               // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`               // Update time
	RoleName  string    `json:"role_name" gorm:"<-:false;-:migration;"` // From Role.Name
}

func (a *UserRole) TableName() string {
	return config.C.FormatTableName("user_role")
}

// Defining the slice of `UserRole` struct.
type UserRoles []*UserRole

func (a UserRoles) ToUserIDMap() map[string]UserRoles {
	m := make(map[string]UserRoles)
	for _, userRole := range a {
		m[userRole.UserID] = append(m[userRole.UserID], userRole)
	}
	return m
}

func (a UserRoles) ToRoleIDs() []string {
	var ids []string
	for _, item := range a {
		ids = append(ids, item.RoleID)
	}
	return ids
}
