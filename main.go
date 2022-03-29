package main

import (
	"github.com/tonnytg/gptm/pkg/webpage"
	"net/http"
)

// main start api
func main() {
	http.HandleFunc("/", webpage.GetProjects)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
