package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"todoapp.com/server/internal/models"
)

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
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	newTask.Id = uuid.NewString()
	task, err := models.CreateTask(&newTask)
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
