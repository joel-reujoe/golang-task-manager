package taskmanager

import (
	// "fmt"
	"net/http"
	"task-manager/Models"
	Utility "task-manager/Utility"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



func Login(c *gin.Context){
	var u taskmanager.User

	if err := c.ShouldBindJSON(&u); err != nil{
		c.JSON(http.StatusUnprocessableEntity,"Incorrect Parameters Provided")
		return
	}


	if u.Email == "" || u.Password == "" {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	 }

	 dbConn, ok := c.MustGet("dbConn").(*gorm.DB)

	 if !ok {
		 c.JSON(http.StatusInternalServerError,"Could not connect to db")
	 }


	 var user taskmanager.User
	 dbConn.First(&user, "email = ? and password = ?", u.Email, u.Password)
	 token, err := Utility.CreateToken(user.Id,user.UserType)

	 if err != nil {
		c.JSON(http.StatusNotFound,"No user exists")
	 }
   
	  if token != ""{
		  c.JSON(http.StatusOK,token)
	  } 
}

