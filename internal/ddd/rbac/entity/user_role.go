package entity

import (
	"time"
)

// User roles for RBAC
type UserRole struct {
	ID        uint      `json:"id" gorm:"primarykey"`                   // Unique ID
	UserID    uint      `json:"user_id" gorm:"index"`                   // From User.ID
	RoleID    uint      `json:"role_id" gorm:"index"`                   // From Role.ID
	CreatedAt time.Time `json:"created_at" gorm:"index;"`               // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`               // Update time
	RoleName  string    `json:"role_name" gorm:"<-:false;-:migration;"` // From Role.Name
}

// Defining the slice of `UserRole` struct.
type UserRoles []*UserRole

func (a UserRoles) ToUserIDMap() map[uint]UserRoles {
	m := make(map[uint]UserRoles)
	for _, userRole := range a {
		m[userRole.UserID] = append(m[userRole.UserID], userRole)
	}
	return m
}

func (a UserRoles) ToRoleIDs() []uint {
	var ids []uint
	for _, item := range a {
		ids = append(ids, item.RoleID)
	}
	return ids
}
