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
	app.router.POST("/note/:noteID", app.getNote) // still handle id retreieving
}

func (app *App) createNote(c *gin.Context) {
	var newNote Note
	if err := c.ShouldBindJSON(&newNote); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Println(newNote)
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
