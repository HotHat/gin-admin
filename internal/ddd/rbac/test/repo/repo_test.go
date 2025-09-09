package repo

import (
	"fmt"
	"os"
	"testing"

	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/test"
	"github.com/LyricTian/gin-admin/v10/pkg/util"
)

var repoTest *RepoTest

func TestMain(m *testing.M) {
	// --- SETUP ---
	fmt.Println("Setting up test suite...")
	// For example, connect to a database
	//setConfig()
	test.SetConfig()
	repoTest, _, _ = BuildRepo(test.TestContext)

	// --- RUN TESTS ---
	code := m.Run()

	// --- TEARDOWN ---
	fmt.Println("Tearing down test suite...")
	// For example, close the database connection

	os.Exit(code)
}

// A simple test to demonstrate
func TestUserQuery(t *testing.T) {
	queryResult, _ := repoTest.UserRepo.Query(test.TestContext, dto.UserQueryParam{
		PaginationParam: util.PaginationParam{
			Pagination: true,
			OnlyCount:  false,
			Current:    1,
			PageSize:   15,
		},
		LikeUsername: "", // Username for login
		LikeName:     "", // Name of user
		Status:       0,
	})

	fmt.Println(queryResult.PageResult.Total, queryResult.PageResult.Current, queryResult.PageResult.PageSize)

	for i, user := range queryResult.Data {
		fmt.Println(i, user)
	}
}
