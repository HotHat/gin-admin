package service

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/entity"
	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/test"
)

var serviceTest *ServiceTest

func TestMain(m *testing.M) {
	// --- SETUP ---
	fmt.Println("Setting up test suite...")
	// For example, connect to a database
	//setConfig()
	test.SetConfig()
	serviceTest, _, _ = BuildService(test.TestContext)

	// --- RUN TESTS ---
	code := m.Run()

	// --- TEARDOWN ---
	fmt.Println("Tearing down test suite...")
	// For example, close the database connection

	os.Exit(code)
}

// A simple test to demonstrate
func TestUserGet(t *testing.T) {
	user, err := serviceTest.userService.Get(test.TestContext, 1)
	if err != nil {
		fmt.Println("Error: ", err)

	}
	fmt.Println(user)
}

func TestCreatesUser(t *testing.T) {

	userForm := dto.UserForm{
		Username: "hello",
		Password: "123456",
		Name:     "world",
		Phone:    "13800138000",
		Email:    "aaa",
		Remark:   "",
		Status:   1,
		Roles: entity.UserRoles{
			{
				RoleID:    1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				RoleName:  "haha",
			},
		},
	}
	err := userForm.Validate()
	if err != nil {
		t.Error(err)
	}

	user, err := serviceTest.userService.Create(test.TestContext, &userForm)

	if err != nil {
		//t.Errorf("Error:", err)
		fmt.Println("Error: ", err)
	}
	fmt.Println(user)
}
