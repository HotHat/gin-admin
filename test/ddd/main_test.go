package ddd

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/HotHat/gin-admin/v10/internal/config"
	"github.com/HotHat/gin-admin/v10/internal/ddd/wirex"
	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
)

const (
	baseAPI = "/api/v1"
)

var (
	app *gin.Engine
)

func TestMain(m *testing.M) {
	// --- SETUP ---
	fmt.Println("Setting up test suite...")
	SetConfig()

	// --- RUN TESTS ---
	code := m.Run()

	// --- TEARDOWN ---
	fmt.Println("Tearing down test suite...")
	// For example, close the database connection

	os.Exit(code)
}

func SetConfig() {
	var workDirFlag = flag.String("workDir", "", "The working directory.")
	flag.Parse()

	fmt.Println("Work dir:", *workDirFlag, flag.Args())

	//workDir := *workDirFlag
	workDir := "D:\\Programming\\gin-admin\\configs"
	staticDir := ""
	config.MustLoad(workDir, strings.Split("dev", ",")...)
	config.C.General.WorkDir = workDir
	config.C.Middleware.Static.Dir = staticDir
	//config.C.Print()
	config.C.PreLoad()
	ctx := context.Background()
	injector, _, err := wirex.BuildInjector(ctx)
	if err != nil {
		panic(err)
	}

	app = gin.New()
	injector.Register(ctx, app)

}

func tester(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(app),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

func printJson(data interface{}) {
	s, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(s))
}
