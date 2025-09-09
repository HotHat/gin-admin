package admin

import (
	"github.com/LyricTian/gin-admin/v10/internal/mods/rbac/api"
	"github.com/gin-gonic/gin"
)

type RBACRouteV1 struct {
	LoginAPI *api.Login
	Handler  *AdminHandler
}

func (a *RBACRouteV1) RegisterRoute(g *gin.Engine) {
	mds := a.Handler.getHandlers()

	v1 := g.Group("v2", mds...)

	captcha := v1.Group("captcha")
	{
		captcha.GET("id", a.LoginAPI.GetCaptcha)
		captcha.GET("image", a.LoginAPI.ResponseCaptcha)
	}

	v1.POST("login", a.LoginAPI.Login)

}
