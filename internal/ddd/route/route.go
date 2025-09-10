package route

import (
	"context"

	"github.com/HotHat/gin-admin/v10/pkg/comm"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Routes []comm.RouteRegister
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) Register(ctx context.Context, gin *gin.Engine) {
	for _, route := range r.Routes {
		route.Register(ctx, gin)
	}
}

func (r *Route) Add(reg comm.RouteRegister) {
	r.Routes = append(r.Routes, reg)
}
