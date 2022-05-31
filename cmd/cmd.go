package cmd

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tonnytg/gptm/internal/webpage"
)

func Flags() {

	webserver := flag.String("webserver", "", "--webserver <STATUS>")
	project := flag.String("project", "", "--project PROJECT_NAME")

	flag.Parse()
	if *project != "" {
		fmt.Printf("Project name: %s\n", *project)
		os.Exit(0)
	}

	if *webserver == "enabled" {
		http.HandleFunc("/", webpage.GetProjects)
		fmt.Printf("webserver status:%s\n", *webserver)
		log.Printf("webserver started: %s\n", time.Now())
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
		log.Printf("webserver stoped\n")
	}
}
