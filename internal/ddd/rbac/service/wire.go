package service

import (
	"github.com/google/wire"
)

// ServiceSet Collection of wire providers
var ServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Struct(new(AuthService), "*"),
)
