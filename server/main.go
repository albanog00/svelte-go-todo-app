package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"todoapp.com/server/api"
	"todoapp.com/server/internal/models"
)

func LoadEnv() {
	var err error

	if gin.Mode() == gin.ReleaseMode {
		err = godotenv.Load("prod.env")
	} else {
		err = godotenv.Load(".env")
	}

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func init() {
	// gin.SetMode(gin.ReleaseMode)

	LoadEnv()
	models.NewMySQLClient()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/tasks", api.GetTasks)
	r.POST("/tasks", api.PostTask)
	r.PUT("/tasks/:id", api.DeleteTask)

	r.POST("/users", api.PostUser)

	r.Run("localhost:3001")
}
