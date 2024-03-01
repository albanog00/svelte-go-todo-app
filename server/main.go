package main

import (
	"log"
	"net/http"
	"time"

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
	gin.SetMode(gin.ReleaseMode)

	LoadEnv()
	models.NewMySQLClient()
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:5173, http://localhost:5173"},
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           15 * time.Minute,
	}))

	r.GET("/auth/validate", api.ValidateAuthUser)
	r.POST("/auth/signin", api.AuthUser)
	r.GET("/auth/signout", api.SignOutUser)

	r.GET("/users", api.GetUserInfo)
	r.POST("/users", api.PostUser)

	r.GET("/tasks", api.GetTasks)
	r.POST("/tasks", api.PostTask)
	r.DELETE("/tasks/:id", api.DeleteTask)

	s := &http.Server{
		Addr:    ":3001",
		Handler: r,
	}

	if err := s.ListenAndServeTLS("app.localhost.com.crt", "app.localhost.com.key"); err != nil {
		log.Fatalf(err.Error())
	}
}
