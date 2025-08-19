package main

import (
	"fmt"
	"net/http"
)

type Request struct {
	Text string `json:"text"`
}

type Response struct {
	Lines int `json:"lines"`
	Words int `json:"words"`
	Bytes int `json:"bytes"`
}

func main() {
	http.HandleFunc("/count", simpleHandler)

	fmt.Println("seRveR running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
