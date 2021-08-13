package webserver

import (
	"fmt"
	"github.com/tonnytg/gptm/projects"
	"log"
	"net/http"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	data, err := projects.GetProjects()
	if err != nil {
		log.Println("error: getting projects:", err)
	}

	fmt.Fprintf(w, "%s", data)
}

