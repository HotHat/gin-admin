package repo

import "github.com/HotHat/gin-admin/v10/internal/ddd/rbac/repo"

type RepoTest struct {
	MenuRepo         *repo.MenuRepo
	MenuResourceRepo repo.MenuResourceRepo
	UserRepo         *repo.UserRepo
	UserRoleRepo     repo.UserRoleRepo
	RoleRepo         *repo.RoleRepo
	RoleMenuRepo     *repo.RoleMenuRepo
}
