package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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

		fmt.Println("Data:\n", string(data))
	}
}

func TestGetProjects() {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, cloudresourcemanager.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	cloudresourcemanagerService, err := cloudresourcemanager.New(c)
	if err != nil {
		log.Fatal(err)
	}

	req := cloudresourcemanagerService.Projects.List()
	if err := req.Pages(ctx, func(page *cloudresourcemanager.ListProjectsResponse) error {
		for _, project := range page.Projects {
			// TODO: Change code below to process each `project` resource:
			fmt.Printf("%#v\n", project)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
