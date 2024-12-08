package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mohanrajnambe/task-manager-golang/controllers"
	"github.com/mohanrajnambe/task-manager-golang/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/tasks", controllers.TasksCreate)
	r.GET("/tasks", controllers.TasksGetAll)
	r.GET("/tasks/:id", controllers.TaskGetByID)
	r.PUT("/tasks/:id", controllers.TaskUpdate)
	r.DELETE("/tasks/:id", controllers.TaskDelete)
	r.Run()
}
