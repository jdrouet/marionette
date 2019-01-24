package parser

import (
	"encoding/json"
	"io/ioutil"
	"os"

	model "github.com/jdrouet/marionette/cmd/model"
)

// Parse the configuration file
func Parse(configurationPath string) (model.Repository, error) {
	var repo model.Repository
	jsonFile, err := os.Open(configurationPath)
	if err != nil {
		return repo, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return repo, err
	}
	err = json.Unmarshal(byteValue, &repo)
	return repo, err
}
