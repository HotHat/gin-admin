package api

import (
	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/service"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"github.com/gin-gonic/gin"
)

// RoleAPI management for RBAC
type RoleAPI struct {
	RoleService *service.RoleService
}

// @Tags RoleAPI
// @Security ApiKeyAuth
// @Summary Query role list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Param name query string false "Display name of role"
// @Param status query string false "Status of role (disabled, enabled)"
// @Success 200 {object} util.ResponseResult{data=[]schema.Role}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/roles [get]
func (a *RoleAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	params := dto.RoleQueryParam{
		ResultType: "select",
		Status:     -1,
		PaginationParam: util.PaginationParam{
			Current:  1,
			PageSize: 15,
		},
	}

	if err := util.ParseQuery(c, &params); err != nil {
		util.RespError(c, err)
		return
	}

	paramMap := c.Request.URL.Query()
	if len(paramMap) > 0 {
		params.ResultType = ""
	}

	result, err := a.RoleService.Query(ctx, params)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespPage(c, result.Data, result.PageResult)
}

// @Tags RoleAPI
// @Security ApiKeyAuth
// @Summary Get role record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Role}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/roles/{id} [get]
func (a *RoleAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id, _ := comm.StrToID(c.Param("id"))
	item, err := a.RoleService.Get(ctx, id)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespSuccess(c, item)
}

// @Tags RoleAPI
// @Security ApiKeyAuth
// @Summary Create role record
// @Param body body schema.RoleForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Role}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/roles [post]
func (a *RoleAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.RoleForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.RespError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.RespError(c, err)
		return
	}

	result, err := a.RoleService.Create(ctx, item)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespSuccess(c, result)
}

// @Tags RoleAPI
// @Security ApiKeyAuth
// @Summary Update role record by ID
// @Param id path string true "unique id"
// @Param body body schema.RoleForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/roles/{id} [put]
func (a *RoleAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.RoleForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.RespError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.RespError(c, err)
		return
	}

	id, _ := comm.StrToID(c.Param("id"))
	err := a.RoleService.Update(ctx, id, item)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespOK(c)
}

// @Tags RoleAPI
// @Security ApiKeyAuth
// @Summary Delete role record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/roles/{id} [delete]
func (a *RoleAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id, _ := comm.StrToID(c.Param("id"))
	err := a.RoleService.Delete(ctx, id)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespOK(c)
}
