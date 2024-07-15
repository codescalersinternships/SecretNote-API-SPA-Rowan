package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Id      int    `json:"id"`
}

func NewNote() Note {
	return Note{}
}

func getNote(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error method of '/getNote' request should be GET!")		
		return
	}
	_, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("error detected while reading request body")
		os.Exit(1)
	}
	// str := string(body)
	// fmt.Println(string(body))
	fmt.Fprint(w, "helloooooooooo")
}

func getNotes(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error method of '/getNotes' request should be GET!")		
		return
	}
	_, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("error detected while reading request body")
		os.Exit(1)
	}
	// str := string(body)
	// fmt.Println(string(body))
	fmt.Fprint(w, "helloooooooooo")
}

func createNote(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
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
}



func main() {
	http.HandleFunc("/getNote", getNote)
	http.HandleFunc("/getNotes", getNotes)
	http.HandleFunc("/createNote", createNote)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error detected while reading request body")
		os.Exit(1)
	}
}
