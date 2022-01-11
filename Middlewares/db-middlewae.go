package taskmanager

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func DbMiddleWare(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbConn",db)
		c.Next()
	}
}
