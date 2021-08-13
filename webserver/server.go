package webserver

import (
	"encoding/json"
	"fmt"
	"github.com/tonnytg/gptm/projects"
	"log"
	"net/http"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {

	var p projects.MapProject
	data, err := projects.GetProjects()
	json.Unmarshal(data, &p)
	if err != nil {
		log.Println("error: getting projects:", err)
	}

	for i := 0; i < len(p.Projects); i++ {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Projects[i].Name, p.Projects[i].ProjectID)
	}
}
