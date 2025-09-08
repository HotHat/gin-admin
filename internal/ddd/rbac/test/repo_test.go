package test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/LyricTian/gin-admin/v10/internal/config"
	"github.com/LyricTian/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/LyricTian/gin-admin/v10/pkg/util"
)

var repoTest *RepoTest

var ctx context.Context

func setConfig() {
	var workDirFlag = flag.String("workDir", "", "The working directory.")
	flag.Parse()

	fmt.Println("Work dir:", *workDirFlag, flag.Args())

	workDir := *workDirFlag
	staticDir := ""
	config.MustLoad(workDir, strings.Split("dev", ",")...)
	config.C.General.WorkDir = workDir
	config.C.Middleware.Static.Dir = staticDir
	config.C.Print()
	config.C.PreLoad()

	ctx = context.Background()

	repoL, clearFun, err := BuildRepo(ctx)

	if err != nil {
		clearFun()
	}

	//
	repoTest = repoL
}

func TestMain(m *testing.M) {
	// --- SETUP ---
	fmt.Println("Setting up test suite...")
	// For example, connect to a database
	setConfig()

	// --- RUN TESTS ---
	code := m.Run()

	// --- TEARDOWN ---
	fmt.Println("Tearing down test suite...")
	// For example, close the database connection

	os.Exit(code)
}

// A simple test to demonstrate
func TestUserQuery(t *testing.T) {
	queryResult, _ := repoTest.UserRepo.Query(ctx, dto.UserQueryParam{
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
