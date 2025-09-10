package entity

import (
	"time"

	"github.com/HotHat/gin-admin/v10/internal/config"
)

// Logger management
type Logger struct {
	ID        uint      `gorm:"primaryKey;" json:"id"`                   // Unique ID
	Level     string    `gorm:"size:20;index;" json:"level"`             // Log level
	TraceID   string    `gorm:"size:64;index;" json:"trace_id"`          // Trace ID
	UserID    string    `gorm:"size:20;index;" json:"user_id"`           // User ID
	Tag       string    `gorm:"size:32;index;" json:"tag"`               // Log tag
	Message   string    `gorm:"size:1024;" json:"message"`               // Log message
	Stack     string    `gorm:"type:text;" json:"stack"`                 // Error stack
	Data      string    `gorm:"type:text;" json:"data"`                  // Log data
	CreatedAt time.Time `gorm:"index;" json:"created_at"`                // Create time
	LoginName string    `json:"login_name" gorm:"<-:false;-:migration;"` // From User.Username
	UserName  string    `json:"user_name" gorm:"<-:false;-:migration;"`  // From User.Name
}

func (a *Logger) TableName() string {
	return config.C.FormatTableName("logger")
}

// Defining the slice of `Logger` struct.
type Loggers []*Logger
