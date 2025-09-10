package admin

import (
	"context"

	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/api"
	"github.com/gin-gonic/gin"
)

type RBACRouteV1 struct {
	LoginAPI *api.LoginApi
	Handler  *AdminHandler
}

func (a *RBACRouteV1) Release(ctx context.Context) error {
	return a.Handler.Release(ctx)
}

func (a *RBACRouteV1) Register(ctx context.Context, g *gin.Engine) {
	mds := a.Handler.GetHandlers()

	v1 := g.Group("v2", mds...)

	captcha := v1.Group("captcha")
	{
		captcha.GET("id", a.LoginAPI.GetCaptcha)
		captcha.GET("image", a.LoginAPI.ResponseCaptcha)
	}

	v1.POST("login", a.LoginAPI.Login)

}
