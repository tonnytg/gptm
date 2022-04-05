package main

import (
	"github.com/tonnytg/gptm/pkg/webpage"
	"net/http"
)

func main() {
	http.HandleFunc("/", webpage.GetProjects)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
