package webpage

import (
	"encoding/json"
	"github.com/tonnytg/gptm/entity/projects"
	"github.com/tonnytg/gptm/pkg/webTool"
	"html/template"
	"log"
	"net/http"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {

	var p projects.MapProject
	url := "https://cloudresourcemanager.googleapis.com/v1/projects"
	data, err := webTool.Get(url)

	json.Unmarshal(data, &p)
	if err != nil {
		log.Println("error: getting projects:", err)
	}

	t, err := template.ParseFiles("./pkg/webpage/templates/projects.html")
	t.Execute(w, p.Projects)
}
