package main

import (
	"encoding/json"
	"flag"
	"github.com/chargehive/configuration/object"
	"github.com/chargehive/configuration/selector"
	"github.com/chargehive/configuration/v1/scheduler"
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
	return scheduler.Initiator{Type: scheduler.InitiatorTypeAuth, InitialConnector: scheduler.ConnectorSelectorConfig,
		AttemptConfig: &scheduler.AttemptConfig{
			PoolType:                 scheduler.PoolTypeCascade,
			MethodSelector:           scheduler.MethodSelectorPrimaryMethod,
			OverridePoolConnectorIDs: []string{"paysafe-connector"}}}
}

func buildSelector() selector.Selector {
	return selector.Selector{Priority: 50, Expressions: []selector.Predicate{
		{Key: selector.KeyChargeAmountCurrency, Operator: selector.PredicateOperatorEqual, Conversion: selector.OperatorConversionDefault, Values: []string{"GBP"}},
	}}
}

func main() {

	flag.Parse()

	spec := buildSpec()

	def := object.DefinitionFromSpec(spec)
	def.MetaData.Name = object.CleanName(*id)
	def.MetaData.ProjectID = *projectID
	def.Selector = buildSelector()

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
