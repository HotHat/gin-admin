package v1

import (
	"github.com/LyricTian/gin-admin/v10/internal/mods/rbac/api"
	"github.com/gin-gonic/gin"
)

type ApiV1 struct {
	LoginAPI *api.Login
}

func (a ApiV1) RegisterRoute(gin *gin.Engine) {
	v1 := gin.Group("v2")

	captcha := v1.Group("captcha")
	{
		captcha.GET("id", a.LoginAPI.GetCaptcha)
		captcha.GET("image", a.LoginAPI.ResponseCaptcha)
	}

	v1.POST("login", a.LoginAPI.Login)

}
