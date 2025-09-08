package repo

import (
	"github.com/google/wire"
)

// EntitySet Collection of wire providers
var RepoSet = wire.NewSet(
	wire.Struct(new(MenuRepo), "*"),
	wire.Struct(new(MenuResourceRepo), "*"),
	wire.Struct(new(RoleRepo), "*"),
	wire.Struct(new(RoleMenuRepo), "*"),
	wire.Struct(new(UserRepo), "*"),
	wire.Struct(new(UserRoleRepo), "*"),
)
