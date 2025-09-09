package route

import (
	"github.com/LyricTian/gin-admin/v10/pkg/comm"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Routes []comm.RouteRegister
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) Register(gin *gin.Engine) {
	for _, route := range r.Routes {
		route.Register(gin)
	}
}

func (r *Route) Add(reg comm.RouteRegister) {
	r.Routes = append(r.Routes, reg)
}
