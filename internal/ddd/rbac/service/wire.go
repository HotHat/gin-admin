package service

import (
	"github.com/google/wire"
)

// ServiceSet Collection of wire providers
var ServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Struct(new(AuthService), "*"),
	wire.Struct(new(MenuService), "*"),
	wire.Struct(new(RoleService), "*"),
)
