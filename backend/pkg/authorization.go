package pkg

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var secret = "Pxyehdyrowans_security"

// CreateJWTCookie creates JWT by SigningMethodHS256
// JWT is created by secretkey relevant to each user
// JWT is stored in browser cookie for later retreival
func (app *App) CreateJWTCookie(user User, c *gin.Context) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		// "username": user.Username,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	secretKey := fmt.Sprint(user.ID) + secret
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create token",
		})
		return "", err
	}
	return tokenString, nil
}

// RequireAuth is our middleware fuction that adds User to the gin context
// Request then proceeds while having user set as a key on the request
func (app *App) RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		claims, _ := token.Claims.(jwt.MapClaims)
		id := uint(claims["id"].(float64))
		secretKey := fmt.Sprint(id) + secret
		return []byte(secretKey), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		user, err := app.dataBase.GetUserByID(uint(claims["id"].(float64)))
		if err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", user)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
}
