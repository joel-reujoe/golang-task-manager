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


func GetAssignedUsers(c *gin.Context){


	if c.MustGet("userType") != taskmanagerModel.Reviewer {
		c.JSON(http.StatusUnauthorized, gin.H{"message":"Please login as reviewer"})
		return
	}


	dbConn, ok := c.MustGet("dbConn").(*gorm.DB)
	 if !ok {
		c.JSON(http.StatusInternalServerError,"Could not connect to db")
		return
	}

	userId := c.MustGet("userId").(uint64)
	if userId > 0{
		var reviewers []taskmanagerModel.ReviewersTask
		var userIds []int

		dbConn.Where("reviewer_id = ?", userId).Find(&reviewers)
		for i := 0; i< len(reviewers);i++ {
			userIds = append(userIds, int(reviewers[i].UserId))
		}

		var users []taskmanagerModel.User
		dbConn.Find(&users, userIds)
		c.JSON(http.StatusOK, users)

	}else{
		c.JSON(http.StatusNotFound,gin.H{"message":"Please login again"})
		return
	}

}


func GetReviewerTaskByUserId(c *gin.Context){


	var userIdObj taskmanager.GetReviewerTaskByUserId

	if c.MustGet("userType") != taskmanagerModel.Reviewer {
		c.JSON(http.StatusUnauthorized, gin.H{"message":"Please login as reviewer"})
		return
	}

	if err := c.ShouldBindUri(&userIdObj); err != nil{
		c.JSON(http.StatusUnprocessableEntity,"Incorrect Parameters Provided")
		return
	}

	dbConn, ok := c.MustGet("dbConn").(*gorm.DB)
	 if !ok {
		c.JSON(http.StatusInternalServerError,"Could not connect to db")
		return
	}

	userId := c.MustGet("userId").(uint64)
	if userId > 0{
		tasks := []taskmanagerModel.Task{}
		reviewerTask := []taskmanagerModel.ReviewersTask{}

		fmt.Println(userIdObj.UserId)
		dbConn.Where("reviewer_id = ?", userId).Find(&reviewerTask)
		result := dbConn.Where("user_id = ?", userIdObj.UserId).Find(&tasks)
		fmt.Println(tasks)

		if result.RowsAffected > 0 {
			dbConn.Find(&tasks,userIdObj.UserId)
			c.JSON(http.StatusOK,tasks)
			return
		}else{
			c.JSON(http.StatusNotFound, gin.H{"message":"No results found"})
		}
	}else{
		c.JSON(http.StatusNotFound,gin.H{"message":"Please login again"})
		return
	}

}

func ApproveTask(c *gin.Context){
	var approveTaskDto taskmanager.ApproveTaskDto

	if err := c.ShouldBindJSON(&approveTaskDto); err != nil{
		c.JSON(http.StatusUnprocessableEntity,"Incorrect Parameters Provided")
		return
	}

	dbConn, ok := c.MustGet("dbConn").(*gorm.DB)
	 if !ok {
		c.JSON(http.StatusInternalServerError,"Could not connect to db")
		return
	}


	userId := c.MustGet("userId").(uint64)
	if userId > 0{

		approvedTask := taskmanagerModel.Task{Approved: true, ApprovedBy: uint32(userId)}

		dbConn.Model(&taskmanagerModel.Task{}).Where("id = ?", approveTaskDto.TaskId).Updates(approvedTask)
		c.JSON(http.StatusNotFound,gin.H{"message":fmt.Sprintf("Task with id %d is approved",approveTaskDto.TaskId)})
		return
	}else{
		c.JSON(http.StatusNotFound,gin.H{"message":"Please login again"})
		return
	}

}