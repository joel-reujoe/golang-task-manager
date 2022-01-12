package taskmanager

import (
	// "fmt"
	"fmt"
	"net/http"
	taskmanager "task-manager/DTO"
	taskmanagerModel "task-manager/Models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)





func GetAllUsers(c *gin.Context){
	



	if c.MustGet("userType") != taskmanagerModel.Admin {
		c.JSON(http.StatusUnauthorized, gin.H{"message":"Please login as admin"})
		return
	}

	dbConn, ok := c.MustGet("dbConn").(*gorm.DB)
	 if !ok {
		c.JSON(http.StatusInternalServerError,"Could not connect to db")
		return
	}

	if c.MustGet("userId").(uint64) > 0{
		result := []taskmanagerModel.User{}
		dbConn.Table("users").Where("id <> ?",c.MustGet("userId")).Find(&result)
		c.JSON(http.StatusOK,result)

	}else{
		c.JSON(http.StatusBadGateway,gin.H{"message":"Could not create task"})
		return
	}
}


func CreateUser(c *gin.Context){

	var user taskmanager.CreateUserDto

	if c.MustGet("userType") != taskmanagerModel.Admin {
		c.JSON(http.StatusUnauthorized, gin.H{"message":"Please login as admin"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusUnprocessableEntity,"Incorrect Parameters Provided")
		return
	}

	dbConn, ok := c.MustGet("dbConn").(*gorm.DB)
	 if !ok {
		c.JSON(http.StatusInternalServerError,"Could not connect to db")
		return
	}

	if c.MustGet("userId").(uint64) > 0{
		
		
		userModel := &taskmanagerModel.User{Email:user.Email,FirstName: user.FirstName, LastName: user.LastName, Password: user.Password}

		if user.IsReviewer {
			userModel.UserType = taskmanagerModel.Reviewer
		}else{
			userModel.UserType = taskmanagerModel.Other
		}

		dbConn.Create(userModel)
		c.JSON(http.StatusOK,gin.H{"message":fmt.Sprintf("User created with id %d",userModel.Id)})
		return
	}else{
		c.JSON(http.StatusBadGateway,gin.H{"message":"Could not create task"})
		return
	}
}


func GetUserById(c *gin.Context){

	var user taskmanager.GetUserByIdDto
	fmt.Println(c.Keys)
	if c.MustGet("userType") != taskmanagerModel.Admin {
		c.JSON(http.StatusUnauthorized, gin.H{"message":"Please login as admin"})
		return
	}

	if err := c.ShouldBindUri(&user); err != nil{
		c.JSON(http.StatusUnprocessableEntity,"Incorrect Parameters Provided")
		return
	}

	dbConn, ok := c.MustGet("dbConn").(*gorm.DB)
	 if !ok {
		c.JSON(http.StatusInternalServerError,"Could not connect to db")
		return
	}

	if c.MustGet("userId").(uint64) > 0{
		var userModel taskmanagerModel.User
		dbConn.Where("id = ?", user.UserId).First(&userModel)

		userResponseDto := taskmanager.GetUserByIdResponseDto{}
		userResponseDto.FirstName = userModel.FirstName
		userResponseDto.LastName = userModel.LastName
		userResponseDto.Email = userModel.Email

		if userModel.UserType == taskmanagerModel.Reviewer{
			userResponseDto.IsReviewer = true
		}else{
			userResponseDto.IsReviewer = false
		}

		c.JSON(http.StatusOK,userResponseDto)
		return
	}else{
		c.JSON(http.StatusNotFound,gin.H{"message":"Please login again"})
	}
}

func GetTaskByUserId(c *gin.Context){
	var userId taskmanager.GetTaskByUserIdDto

	if c.MustGet("userType") != taskmanagerModel.Admin {
		c.JSON(http.StatusUnauthorized, gin.H{"message":"Please login as admin"})
		return
	}

	if err := c.ShouldBindUri(&userId); err != nil{
		c.JSON(http.StatusUnprocessableEntity,"Incorrect Parameters Provided")
		return
	}


	dbConn, ok := c.MustGet("dbConn").(*gorm.DB)
	 if !ok {
		c.JSON(http.StatusInternalServerError,"Could not connect to db")
		return
	}

	if c.MustGet("userId").(uint64) > 0{
		result := []taskmanagerModel.Task{}

		dbConn.Where("user_id = ?", userId.UserId).Find(&result)
		c.JSON(http.StatusOK, result)	
	}else{
		c.JSON(http.StatusNotFound,gin.H{"message":"Please login again"})
	}


}

func AssignReviewerToUser(c *gin.Context){

	var assignRevierDto taskmanager.AssignTaskToReviewerDto
	if c.MustGet("userType") != taskmanagerModel.Admin {
		c.JSON(http.StatusUnauthorized, gin.H{"message":"Please login as admin"})
		return
	}

	if err := c.ShouldBindJSON(&assignRevierDto); err != nil{
		c.JSON(http.StatusUnprocessableEntity,"Incorrect Parameters Provided")
		return
	}


	dbConn, ok := c.MustGet("dbConn").(*gorm.DB)
	 if !ok {
		c.JSON(http.StatusInternalServerError,"Could not connect to db")
		return
	}

	if c.MustGet("userId").(uint64) > 0{

		reviewerTasks := []taskmanagerModel.ReviewersTask{}

		var reviewers []taskmanagerModel.User

		dbConn.Find(&reviewers, assignRevierDto.ReviewerId)
		for i := 0; i<len(assignRevierDto.ReviewerId);i++{
				if reviewers[i].UserType == taskmanagerModel.Reviewer{
				reviewerTask := taskmanagerModel.ReviewersTask{
					UserId: assignRevierDto.UserId,
					ReviewerId: assignRevierDto.ReviewerId[i],
				}

				reviewerTasks = append(reviewerTasks,reviewerTask)
			}
		}
		dbConn.Create(reviewerTasks)
		c.JSON(http.StatusOK, gin.H{"message":"Reviewers assigned to user"})
		return
		
	}else{
		c.JSON(http.StatusNotFound,gin.H{"message":"Please login again"})
		return
	}
}