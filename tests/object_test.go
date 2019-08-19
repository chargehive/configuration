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

	o := object.InstanceFromSpec(scheduler.Initiator{Type: scheduler.InitiatorTypeAuth, InitialConnector: scheduler.ConnectorSelectorStickyAny})

	jsn, err := json.Marshal(o)
	if err != nil {
		t.Error(err)
	}

	log.Print(string(jsn))

	obj, err := object.FromJson([]byte(`{"kind":"Initiator","metadata":{"projectId":"","name":""},"specVersion":"v1","selector":{},"spec":{"Type":"auth","InitialConnector":"sticky-any","AttemptConfig":null}}`))
	ii, err := scheduler.NewInitiatorInstance(obj)
	if err != nil {
		t.Error(err)
	}

	jsn, err = json.Marshal(ii)
	log.Print(string(jsn))
	jsn, err = json.Marshal(obj)
	log.Print(string(jsn))
	if err != nil {
		t.Error(err)
	}

	log.Print(string(jsn))
}
