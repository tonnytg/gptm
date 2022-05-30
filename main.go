package main

import (
	"net/http"

	"github.com/tonnytg/gptm/internal/webpage"
)

// main start api
func main() {
	http.HandleFunc("/", webpage.GetProjects)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
