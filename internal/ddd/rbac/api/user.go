package api

import (
	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/service"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"github.com/gin-gonic/gin"
)

// UserAPI management for RBAC
type UserAPI struct {
	UserService *service.UserService
	AuthService *service.AuthService
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Query user list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Param username query string false "Username for login"
// @Param name query string false "Name of user"
// @Param status query string false "Status of user (activated, freezed)"
// @Success 200 {object} util.ResponseResult{data=[]schema.User}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users [get]
func (a *UserAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	params := dto.UserQueryParam{
		Status: -1,
		PaginationParam: util.PaginationParam{
			Current:  1,
			PageSize: 15,
		},
	}

	if err := util.ParseQuery(c, &params); err != nil {
		util.RespError(c, err)
		return
	}

	result, err := a.UserService.Query(ctx, params)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespPage(c, result.Data, result.PageResult)
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Get user record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.User}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users/{id} [get]
func (a *UserAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := comm.StrToID(c.Param("id"))
	if err != nil {
		util.RespError(c, err)
	}
	item, err := a.UserService.Get(ctx, id)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespSuccess(c, item)
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Create user record
// @Param body body schema.UserForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.User}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users [post]
func (a *UserAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.UserForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.RespError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.RespError(c, err)
		return
	}

	result, err := a.UserService.Create(ctx, item)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespSuccess(c, result)
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Update user record by ID
// @Param id path string true "unique id"
// @Param body body schema.UserForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users/{id} [put]
func (a *UserAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.UserForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.RespError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.RespError(c, err)
		return
	}
	id, err := comm.StrToID(c.Param("id"))
	if err != nil {
		util.RespError(c, err)
	}
	err = a.UserService.Update(ctx, id, item)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespOK(c)
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Delete user record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users/{id} [delete]
func (a *UserAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := comm.StrToID(c.Param("id"))
	if err != nil {
		util.RespError(c, err)
	}
	err = a.UserService.Delete(ctx, id)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespOK(c)
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Reset user password by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users/{id}/reset-pwd [patch]
func (a *UserAPI) ResetPassword(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := comm.StrToID(c.Param("id"))
	if err != nil {
		util.RespError(c, err)
	}
	err = a.UserService.ResetPassword(ctx, id)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespOK(c)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Get current user info
// @Success 200 {object} util.ResponseResult{data=schema.User}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/current/user [get]
func (a *UserAPI) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := a.UserService.GetUserInfo(ctx)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespSuccess(c, data)
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
func (a *UserAPI) UpdatePassword(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.UpdateLoginPassword)
	if err := util.ParseJSON(c, item); err != nil {
		util.RespError(c, err)
		return
	}

	err := a.UserService.UpdatePassword(ctx, item)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespOK(c)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Query current user menus based on the current user role
// @Success 200 {object} util.ResponseResult{data=[]schema.Menu}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/current/menus [get]
func (a *UserAPI) QueryMenus(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := a.UserService.QueryMenus(ctx)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespSuccess(c, data)
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
func (a *UserAPI) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.UpdateCurrentUser)
	if err := util.ParseJSON(c, item); err != nil {
		util.RespError(c, err)
		return
	}

	err := a.UserService.UpdateUser(ctx, item)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespOK(c)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Logout system
// @Success 200 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/current/logout [post]
func (a *UserAPI) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.AuthService.Logout(ctx)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespOK(c)
}

// @Tags LoginAPI
// @Security ApiKeyAuth
// @Summary Refresh current access token
// @Success 200 {object} util.ResponseResult{data=schema.LoginToken}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/current/refresh-token [post]
func (a *UserAPI) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	data, err := a.AuthService.RefreshToken(ctx)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespSuccess(c, data)
}
