package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// CreateNote handles "/note" endpoint
// requires authentication
func (app *App) CreateNote(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var dummyNote dummyNote
	if err := c.ShouldBindJSON(&dummyNote); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Println(dummyNote)
	dummyUser, _ := c.Get("user")
	fmt.Println(dummyUser)
	fmt.Println("hellooo")
	user, err := app.dataBase.GetUserByUsername(dummyUser.(User).Username)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	newNote := Note{Title: dummyNote.Title, Content: dummyNote.Content}
	note, err := app.dataBase.CreateNote(newNote, user)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, note)
}

// GetNote handles "/note/:noteID" endpoint
// doesn't requires authentication because it's fetched by its uuid which user shares
func (app *App) GetNote(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	noteID := c.Param("noteID")
	fmt.Println(noteID)
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
	// if uint(note.ID) == 0 {

	// }
	c.JSON(http.StatusOK, note)
}

// GetNotes handles "/notes" endpoint
// requires authentication
func (app *App) GetNotes(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	dummyUser, _ := c.Get("user")
	notes, err := app.dataBase.GetNotes(dummyUser.(User))
	fmt.Println(notes)
	// db returns errors if null found / empty
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, notes)
}

// GetExpiredNotes handles "/expiredNotes" endpoint
// requires authentication
func (app *App) GetExpiredNotes(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	dummyUser, _ := c.Get("user")
	notes, _ := app.dataBase.GetExpiredNotes(dummyUser.(User))
	// if err != nil {
	// 	c.Error(err)
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// }
	c.JSON(http.StatusOK, notes)
}

// SignUp handles "/signup" endpoint
// hashes user password before saving it to db
// handles Authentication
func (app *App) SignUp(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
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

// Login handles "/login" endpoint
// handles Authorization of user
// Sets up JWT in cookies of browser
func (app *App) Login(c *gin.Context) {
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
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
	c.SetCookie("Authorization", tokenString, 3600, "/", "localhost", false, true)
	// c.Status(http.StatusAccepted)
	// c.SetSameSite(http.SameSiteNoneMode)
	c.JSON(http.StatusAccepted, actualUser)
}

type dummyNote struct {
	Title   string
	Content string
}
type dummyUser struct {
	Username string
	Password string
}
