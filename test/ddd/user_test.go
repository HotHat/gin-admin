package ddd

import (
	"net/http"
	"testing"

	"github.com/HotHat/gin-admin/v10/internal/ddd/comm"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/entity"
	"github.com/HotHat/gin-admin/v10/pkg/crypto/hash"
	"github.com/HotHat/gin-admin/v10/pkg/util"
	Assert "github.com/stretchr/testify/assert"
)

func TestUserInfo(t *testing.T) {
	e := tester(t)
	var user dto.UserInfo
	e.GET(baseAPI+"/user").WithHeader(
		"Authorization",
		"Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTc5ODU3NjEsImlhdCI6MTc1Nzg5OTM2MSwibmJmIjoxNzU3ODk5MzYxLCJzdWIiOiI1In0.kMeckM2m6bWeiF73_NWemGpKrRo1HZ2Orq7FVTqrAyF-nYKOaPTDqyjLpUE9I-zp1hG5zbaQQ7RK9d8ZaqWuZw").
		Expect().Status(http.StatusOK).JSON().Decode(&user)

	printJson(user)
}

func TestUserList(t *testing.T) {
	e := tester(t)
	var user dto.UserInfo
	e.GET(baseAPI+"/users").WithHeader(
		"Authorization",
		authToken).
		Expect().Status(http.StatusOK).JSON().Decode(&user)

	printJson(user)
}

func TestUser(t *testing.T) {
	e := tester(t)

	menuFormItem := dto.MenuForm{
		Code:        "user",
		Name:        "User management",
		Description: "User management",
		Sequence:    7,
		Type:        "page",
		Path:        "/system/user",
		Properties:  `{"icon":"user"}`,
		Status:      entity.MenuStatusEnabled,
	}

	var menu entity.Menu
	e.POST(baseAPI + "/menus").WithJSON(menuFormItem).
		Expect().Status(http.StatusOK).JSON().Decode(&util.ResponseResult{Data: &menu})

	assert := Assert.New(t)
	assert.NotEmpty(menu.ID)
	assert.Equal(menuFormItem.Code, menu.Code)
	assert.Equal(menuFormItem.Name, menu.Name)
	assert.Equal(menuFormItem.Description, menu.Description)
	assert.Equal(menuFormItem.Sequence, menu.Sequence)
	assert.Equal(menuFormItem.Type, menu.Type)
	assert.Equal(menuFormItem.Path, menu.Path)
	assert.Equal(menuFormItem.Properties, menu.Properties)
	assert.Equal(menuFormItem.Status, menu.Status)

	menuIDStr := comm.IDToStr(menu.ID)

	roleFormItem := dto.RoleForm{
		Code: "user",
		Name: "Normal",
		Menus: entity.RoleMenus{
			{MenuID: menu.ID},
		},
		Description: "Normal",
		Sequence:    8,
		Status:      entity.RoleStatusEnabled,
	}

	var role entity.Role
	e.POST(baseAPI + "/roles").WithJSON(roleFormItem).Expect().Status(http.StatusOK).JSON().Decode(&util.ResponseResult{Data: &role})
	assert.NotEmpty(role.ID)
	assert.Equal(roleFormItem.Code, role.Code)
	assert.Equal(roleFormItem.Name, role.Name)
	assert.Equal(roleFormItem.Description, role.Description)
	assert.Equal(roleFormItem.Sequence, role.Sequence)
	assert.Equal(roleFormItem.Status, role.Status)
	assert.Equal(len(roleFormItem.Menus), len(role.Menus))

	roleIDStr := comm.IDToStr(role.ID)
	userFormItem := dto.UserForm{
		Username: "test",
		Name:     "Test",
		Password: hash.MD5String("test"),
		Phone:    "0720",
		Email:    "test@gmail.com",
		Remark:   "test user",
		Status:   entity.UserStatusActivated,
		Roles:    entity.UserRoles{{RoleID: role.ID}},
	}

	var user entity.User
	e.POST(baseAPI + "/users").WithJSON(userFormItem).Expect().Status(http.StatusOK).JSON().Decode(&util.ResponseResult{Data: &user})
	assert.NotEmpty(user.ID)
	assert.Equal(userFormItem.Username, user.Username)
	assert.Equal(userFormItem.Name, user.Name)
	assert.Equal(userFormItem.Phone, user.Phone)
	assert.Equal(userFormItem.Email, user.Email)
	assert.Equal(userFormItem.Remark, user.Remark)
	assert.Equal(userFormItem.Status, user.Status)
	assert.Equal(len(userFormItem.Roles), len(user.Roles))

	var users entity.Users
	e.GET(baseAPI+"/users").WithQuery("username", userFormItem.Username).Expect().Status(http.StatusOK).JSON().Decode(&util.ResponseResult{Data: &users})
	assert.GreaterOrEqual(len(users), 1)

	newName := "Test 1"
	newStatus := entity.UserStatusFreezed
	user.Name = newName
	user.Status = newStatus
	userIDStr := comm.IDToStr(user.ID)
	e.PUT(baseAPI + "/users/" + userIDStr).WithJSON(user).Expect().Status(http.StatusOK)

	var getUser entity.User
	e.GET(baseAPI + "/users/" + userIDStr).Expect().Status(http.StatusOK).JSON().Decode(&util.ResponseResult{Data: &getUser})
	assert.Equal(newName, getUser.Name)
	assert.Equal(newStatus, getUser.Status)

	e.DELETE(baseAPI + "/users/" + userIDStr).Expect().Status(http.StatusOK)
	e.GET(baseAPI + "/users/" + userIDStr).Expect().Status(http.StatusNotFound)

	e.DELETE(baseAPI + "/roles/" + roleIDStr).Expect().Status(http.StatusOK)
	e.GET(baseAPI + "/roles/" + roleIDStr).Expect().Status(http.StatusNotFound)

	e.DELETE(baseAPI + "/menus/" + menuIDStr).Expect().Status(http.StatusOK)
	e.GET(baseAPI + "/menus/" + menuIDStr).Expect().Status(http.StatusNotFound)
}
