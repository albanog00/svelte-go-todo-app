package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"todoapp.com/server/api"
	"todoapp.com/server/internal/models"
)

func init() {
	models.NewMySQLClient()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/tasks", api.GetTasks)
	r.POST("/tasks", api.PostTask)
	r.PUT("/tasks/:id", api.DeleteTask)

	r.Run("localhost:3001")
}
