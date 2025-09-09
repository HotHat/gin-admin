package admin

import (
	"context"
	"path/filepath"

	"github.com/LyricTian/gin-admin/v10/internal/config"
	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/entity"
	"github.com/LyricTian/gin-admin/v10/pkg/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	apiPrefix = "/api/admin"
)

type AdminMod struct {
	DB      *gorm.DB
	Casbinx *Casbinx
}

func (a *AdminMod) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(
		new(entity.Menu),
		new(entity.MenuResource),
		new(entity.Role),
		new(entity.RoleMenu),
		new(entity.User),
		new(entity.UserRole),
	)
}

func (a *AdminMod) Init(ctx context.Context) error {
	if config.C.Storage.DB.AutoMigrate {
		if err := a.AutoMigrate(ctx); err != nil {
			return err
		}
	}

	if err := a.Casbinx.Load(ctx); err != nil {
		return err
	}

	if name := config.C.General.MenuFile; name != "" {
		fullPath := filepath.Join(config.C.General.WorkDir, name)
		if err := a.MenuAPI.MenuBIZ.InitFromFile(ctx, fullPath); err != nil {
			logging.Context(ctx).Error("failed to init menu data", zap.Error(err), zap.String("file", fullPath))
		}
	}

	return nil
}

func (a *AdminMod) RouterPrefixes() []string {
	return []string{
		apiPrefix,
	}
}

func (a *AdminMod) RegisterRouters(ctx context.Context, e *gin.Engine) error {
	gAPI := e.Group(apiPrefix)
	v1 := gAPI.Group("v1")

	if err := a.RBAC.RegisterV1Routers(ctx, v1); err != nil {
		return err
	}

	return nil
}

func (a *AdminMod) Release(ctx context.Context) error {
	if err := a.RBAC.Release(ctx); err != nil {
		return err
	}

	return nil
}
