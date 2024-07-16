package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewNote() Note {
	return Note{}
}
func getNote(c *gin.Context) {
	_, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("error detected while reading request body")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("error marshalling")
	}
	c.JSON(http.StatusOK, gin.H{
		"id":      6,
		"title":   "hello world",
		"content": "cute",
	})
}

func getNotes(c *gin.Context) {
	_, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("error detected while reading request body")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println("error marshalling")
	}
	var notes []Note
	notes = append(notes, Note{50, "sheer", "cups"})
	notes = append(notes, Note{18, "pizza", "ranch"})
	notes = append(notes, Note{16, "Lamin", "Yamal"})
	c.JSON(http.StatusOK, notes)
}

func createNote(c *gin.Context) {
	var newNote Note
	if err := c.ShouldBindJSON(&newNote); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Println(newNote)
	c.JSON(http.StatusOK, newNote)
}

func main() {
	router := gin.Default()
	router.GET("/getNote", getNote)
	router.GET("/getNotes", getNotes)
	router.POST("/createNote", createNote)
	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("error starting the server: %s\n", err)
		os.Exit(1)
	}
}
