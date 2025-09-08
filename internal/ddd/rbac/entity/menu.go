package entity

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/LyricTian/gin-admin/v10/internal/config"
	"github.com/LyricTian/gin-admin/v10/pkg/errors"
	"github.com/LyricTian/gin-admin/v10/pkg/util"
)

const (
	MenuStatusDisabled = "disabled"
	MenuStatusEnabled  = "enabled"
)

var (
	MenusOrderParams = []util.OrderByParam{
		{Field: "sequence", Direction: util.DESC},
		{Field: "created_at", Direction: util.DESC},
	}
)

// Menu management for RBAC
type Menu struct {
	ID          int           `json:"id" gorm:"primarykey;"`              // Unique ID
	Code        string        `json:"code" gorm:"size:32;index;"`         // Code of menu (unique for each level)
	Name        string        `json:"name" gorm:"size:128;index"`         // Display name of menu
	Description string        `json:"description" gorm:"size:1024"`       // Details about menu
	Sequence    int           `json:"sequence" gorm:"index;"`             // Sequence for sorting (Order by desc)
	Type        string        `json:"type" gorm:"size:20;index"`          // Type of menu (page, button)
	Path        string        `json:"path" gorm:"size:255;"`              // Access path of menu
	Properties  string        `json:"properties" gorm:"type:text;"`       // Properties of menu (JSON)
	Status      string        `json:"status" gorm:"size:20;index"`        // Status of menu (enabled, disabled)
	ParentID    int           `json:"parent_id" gorm:"size:20;index;"`    // Parent ID (From Menu.ID)
	ParentPath  string        `json:"parent_path" gorm:"size:255;index;"` // Parent path (split by .)
	Children    *Menus        `json:"children" gorm:"-"`                  // Child menus
	CreatedAt   time.Time     `json:"created_at" gorm:"index;"`           // Create time
	UpdatedAt   time.Time     `json:"updated_at" gorm:"index;"`           // Update time
	Resources   MenuResources `json:"resources" gorm:"-"`                 // Resources of menu
}

func (a *Menu) TableName() string {
	return config.C.FormatTableName("menu")
}

// Defining the query parameters for the `Menu` struct.

// Defining the slice of `Menu` struct.
type Menus []*Menu

func (a Menus) Len() int {
	return len(a)
}

func (a Menus) Less(i, j int) bool {
	if a[i].Sequence == a[j].Sequence {
		return a[i].CreatedAt.Unix() > a[j].CreatedAt.Unix()
	}
	return a[i].Sequence > a[j].Sequence
}

func (a Menus) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Menus) ToMap() map[int]*Menu {
	m := make(map[int]*Menu)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}

func (a Menus) SplitParentIDs() []string {
	parentIDs := make([]string, 0, len(a))
	idMapper := make(map[int]struct{})
	for _, item := range a {
		if _, ok := idMapper[item.ID]; ok {
			continue
		}
		idMapper[item.ID] = struct{}{}
		if pp := item.ParentPath; pp != "" {
			for _, pid := range strings.Split(pp, util.TreePathDelimiter) {
				if pid == "" {
					continue
				}
				pidInt, _ := strconv.ParseInt(pid, 10, 64)
				var pid32 = int(pidInt)

				if _, ok := idMapper[pid32]; ok {
					continue
				}
				parentIDs = append(parentIDs, pid)
				idMapper[pid32] = struct{}{}
			}
		}
	}
	return parentIDs
}

func (a Menus) ToTree() Menus {
	var list Menus
	m := a.ToMap()
	for _, item := range a {
		if item.ParentID == 0 {
			list = append(list, item)
			continue
		}
		if parent, ok := m[item.ParentID]; ok {
			if parent.Children == nil {
				children := Menus{item}
				parent.Children = &children
				continue
			}
			*parent.Children = append(*parent.Children, item)
		}
	}
	return list
}
