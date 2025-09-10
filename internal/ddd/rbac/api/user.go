package api

import (
	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/service"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"github.com/gin-gonic/gin"
)

// UserApi management for RBAC
type UserApi struct {
	UserService *service.UserService
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
func (a *UserApi) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto.UserQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.UserService.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Get user record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.User}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users/{id} [get]
func (a *UserApi) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := comm.StrToID(c.Param("id"))
	if err != nil {
		util.ResError(c, err)
	}
	item, err := a.UserService.Get(ctx, id)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
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
func (a *UserApi) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.UserForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.UserService.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
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
func (a *UserApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.UserForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}
	id, err := comm.StrToID(c.Param("id"))
	if err != nil {
		util.ResError(c, err)
	}
	err = a.UserService.Update(ctx, id, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Delete user record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users/{id} [delete]
func (a *UserApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := comm.StrToID(c.Param("id"))
	if err != nil {
		util.ResError(c, err)
	}
	err = a.UserService.Delete(ctx, id)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags UserAPI
// @Security ApiKeyAuth
// @Summary Reset user password by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/users/{id}/reset-pwd [patch]
func (a *UserApi) ResetPassword(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := comm.StrToID(c.Param("id"))
	if err != nil {
		util.ResError(c, err)
	}
	err = a.UserService.ResetPassword(ctx, id)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
