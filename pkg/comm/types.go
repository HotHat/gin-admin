package comm

import (
	"context"

	"github.com/gin-gonic/gin"
)

type RouteRegister interface {
	Register(ctx context.Context, gin *gin.Engine)
}
