package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type App struct {
	router   *gin.Engine
	dataBase DB
}

func NewApp() *App {
	app := App{router: gin.Default(), dataBase: *NewDB()}
	return &app
}

func (app *App) registerRoutes() {
	app.router.POST("/note/", app.createNote)
	app.router.GET("/note/:noteID", app.getNote) // still handle id retreieving
	app.router.POST("/signup", app.SignUp)
}

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

	if err := app.dataBase.CreateUser(dummyuser.Username, dummyuser.Password); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	return
}

type dummyNote struct {
	Title   string
	Content string
}
type dummyUser struct {
	Username string
	Password string
}
