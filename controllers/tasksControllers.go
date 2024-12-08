package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohanrajnambe/task-manager-golang/initializers"
	"github.com/mohanrajnambe/task-manager-golang/models"
)

func TasksCreate(c *gin.Context) {

	var body struct {
		Title       string
		Description string
		Status      string
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"data":  nil,
			"error": err,
		})
	}
	fmt.Printf("%v/n", body)
	task := models.Task{Title: body.Title, Description: body.Description, Status: body.Status}

	result := initializers.DB.Create(&task) // pass pointer of data to Create

	if result.Error != nil {
		//not the right way to return a error
		c.JSON(400, gin.H{
			"data":  nil,
			"error": result.Error,
		})
	}

	c.JSON(200, gin.H{
		"data":  task,
		"error": nil,
	})
}

func TasksGetAll(c *gin.Context) {

	var tasks []models.Task
	initializers.DB.Find(&tasks)

	c.JSON(200, gin.H{
		"data":  tasks,
		"error": nil,
	})
}

func TaskGetByID(c *gin.Context) {

	id := c.Param("id")
	var task models.Task
	initializers.DB.First(&task, id)

	c.JSON(200, gin.H{
		"data":  task,
		"error": nil,
	})
}
func TaskUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Status      *string `json:"status"`
	}

	// Parse JSON body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}

	// Fetch the existing task
	var task models.Task
	if err := initializers.DB.First(&task, id).Error; err != nil {
		c.JSON(404, gin.H{
			"data":  nil,
			"error": "Task Not Found",
		})
		return
	}

	// Prepare updated fields
	updatedTask := map[string]interface{}{}
	if body.Title != nil {
		updatedTask["title"] = *body.Title
	}
	if body.Description != nil {
		updatedTask["description"] = *body.Description
	}
	if body.Status != nil {
		updatedTask["status"] = *body.Status
	}

	// Apply updates
	initializers.DB.Model(&task).Updates(updatedTask)

	// Return the updated task
	c.JSON(200, gin.H{
		"data":  task,
		"error": nil,
	})
}

func TaskDelete(c *gin.Context) {

	id := c.Param("id")

	initializers.DB.Delete(&models.Task{}, id)

	c.JSON(200, gin.H{
		"data":  "Delete success",
		"error": nil,
	})
}
