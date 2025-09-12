package api

import (
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/service"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"github.com/gin-gonic/gin"
)

type LoginAPI struct {
	AuthService *service.AuthService
}

// @Tags LoginAPI
// @Summary Get captcha ID
// @Success 200 {object} util.ResponseResult{data=schema.Captcha}
// @Router /api/v1/captcha/id [get]
func (a *LoginAPI) GetCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := a.AuthService.GetCaptcha(ctx)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespSuccess(c, data)
}

// @Tags LoginAPI
// @Summary Response captcha image
// @Param id query string true "Captcha ID"
// @Param reload query number false "Reload captcha image (reload=1)"
// @Produce image/png
// @Success 200 "Captcha image"
// @Failure 404 {object} util.ResponseResult
// @Router /api/v1/captcha/image [get]
func (a *LoginAPI) ResponseCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.AuthService.ResponseCaptcha(ctx, c.Writer, c.Query("id"), c.Query("reload") == "1")
	if err != nil {
		util.RespError(c, err)
	}
}

// @Tags LoginAPI
// @Summary LoginAPI system with username and password
// @Param body body schema.LoginForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.LoginToken}
// @Failure 400 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/login [post]
func (a *LoginAPI) Login(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.LoginForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.RespError(c, err)
		return
	}

	data, err := a.AuthService.Login(ctx, item.Trim())
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespSuccess(c, data)
}
