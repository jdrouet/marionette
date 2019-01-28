package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	model "github.com/jdrouet/marionette/cmd/model"
	parser "github.com/jdrouet/marionette/cmd/parser"
)

func check(list []model.Project, project string) {
	for _, item := range list {
		if item.Name == project {
			fmt.Println("true")
			return
		}
	}
	fmt.Println("false")
}

func display(projects []model.Project) {
	jsonOutput, err := json.Marshal(projects)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem converting to json\n")
		os.Exit(4)
	}
	fmt.Println(string(jsonOutput))
}

func main() {
	currentDirector, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Impossible to get the current directory\n")
		os.Exit(1)
	}

	contextPtr := flag.String("context", currentDirector, "Directory of the repository")
	configFilePtr := flag.String("config", "marionette.json", "The configuration file containing the description of the project")
	referencePtr := flag.String("reference", "master", "The branch that is the reference")
	checkPtr := flag.String("check", "", "Check if the project has changed")
	flag.Parse()

	repo, err := parser.Parse(*configFilePtr)

	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "The configuration file %q cannot be found\n", *configFilePtr)
		os.Exit(2)
	}

	err = repo.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while init...\n")
		os.Exit(2)
	}

	projects, err := repo.Digest(*contextPtr, *referencePtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while running...\n")
		os.Exit(2)
	}

	if len(*checkPtr) > 0 {
		check(projects, *checkPtr)
	} else {
		display(projects)
	}
	os.Exit(0)
}
