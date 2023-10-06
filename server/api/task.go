package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"todoapp.com/server/internal/models"
)

type CreateTaskDTO struct {
	Description string
	Date        time.Time
}

func GetTasks(c *gin.Context) {
	var pageNum uint = 0
	pageString := c.Query("page")

	if pageString != "" {
		if parsedString, err := strconv.ParseUint(pageString, 10, 32); err != nil {
			pageNum = 0
		} else {
			pageNum = uint(parsedString)
		}
	}

	tasks, count, err := models.GetTasks(pageNum)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"tasks": tasks,
		"count": count,
	})
}

func PostTask(c *gin.Context) {
	var newTask CreateTaskDTO
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	task := &models.Task{
		Id:          uuid.NewString(),
		Description: newTask.Description,
		Date:        newTask.Date,
	}

	task, err := models.CreateTask(task)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := models.DeleteTask(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "task not found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "task deleted successfully",
	})
}
