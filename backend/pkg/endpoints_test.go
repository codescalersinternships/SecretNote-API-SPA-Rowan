package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func Test_SignUp(t *testing.T) {
	router := gin.Default()
	user := dummyUser{
		Username: "rowan",
		Password: "slays",
	}
	userToSend, _ := json.Marshal(user)
	request, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(userToSend))
	response := httptest.NewRecorder()
	router.POST("/signup", SignUp)
	router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusAccepted, response.Code)
}

func Test_Login(t *testing.T) {
	router := gin.Default()
	user := dummyUser{
		Username: "rowan",
		Password: "slays",
	}
	userToSend, _ := json.Marshal(user)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userToSend))
	response := httptest.NewRecorder()
	router.POST("/login", Login)
	router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusAccepted, response.Code)
}

func SignUp(c *gin.Context) {
	var dummyuser dummyUser
	fmt.Println(dummyuser)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dummyuser.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
		return
	}
	dummyuser.Password = string(hashedPassword)
	c.Status(http.StatusAccepted)
}

func Login(c *gin.Context) {
	var dummyuser dummyUser
	if err := c.ShouldBindJSON(&dummyuser); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}

	actualUser := dummyUser{
		Username: "rowan",
		Password: "slays",
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(actualUser.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
		return
	}
	actualUser.Password = string(hashedPassword)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	err = bcrypt.CompareHashAndPassword([]byte(actualUser.Password), []byte(dummyuser.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username or password",
		})
		return
	}
	// tokenString, err := app.CreateJWTCookie(actualUser, c)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "failed to create token",
	// 	})
	// 	return
	// }
	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)
	c.Status(http.StatusAccepted)
}
