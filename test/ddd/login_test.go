package ddd

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/HotHat/gin-admin/v10/internal/ddd/rbac/dto"
	"github.com/HotHat/gin-admin/v10/pkg/util"
)

func TestGetCaptcha(t *testing.T) {
	e := tester(t)

	var captcha dto.Captcha
	fmt.Println("captcha:", captcha)

	e.GET(baseAPI + "/captcha/id").
		Expect().Status(http.StatusOK).JSON().Decode(&util.ResponseResult{Data: &captcha})

	fmt.Println("GetCaptcha:", captcha)

	//url := fmt.Sprintf(baseAPI+"/captcha/image?id=%s&reload=1", captcha.CaptchaID)
	e.GET(baseAPI+"/captcha/image").
		WithQuery("id", captcha.CaptchaID).
		WithQuery("reload", "1").
		Expect().Status(http.StatusOK).JSON().Decode(&util.ResponseResult{Data: &captcha})

}

func TestGetCaptchaImage(t *testing.T) {
	e := tester(t)

	captchaID := "abc123"
	fmt.Println("captcha:", captchaID)
	//url := fmt.Sprintf(baseAPI+"/captcha/image?id=%s&reload=1", captcha.CaptchaID)
	e.GET(baseAPI+"/captcha/image").
		WithQuery("id", captchaID).
		WithQuery("reload", "1").
		Expect().Status(http.StatusOK)

}
