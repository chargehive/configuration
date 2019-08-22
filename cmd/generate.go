package main

import (
	"encoding/json"
	"flag"
	"github.com/chargehive/configuration/object"
	"github.com/chargehive/configuration/v1/policy"
	"io/ioutil"
	"log"
	"os"
)

var outputFile = flag.String("out", "", "Output File Path")
var id = flag.String("id", "test-spec", "Spec ID")
var projectID = flag.String("project", "test-project", "Project ID")

func buildSpec() object.Specification {
	return policy.ScaPolicy{ShouldIdentify: true}
}

func main() {

	flag.Parse()

	spec := buildSpec()

	def := object.DefinitionFromSpec(spec)
	def.MetaData.Name = object.CleanName(*id)
	def.MetaData.ProjectID = *projectID

	data, err := json.Marshal(def)

	if err != nil {
		log.Print("Unable to convert spec to json ", err)
		return
	}

	if outputFile == nil || *outputFile == "" {
		log.Print(string(data))
	} else {
		err := ioutil.WriteFile(*outputFile, data, os.ModePerm)
		if err != nil {
			log.Print(err)
		} else {
			log.Print("Written data to ", *outputFile)
		}
	}
}
