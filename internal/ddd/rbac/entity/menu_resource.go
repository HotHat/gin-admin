package entity

import (
	"time"

	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
)

// Menu resource management for RBAC
type MenuResource struct {
	ID        comm.ID   `json:"id" gorm:"size:20;primarykey"` // Unique ID
	MenuID    comm.ID   `json:"menu_id" gorm:"size:20;index"` // From Menu.ID
	Method    string    `json:"method" gorm:"size:20;"`       // HTTP method
	Path      string    `json:"path" gorm:"size:255;"`        // API request path (e.g. /api/v1/users/:id)
	CreatedAt time.Time `json:"created_at" gorm:"index;"`     // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`     // Update time
}

// Defining the slice of `MenuResource` struct.
type MenuResources []*MenuResource
