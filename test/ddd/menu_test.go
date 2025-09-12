package ddd

import (
	"net/http"
	"testing"
)

func TestUserMenus(t *testing.T) {
	e := tester(t)

	e.GET(baseAPI+"/user/menus").WithHeader(
		"Authorization",
		"Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTc3MzAxODQsImlhdCI6MTc1NzY0Mzc4NCwibmJmIjoxNzU3NjQzNzg0LCJzdWIiOiI1In0.lcmq-iwYdmwaIYYnLtQkcwZMv4LbpeP2LTlW_ZlhVX_t7oKXFsNPFE8V-nlW8xg5T66e1DeZRvUVItoNUmCPeQ").
		Expect().Status(http.StatusOK).JSON()

}
