package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"todoapp.com/server/api"
	"todoapp.com/server/internal/models"
)

func LoadEnv() {
	var err error

	if gin.IsDebugging() {
		err = godotenv.Load(".env")
	} else {
		err = godotenv.Load("prod.env")
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
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))

	r.GET("/tasks", api.GetTasks)
	r.POST("/tasks", api.PostTask)
	r.PUT("/tasks/:id", api.DeleteTask)

	r.POST("/users", api.PostUser)
	r.GET("/users", api.GetUserInfo)

	r.POST("/auth/signin", api.AuthUser)
	r.GET("/auth/validate", api.ValidateAuthUser)

	s := &http.Server{
		Addr:    ":3001",
		Handler: r,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
	}
}
