package api

import (
	"github.com/google/wire"
)

// ApiSet Collection of wire providers
var ApiSet = wire.NewSet(
	wire.Struct(new(LoginAPI), "*"),
	wire.Struct(new(UserAPI), "*"),
	wire.Struct(new(MenuAPI), "*"),
	wire.Struct(new(RoleAPI), "*"),
)
