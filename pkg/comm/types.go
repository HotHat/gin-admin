package comm

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ModsRegister interface {
	Init(ctx context.Context)
	Register(gin *gin.Engine)
	Release(ctx context.Context)
}
