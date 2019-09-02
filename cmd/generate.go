package main

import (
	"encoding/json"
	"flag"
	"github.com/LucidCube/chargehive-transport-config/connectorconfig"
	"github.com/chargehive/configuration/object"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var outputFile = flag.String("out", "", "Output File Path")
var outputDir = flag.String("outd", "", "Output File Directory")
var id = flag.String("id", "test-spec", "Spec ID")
var projectID = flag.String("project", "", "Project ID")

func buildSpec() object.Specification {
	creds := connectorconfig.SandboxCredentials{
		Mode:                connectorconfig.SandboxModeDynamic,
		TransactionIDPrefix: "staging-",
	}
	return creds.ToConnector()
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

	outputLocation := ""
	if outputDir != nil {
		outputLocation = *outputDir
	}

	if outputFile != nil && *outputFile != "" {
		outputLocation = path.Join(outputLocation, *outputFile)
	} else if outputLocation != "" {
		outputLocation = path.Join(outputLocation, def.GetID()+".json")
	}

	if outputLocation == "" {
		log.Print(string(data))
	} else {
		err := ioutil.WriteFile(outputLocation, data, os.ModePerm)
		if err != nil {
			log.Print(err)
		} else {
			log.Print("Written data to ", outputLocation)
		}
	}
}
