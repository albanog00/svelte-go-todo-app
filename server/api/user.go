package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"todoapp.com/server/internal/models"
)

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func PostUser(c *gin.Context) {
	var newUser CreateUserDTO
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	user := &models.User{
		Username: newUser.Username,
		Password: newUser.Password,
	}

	task, err := models.CreateUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, task)
}
