package service

import "github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/service"

type ServiceTest struct {
	AuthService *service.AuthService
	userService *service.UserService
}
