package taskmanager

import (
	"fmt"
	"net/http"
	taskmanagerLogin "task-manager/Utility"

	"github.com/gin-gonic/gin"
)





func VerifyToken() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenAuth, err := taskmanagerLogin.ExtractTokenMetadata(c.Request)
		fmt.Println(err)
		if err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{"message":"Incorrect token"})
			return
		}
		c.Set("userId",tokenAuth.UserId)
		c.Set("userType",tokenAuth.UserType)
	}
}