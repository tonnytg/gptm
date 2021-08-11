package main

import (
	"github.com/tonnytg/gptm/projects"
	"log"
)

func main() {
	err := projects.GetProjects()
	if err != nil {
		log.Println("error: getting projects:", err)
	}
}
