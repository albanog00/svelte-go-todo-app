package api

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"todoapp.com/server/internal/models"
)

type AuthUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func generateJWT(user *models.User) (string, error) {
	jwtExpiresIn, err := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      "golang-auth-server",
		"sub":      user.Username,
		"iat":      time.Now().Unix(),
		"nbf":      time.Now().Unix(),
		"username": user.Username,
		"exp":      time.Now().Add(time.Duration(jwtExpiresIn) * time.Minute).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func validateJWT(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	return err
}

func ValidateAuthUser(c *gin.Context) {
	token, err := c.Cookie("auth-jwt")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := validateJWT(token); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{
		"message": "Authorized",
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

	c.SetCookie("auth-jwt", jwt, 60*60, "/", "localhost", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{
		"success": true,
	})
}
