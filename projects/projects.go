package projects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type MapProject struct {
	Projects []struct {
		ProjectNumber  string            `json:"projectNumber"`
		ProjectID      string            `json:"projectId"`
		LifecycleState string            `json:"lifecycleState"`
		Name           string            `json:"name"`
		Labels         map[string]string `json:"labels"`
		CreateTime     time.Time         `json:"createTime"`
	} `json:"projects"`
}

// GetProjects list each project and show CreateAT, Name and Labels
func GetProjects() {

	token := os.Getenv("GCP_API_KEY")
	//project := os.Getenv("PROJECT_ID")
	url := "https://cloudresourcemanager.googleapis.com/v1/projects"

	bearer := "Bearer " + token

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	} else {
		defer resp.Body.Close()

		data, _ := ioutil.ReadAll(resp.Body)
		var ps MapProject
		json.Unmarshal(data, &ps)

		//fmt.Println(string(data))

		for i := 0; i < len(ps.Projects); i++ {
			fmt.Println("Project:", ps.Projects[i].Name ,"-", ps.Projects[i].ProjectNumber)
			fmt.Println("CreateAt:", ps.Projects[i].CreateTime)
			fmt.Println("Tags:")
			for key, value := range ps.Projects[i].Labels {
				fmt.Printf("\t%s:%s\n", key, value)
			}
			fmt.Println("---")
		}
	}
}
