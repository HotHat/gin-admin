package service

import (
	"fmt"
	"os"
	"testing"

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
