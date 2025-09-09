package comm

import (
	"github.com/gin-gonic/gin"
)

type RouteRegister interface {
	Register(gin *gin.Engine)
}
