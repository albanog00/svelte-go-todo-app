package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"todoapp.com/server/internal/models"
)

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func PostUser(c *gin.Context) {
	var user CreateUserDTO
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	newUser := &models.User{
		Id:       uuid.NewString(),
		Username: user.Username,
		Password: user.Password,
	}

	_, err := models.CreateUser(newUser)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, user)
}
