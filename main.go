package main

import (
	"log"
	"os"
	"task-manager/Controller"
	middlewares "task-manager/Middlewares"
	util "task-manager/Utility"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var accessSecret string
var dbConn *gorm.DB

var (
	router = gin.Default()
)

func main(){
 err := godotenv.Load(".env")

 if err != nil{
	log.Fatal("Error loading env file")
	
 }
	accessSecret = os.Getenv("ACCESS_SECRET")
	dbConn := util.ConnectToDb()
	router.Use(middlewares.DbMiddleWare(dbConn))
	router.POST("/login",taskmanager.Login)
 	router.Use(middlewares.VerifyToken())
	router.POST("/task",taskmanager.CreateTask)
	router.PUT("/task",taskmanager.UpdateTask)
	router.GET("/task/:taskId",taskmanager.GetTasksById)
	router.GET("/admin/user",taskmanager.GetAllUsers)
	router.POST("/admin/user",taskmanager.CreateUser)
	router.GET("/user/:userId",taskmanager.GetUserById)
	router.GET("/admin/:userId/tasks",taskmanager.GetTaskByUserId)

	router.Run(":8000")

}