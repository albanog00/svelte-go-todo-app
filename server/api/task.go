package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"todoapp.com/server/internal/models"
)

type CreateTaskDTO struct {
	Description string
	Date        string
	Time        string
}

func GetTasks(c *gin.Context) {
	tasks, err := models.GetTasks()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func PostTask(c *gin.Context) {
	var newTask CreateTaskDTO
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	task := &models.Task{
		Id:          uuid.NewString(),
		Description: newTask.Description,
		Date:        newTask.Date,
		Time:        newTask.Time,
	}

	task, err := models.CreateTask(task)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err,
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