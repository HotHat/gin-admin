package entity

import (
	"time"

	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
)

// User roles for RBAC
type UserRole struct {
	ID        comm.ID   `json:"id" gorm:"primarykey"`                   // Unique ID
	UserID    comm.ID   `json:"user_id" gorm:"index"`                   // From User.ID
	RoleID    comm.ID   `json:"role_id" gorm:"index"`                   // From Role.ID
	CreatedAt time.Time `json:"created_at" gorm:"index;"`               // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`               // Update time
	RoleName  string    `json:"role_name" gorm:"<-:false;-:migration;"` // From Role.Name
}

// Defining the slice of `UserRole` struct.
type UserRoles []*UserRole

func (a UserRoles) ToUserIDMap() map[comm.ID]UserRoles {
	m := make(map[comm.ID]UserRoles)
	for _, userRole := range a {
		m[userRole.UserID] = append(m[userRole.UserID], userRole)
	}
	return m
}

func (a UserRoles) ToRoleIDs() []comm.ID {
	var ids []comm.ID
	for _, item := range a {
		ids = append(ids, item.RoleID)
	}
	return ids
}
