package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	// "time"
)

func getNote(w http.ResponseWriter, req *http.Request) {
	_, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("error detected while reading request body")
		os.Exit(1)
	}
	// str := string(body)
	// fmt.Println(string(body))
	fmt.Fprint(w, "helloooooooooo")
}

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func createNote(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprint(w, "Error method of '/createNote' request should be POST!")
		return
	}
	str, _ := io.ReadAll(req.Body)
	var note Note
	// fmt.Println(string(str))
	// err := json.NewDecoder(req.Body).Decode(&note)
	err := json.Unmarshal(str, &note)
	if err != nil {
		fmt.Println("Error decoding")
	}
	fmt.Println(note)
	// fmt.Println(note.Id)
	// fmt.Println(note.Title)
}

func main() {
	http.HandleFunc("/getNote", getNote)
	http.HandleFunc("/createNote", createNote)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error detected while reading request body")
		os.Exit(1)
	}
}
