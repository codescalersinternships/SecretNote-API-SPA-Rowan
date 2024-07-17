package pkg

import (
	"github.com/gin-gonic/gin"
	//  "github.com/golang-jwt/jwt"
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
	app.router.POST("/note/", app.createNote)
	app.router.GET("/note/:noteID", app.getNote) // still handle id retreieving
	app.router.POST("/signup", app.SignUp)
	app.router.GET("/validate", app.RequireAuth, app.Validate)
	app.router.POST("/login", app.Login)
}
func (app *App) Run() error {
	return app.router.Run()
}
