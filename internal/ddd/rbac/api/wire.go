package api

import (
	"github.com/google/wire"
)

// ApiSet Collection of wire providers
var ApiSet = wire.NewSet(
	wire.Struct(new(LoginApi), "*"),
	wire.Struct(new(UserApi), "*"),
)
