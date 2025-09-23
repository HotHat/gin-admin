package api

import (
	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/service"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	"github.com/gin-gonic/gin"
)

// MenuAPI management for RBAC
type MenuAPI struct {
	MenuService *service.MenuService
}

// @Tags MenuAPI
// @Security ApiKeyAuth
// @Summary Query menu tree data
// @Param code query string false "Code path of menu (like xxx.xxx.xxx)"
// @Param name query string false "Name of menu"
// @Param includeResources query bool false "Whether to include menu resources"
// @Success 200 {object} util.ResponseResult{data=[]schema.Menu}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/menus [get]
func (a *MenuAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	params := dto.MenuQueryParam{
		Status: -1,
		PaginationParam: util.PaginationParam{
			Pagination: false,
			Current:    1,
			PageSize:   15,
		},
	}

	if err := util.ParseQuery(c, &params); err != nil {
		util.RespError(c, err)
		return
	}

	if len(c.Request.URL.RawQuery) == 0 {
		params.IncludeResources = true
	}

	result, err := a.MenuService.Query(ctx, params)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespPage(c, result.Data, result.PageResult)
}

// @Tags MenuAPI
// @Security ApiKeyAuth
// @Summary Get menu record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Menu}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/menus/{id} [get]
func (a *MenuAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	id, _ := comm.StrToID(c.Param("id"))
	item, err := a.MenuService.Get(ctx, id)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespSuccess(c, item)
}

// @Tags MenuAPI
// @Security ApiKeyAuth
// @Summary Create menu record
// @Param body body schema.MenuForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Menu}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/menus [post]
func (a *MenuAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.MenuForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.RespError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.RespError(c, err)
		return
	}

	result, err := a.MenuService.Create(ctx, item)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespSuccess(c, result)
}

// @Tags MenuAPI
// @Security ApiKeyAuth
// @Summary Update menu record by ID
// @Param id path string true "unique id"
// @Param body body schema.MenuForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/menus/{id} [put]
func (a *MenuAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(dto.MenuForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.RespError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.RespError(c, err)
		return
	}

	id, _ := comm.StrToID(c.Param("id"))
	err := a.MenuService.Update(ctx, id, item)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespOK(c)
}

// @Tags MenuAPI
// @Security ApiKeyAuth
// @Summary Delete menu record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/menus/{id} [delete]
func (a *MenuAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id, _ := comm.StrToID(c.Param("id"))
	err := a.MenuService.Delete(ctx, id)
	if err != nil {
		util.RespError(c, err)
		return
	}
	util.RespOK(c)
}
