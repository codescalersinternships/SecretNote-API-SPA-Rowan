package pkg

import (
	"github.com/gin-gonic/gin"
)

// App struct holds main logic of our backend and ORM together
// It has 2 instances: gin router and sqlite db
type App struct {
	router   *gin.Engine
	dataBase DB
}

// NewApp initializes instance of App
// NewApp calls NewDB which by default initializes db and auto migrates
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
	app.router.GET("/note/:noteID", app.GetNote) // because if someone has link then they were sent it by original user
	app.router.GET("/notes", app.RequireAuth, app.GetNotes)
	app.router.GET("/expiredNotes", app.RequireAuth, app.GetExpiredNotes)
}

// Run function runs the router instance
func (app *App) Run() error {
	return app.router.Run()
}
