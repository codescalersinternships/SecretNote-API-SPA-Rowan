package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (app *App) createNote(c *gin.Context) {
	var dummynote dummyNote
	if err := c.ShouldBindJSON(&dummynote); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Println(dummynote)
	newNote := Note{Title: dummynote.Title, Content: dummynote.Content}
	if err := app.dataBase.CreateNote(newNote); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, newNote)
}
func (app *App) getNote(c *gin.Context) {
	noteID := c.Param("noteID")
	uuID, err := uuid.Parse(noteID)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}

	note, err := app.dataBase.GetNote(uuID)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, note)
}
func (app *App) SignUp(c *gin.Context) {
	var dummyuser dummyUser
	if err := c.ShouldBindJSON(&dummyuser); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Println(dummyuser)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dummyuser.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
		return
	}
	dummyuser.Password = string(hashedPassword)
	if err := app.dataBase.CreateUser(dummyuser.Username, dummyuser.Password); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.Status(http.StatusAccepted)
}

func (app *App) Login(c *gin.Context) {
	var dummyuser dummyUser
	if err := c.ShouldBindJSON(&dummyuser); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Println(dummyuser)

	actualUser, err := app.dataBase.GetUserByUsername(dummyuser.Username)
	fmt.Println(actualUser)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	if actualUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username or password",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(actualUser.Password), []byte(dummyuser.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	tokenString, err := app.CreateJWTCookie(actualUser, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)
	c.Status(http.StatusAccepted)
}

type dummyNote struct {
	Title   string
	Content string
}
type dummyUser struct {
	Username string
	Password string
}
