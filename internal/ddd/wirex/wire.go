//go:build wireinject
// +build wireinject

package wirex

import (
	"context"

	rbacApi "github.com/HotHat/gin-admin/v10/internal/ddd/rbac/api"
	rbacRepo "github.com/HotHat/gin-admin/v10/internal/ddd/rbac/repo"
	rbacService "github.com/HotHat/gin-admin/v10/internal/ddd/rbac/service"
	"github.com/HotHat/gin-admin/v10/internal/ddd/route/admin"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"github.com/google/wire"
)

func BuildInjector(ctx context.Context) (*Injector, func(), error) {
	wire.Build(
		InitDB,
		InitCacher,
		InitAuth,
		wire.NewSet(wire.Struct(new(util.Trans), "*")),
		//entity.EntitySet,
		rbacRepo.RepoSet,
		rbacService.ServiceSet,
		rbacApi.ApiSet,
		wire.Struct(new(admin.Casbinx), "*"),
		//wire.Struct(new(admin.AdminHandler), "*"),
		admin.InitAdminHandler,
		wire.Struct(new(admin.RBACRouteV1), "*"),
		wire.Struct(new(Injector), "*"),
	) // end
	return &Injector{}, nil, nil
}
