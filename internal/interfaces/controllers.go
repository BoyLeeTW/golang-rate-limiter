package interfaces

import "github.com/gin-gonic/gin"

type NormalControllerInterface interface {
	Get(c *gin.Context)
}
