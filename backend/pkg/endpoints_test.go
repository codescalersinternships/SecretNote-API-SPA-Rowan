package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_SignUp(t *testing.T) {
	router := gin.Default()
	db, _ := NewDB("test.db")
	app := App{router: gin.Default(), dataBase: db}
	user := dummyUser{
		Username: "rowan",
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
		Content: "I DON'T LIKE FISH",
	}
	noteToSend, _ := json.Marshal(note)
	request, _ := http.NewRequest(http.MethodPost, "/note", bytes.NewBuffer(noteToSend))
	// request.URL.Query().Add("id", fmt.Sprint(1))
	response := httptest.NewRecorder()
	router.POST("/note", app.CreateNote)
	router.ServeHTTP(response, request)
	body, _ := io.ReadAll(response.Body)
	var noteCreated dummyNote
	err := json.Unmarshal(body, &noteCreated)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, http.StatusAccepted, response.Code)
	// assert.Equal(t, dummyNote.Title, noteCreated.Title)
	fmt.Println(noteCreated)
}

func Test_GetNote(t *testing.T) {
	router := gin.Default()
	db, _ := NewDB("test.db")
	app := App{router: gin.Default(), dataBase: db}
	request, _ := http.NewRequest(http.MethodPost, "/note", nil)
	response := httptest.NewRecorder()
	router.GET("/note/987654345678", app.GetNote) // non existing note
	router.ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code)
}
