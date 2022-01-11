package taskmanager

import (
	// "fmt"
	"net/http"
	taskmanager "task-manager/DTO"
	taskmanagerModel "task-manager/Models"
	taskmanagerLogin "task-manager/Utility"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTask(c *gin.Context){
	var task taskmanager.CreateTaskDto

	if err := c.ShouldBindJSON(&task); err != nil{
		c.JSON(http.StatusUnprocessableEntity,"Incorrect Parameters Provided")
		return
	}

	tokenAuth, err := taskmanagerLogin.ExtractTokenMetadata(c.Request)
	if err != nil{
		c.JSON(http.StatusForbidden, "Incorrect Token")
		return

	}
	if task.Title == "" || task.Description == "" {
		c.JSON(http.StatusUnauthorized, "Please provide valid task title and description")
		return
	 }

	 dbConn, ok := c.MustGet("dbConn").(*gorm.DB)
	 if !ok {
		c.JSON(http.StatusInternalServerError,"Could not connect to db")
		return
	}


	if uint32(tokenAuth.UserId) > 0{
		taskToCreate := &taskmanagerModel.Task{Title: task.Title, Description: task.Description, UserId: uint32(tokenAuth.UserId)}
		dbConn.Create(taskToCreate)
		if taskToCreate.ID > 0{
			c.JSON(http.StatusOK,gin.H{"taskId":taskToCreate.ID,"message":"TaskCreated"})
			return
		}	
	}else{
		c.JSON(http.StatusBadGateway,gin.H{"message":"Could not create task"})
		return
	}

	 
}


func UpdateTask(c *gin.Context){
	 var task taskmanager.UpdateTaskDto

	 if err := c.ShouldBindJSON(&task); err != nil{
		c.JSON(http.StatusUnprocessableEntity,"Incorrect Parameters Provided")
		return
	}

	tokenAuth, err := taskmanagerLogin.ExtractTokenMetadata(c.Request)

	if err != nil{
		c.JSON(http.StatusForbidden, "Incorrect Token")

	}
	if task.TaskId == "" || task.Title == "" || task.Description == "" {
		c.JSON(http.StatusUnauthorized, "Please provide valid taskid, task title and description")
		return
	 }


	dbConn, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
	   c.JSON(http.StatusInternalServerError,"Could not connect to db")
   	}

	if uint32(tokenAuth.UserId) > 0{
		var taskToUpdate taskmanagerModel.Task
		dbConn.First(&taskToUpdate,"id = ?", task.TaskId)
		dbConn.Model(&taskToUpdate).Updates(taskmanagerModel.Task{Title:task.Title,Description:task.Description})
		c.JSON(http.StatusOK,gin.H{"message":"Task with id "+task.TaskId+"updated successfully"})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{"message":"Task with id "+task.TaskId+"updated failed"})
		return
	}

	
}

func GetTasksById(c *gin.Context){
	var task taskmanager.GetTasksByIdDto

	if err := c.ShouldBindUri(&task); err != nil{
		c.JSON(http.StatusUnprocessableEntity,"Incorrect Parameters Provided")
		return
	}

	tokenAuth, err := taskmanagerLogin.ExtractTokenMetadata(c.Request)

	if err != nil{
		c.JSON(http.StatusForbidden, "Incorrect Token")

	}
	if task.TaskId == "" {
		c.JSON(http.StatusUnauthorized, "Please provide valid taskid")
		return
	 }


	dbConn, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
	   c.JSON(http.StatusInternalServerError,"Could not connect to db")
   	}

	if uint32(tokenAuth.UserId) > 0{
		var taskToGet taskmanagerModel.Task
		var taskToSend taskmanager.TaskDetailsDto
		dbConn.First(&taskToGet, "id = ? and user_id = ?", task.TaskId, tokenAuth.UserId)
		taskToSend.TaskId = taskToGet.ID
		taskToSend.Title = taskToGet.Title
		taskToSend.Description = taskToGet.Description
		c.JSON(http.StatusOK,taskToSend)
		return
	}else{
		c.JSON(http.StatusOK,gin.H{"message":"Could not get task with id "+task.TaskId})
		return
	}


}