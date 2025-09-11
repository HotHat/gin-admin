package admin

import (
	"context"

	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/api"
	"github.com/gin-gonic/gin"
)

type RBACRouteV1 struct {
	Handler  *AdminHandler
	LoginAPI *api.LoginAPI
	UserAPI  *api.UserAPI
	MenuAPI  *api.MenuAPI
	RoleAPI  *api.RoleAPI
}

func (a *RBACRouteV1) Release(ctx context.Context) error {
	//return a.Handler.Release(ctx)
	return nil
}

func (a *RBACRouteV1) Register(ctx context.Context, g *gin.Engine) error {
	mds := a.Handler.GetHandlers()

	v1 := g.Group(baseAPI+"v1", mds...)

	captcha := v1.Group("captcha")
	{
		captcha.GET("id", a.LoginAPI.GetCaptcha)
		captcha.GET("image", a.LoginAPI.ResponseCaptcha)
	}

	v1.POST("login", a.LoginAPI.Login)

	user := v1.Group("user")
	{
		user.GET("", a.UserAPI.GetUserInfo)
		user.POST("refresh-token", a.UserAPI.RefreshToken)
		user.GET("menus", a.UserAPI.QueryMenus)
		user.PUT("password", a.UserAPI.UpdatePassword)
		user.PUT("", a.UserAPI.UpdateUser)
		user.POST("logout", a.UserAPI.Logout)
	}

	menu := v1.Group("menus")
	{
		menu.GET("", a.MenuAPI.Query)
		menu.GET(":id", a.MenuAPI.Get)
		menu.POST("", a.MenuAPI.Create)
		menu.PUT(":id", a.MenuAPI.Update)
		menu.DELETE(":id", a.MenuAPI.Delete)
	}

	role := v1.Group("roles")
	{
		role.GET("", a.RoleAPI.Query)
		role.GET(":id", a.RoleAPI.Get)
		role.POST("", a.RoleAPI.Create)
		role.PUT(":id", a.RoleAPI.Update)
		role.DELETE(":id", a.RoleAPI.Delete)
	}

	users := v1.Group("users")
	{
		users.GET("", a.UserAPI.Query)
		users.GET(":id", a.UserAPI.Get)
		users.POST("", a.UserAPI.Create)
		users.PUT(":id", a.UserAPI.Update)
		users.DELETE(":id", a.UserAPI.Delete)
		users.PATCH(":id/reset-pwd", a.UserAPI.ResetPassword)
	}

	//logger := v1.Group("loggers")
	//{
	//	logger.GET("", a.LsoggerAPI.Query)
	//}
	return nil
}
