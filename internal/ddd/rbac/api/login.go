package api

import (
	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/service"
	"github.com/LyricTian/gin-admin/v10/pkg/util"
	"github.com/gin-gonic/gin"
)

type LoginApi struct {
	AuthService *service.AuthService
}

// @Tags LoginAPI
// @Summary Get captcha ID
// @Success 200 {object} util.ResponseResult{data=schema.Captcha}
// @Router /api/v1/captcha/id [get]
func (a *LoginApi) GetCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := a.AuthService.GetCaptcha(ctx)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, data)
}

// @Tags LoginAPI
// @Summary Response captcha image
// @Param id query string true "Captcha ID"
// @Param reload query number false "Reload captcha image (reload=1)"
// @Produce image/png
// @Success 200 "Captcha image"
// @Failure 404 {object} util.ResponseResult
// @Router /api/v1/captcha/image [get]
func (a *LoginApi) ResponseCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.AuthService.ResponseCaptcha(ctx, c.Writer, c.Query("id"), c.Query("reload") == "1")
	if err != nil {
		util.ResError(c, err)
	}
}

// @Tags LoginAPI
// @Summary LoginApi system with username and password
// @Param body body schema.LoginForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.LoginToken}
// @Failure 400 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/login [post]
func (a *LoginApi) Login(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.LoginForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	}

	data, err := a.AuthService.Login(ctx, item.Trim())
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, data)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Logout system
// @Success 200 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/current/logout [post]
func (a *LoginApi) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.AuthService.Logout(ctx)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Refresh current access token
// @Success 200 {object} util.ResponseResult{data=schema.LoginToken}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/current/refresh-token [post]
func (a *LoginApi) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := a.AuthService.RefreshToken(ctx)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, data)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Get current user info
// @Success 200 {object} util.ResponseResult{data=schema.User}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/current/user [get]
func (a *LoginApi) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := a.AuthService.GetUserInfo(ctx)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, data)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Change current user password
// @Param body body schema.UpdateLoginPassword true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/current/password [put]
func (a *LoginApi) UpdatePassword(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.UpdateLoginPassword)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.AuthService.UpdatePassword(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Query current user menus based on the current user role
// @Success 200 {object} util.ResponseResult{data=[]schema.Menu}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/current/menus [get]
func (a *LoginApi) QueryMenus(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := a.AuthService.QueryMenus(ctx)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, data)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Update current user info
// @Param body body schema.UpdateCurrentUser true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/current/user [put]
func (a *LoginApi) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.UpdateCurrentUser)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.AuthService.UpdateUser(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
