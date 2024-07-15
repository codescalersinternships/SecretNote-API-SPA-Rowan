package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
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

func main() {
	http.HandleFunc("/getNote", getNote)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error detected while reading request body")
		os.Exit(1)
	}
}
