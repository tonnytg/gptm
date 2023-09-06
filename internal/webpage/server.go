package webpage

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/tonnytg/gptm/entity/projects"
	"github.com/tonnytg/gptm/internal/webTool"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {

	var p projects.MapProject
	url := "https://cloudresourcemanager.googleapis.com/v1/projects"
	data, err := webTool.Get(url)

	json.Unmarshal(data, &p)
	if err != nil {
		log.Println("error: getting projects:", err)
	}

	t, err := template.ParseFiles("./internal/webpage/templates/projects.html")
	t.Execute(w, p.Projects)
}
