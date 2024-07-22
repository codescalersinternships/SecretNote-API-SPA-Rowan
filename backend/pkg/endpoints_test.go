package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_SignUp(t *testing.T) {
	router := gin.Default()
	db, _ := NewDB("test.db")
	app := App{router: gin.Default(), dataBase: db}
	user := dummyUser{
		Username: "rue",
		Password: "slays",
	}
	userToSend, _ := json.Marshal(user)
	request, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(userToSend))
	response := httptest.NewRecorder()
	router.POST("/signup", app.SignUp)
	router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusAccepted, response.Code)
}

func Test_Login(t *testing.T) {
	router := gin.Default()
	db, _ := NewDB("test.db")
	app := App{router: gin.Default(), dataBase: db}
	user := dummyUser{
		Username: "rowan",
		Password: "slays",
	}
	userToSend, _ := json.Marshal(user)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userToSend))
	response := httptest.NewRecorder()
	router.POST("/login", app.Login)
	router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusAccepted, response.Code)
}

func Test_CreateNote(t *testing.T) {
	router := gin.Default()
	db, _ := NewDB("test.db")
	app := App{router: gin.Default(), dataBase: db}
	note := dummyNote{
		Title:   "important",
		Content: "I DON'T",
	}
	noteToSend, _ := json.Marshal(note)
	request, _ := http.NewRequest(http.MethodPost, "/note", bytes.NewBuffer(noteToSend))
	// request.URL.Query().Add("id", fmt.Sprint(1))
	response := httptest.NewRecorder()
	router.POST("/note", app.SetUserMiddleware, app.CreateNote)
	router.ServeHTTP(response, request)
	body, _ := io.ReadAll(response.Body)
	var noteCreated dummyNote
	err := json.Unmarshal(body, &noteCreated)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, note, noteCreated)
	fmt.Println(noteCreated)
}

func Test_GetNote(t *testing.T) {
	router := gin.Default()
	db, _ := NewDB("test.db")
	app := App{router: gin.Default(), dataBase: db}
	t.Run("non existing note", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/note/987654345678", nil)
		response := httptest.NewRecorder()
		router.GET("/note/987654345678", app.SetUserMiddleware, app.GetNote) // non existing note
		router.ServeHTTP(response, request)
		assert.NotEqual(t, 200, response.Code)
	})
	t.Run("existing note", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/note", nil)
		response := httptest.NewRecorder()
		router.GET("/note", app.SetPathParam, app.GetNote) // non existing note
		router.ServeHTTP(response, request)
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
		var receivedNote dummyNote
		err := json.Unmarshal(data, &receivedNote)
		if err != nil {
			fmt.Println(err)
		}
		uuID, _ := uuid.Parse("b7c4a1f9-4695-11ef-892b-e454e83d4f2b")
		noteExpected, _ := app.dataBase.GetNote(uuID)
		fmt.Println(noteExpected)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, noteExpected.Title, receivedNote.Title)
		assert.Equal(t, noteExpected.Content, receivedNote.Content)
	})
}

func (app *App) SetUserMiddleware(c *gin.Context) {
	user := User{}
	user.ID = 2
	c.Set("user", user)
	c.Next()
}

func (app *App) SetPathParam(c *gin.Context) {
	c.Params = []gin.Param{
		{
			Key:   "noteID",
			Value: "b7c4a1f9-4695-11ef-892b-e454e83d4f2b",
		},
	}
	c.Next()
}
