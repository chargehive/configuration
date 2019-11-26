package tests

import (
	"encoding/json"
	"github.com/chargehive/configuration"
	"github.com/chargehive/configuration/object"
	"github.com/chargehive/configuration/v1/scheduler"
	"log"
	"testing"
)

func TestUnmarshall(t *testing.T) {
	configuration.Initialise()

	o := object.DefinitionFromSpec(scheduler.Initiator{Type: scheduler.InitiatorTypeAuth, InitialConnector: scheduler.ConnectorSelectorStickyAny})

	jsn, err := json.Marshal(o)
	if err != nil {
		t.Error(err)
	}

	log.Print(string(jsn))

	obj, err := object.FromJson([]byte(`{"kind":"Initiator","metadata":{"projectId":"","name":""},"specVersion":"v1","selector":{},"spec":{"Type":"auth","InitialConnector":"sticky-any","AttemptConfig":null}}`))
	if err != nil {
		t.Error(err)
	}

	ii, err := scheduler.NewInitiatorDefinition(obj)
	if err != nil {
		t.Error(err)
	}

	jsn, err = json.Marshal(ii)
	if err != nil {
		t.Error(err)
	}
	log.Print(string(jsn))
	jsn, err = json.Marshal(obj)
	log.Print(string(jsn))
	if err != nil {
		t.Error(err)
	}

	log.Print(string(jsn))
}
