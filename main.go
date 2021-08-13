package main

import (
	"github.com/tonnytg/gptm/webserver"
	"net/http"
)

func main() {
	http.HandleFunc("/", webserver.GetProjects )
	http.ListenAndServe(":8080", nil)
}
