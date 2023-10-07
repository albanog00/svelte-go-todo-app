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
			"message": err.Error(),
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
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": gin.H{
			"username": user.Username,
		},
	})
}

func GetUserInfo(c *gin.Context) {
	token, err := c.Cookie("auth-jwt")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	claims, err := RetrieveClaimsFromToken(token)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	username, err := claims.GetSubject()
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := models.GetUser(username)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Authorized",
		"data": gin.H{
			"username": user.Username,
		},
	})
}

func AuthUser(c *gin.Context) {
	var authUser AuthUserDto
	if err := c.BindJSON(&authUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := &models.User{
		Username: authUser.Username,
		Password: authUser.Password,
	}

	_, err := models.AuthUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	jwt, err := generateJWT(user)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.SetCookie("auth-jwt", jwt, 60*90, "/", "localhost", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"username": user.Username,
		},
	})
}

func ValidateAuthUser(c *gin.Context) {
	tokenString, err := c.Cookie("auth-jwt")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = validateJWT(tokenString)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "authorized",
	})
}

func SignOutUser(c *gin.Context) {
	c.SetCookie("auth-jwt", "", -1, "/", "", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
