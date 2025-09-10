package test

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/HotHat/gin-admin/v10/internal/config"
)

var TestContext context.Context

func SetConfig() {
	var workDirFlag = flag.String("workDir", "", "The working directory.")
	flag.Parse()

	fmt.Println("Work dir:", *workDirFlag, flag.Args())

	workDir := *workDirFlag
	staticDir := ""
	config.MustLoad(workDir, strings.Split("dev", ",")...)
	config.C.General.WorkDir = workDir
	config.C.Middleware.Static.Dir = staticDir
	//config.C.Print()
	config.C.PreLoad()

	TestContext = context.Background()

}
