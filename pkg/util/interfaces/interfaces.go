package interfaces

import "github.com/gin-gonic/gin"

type subHandler interface {
	GetPageNCount(c *gin.Context) (page int, count int, ok bool)
	BindRequest(c *gin.Context, obj any) bool
}
