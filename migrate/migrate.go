package main

import (
	"github.com/mohanrajnambe/task-manager-golang/initializers"
	"github.com/mohanrajnambe/task-manager-golang/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Task{})
}
