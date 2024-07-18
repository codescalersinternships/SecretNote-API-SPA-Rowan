package pkg

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	router   *gin.Engine
	dataBase DB
}

func NewApp() (App, error) {
	db, err := NewDB()
	app := App{router: gin.Default(), dataBase: db}
	app.registerRoutes()
	return app, err
}

func (app *App) registerRoutes() {
	app.router.POST("/signup", app.SignUp)
	app.router.POST("/login", app.Login)
	app.router.POST("/note", app.RequireAuth, app.createNote)
	app.router.GET("/note/:noteID", app.getNote) // because if someone has link then they were sent it by original user
	app.router.GET("/notes", app.RequireAuth, app.GetNotes)
	app.router.GET("/expiredNotes", app.RequireAuth, app.GetExpiredNotes)
}
func (app *App) Run() error {
	return app.router.Run()
}
