//go:build wireinject
// +build wireinject

package service

import (
	"context"

	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/repo"
	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/service"
	"github.com/LyricTian/gin-admin/v10/internal/wirex"
	"github.com/LyricTian/gin-admin/v10/pkg/util"
	"github.com/google/wire"
)

func BuildService(ctx context.Context) (*ServiceTest, func(), error) {
	wire.Build(
		wirex.InitDB,
		wirex.InitCacher,
		wirex.InitAuth,
		wire.NewSet(wire.Struct(new(util.Trans), "*")),
		//entity.EntitySet,
		repo.RepoSet,
		service.ServiceSet,
		wire.Struct(new(ServiceTest), "*"),
	) // end
	return &ServiceTest{}, nil, nil
}
