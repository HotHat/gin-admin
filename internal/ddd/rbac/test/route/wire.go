//go:build wireinject
// +build wireinject

package route

import (
	"context"

	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/api"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/repo"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/service"
	"github.com/HotHat/gin-admin/v10/internal/ddd/route/admin"
	"github.com/HotHat/gin-admin/v10/internal/wirex"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"github.com/google/wire"
)

func BuildRoute(ctx context.Context) (*RouteTest, func(), error) {
	wire.Build(
		wirex.InitDB,
		wirex.InitCacher,
		wirex.InitAuth,
		wire.NewSet(wire.Struct(new(util.Trans), "*")),
		//entity.EntitySet,
		repo.RepoSet,
		service.ServiceSet,
		api.ApiSet,
		wire.Struct(new(admin.Casbinx), "*"),
		wire.Struct(new(admin.AdminHandler), "*"),
		wire.Struct(new(admin.RBACRouteV1), "*"),
		wire.Struct(new(RouteTest), "*"),
	) // end
	return &RouteTest{}, nil, nil
}
