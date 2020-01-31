package tests

import (
	"encoding/json"
	"fmt"
	"github.com/chargehive/configuration"
	"github.com/chargehive/configuration/object"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestValidation(t *testing.T) {
	configuration.Initialise()
	obj := &object.Definition{}
	rawJson := []byte(`{
  "kind": "SchedulerSequential-bad",
  "metadata": {
    "projectId": "test-project",
    "name": "refund-scheduler"
  },
  "selector": {
    "priority": 10,
    "expressions": [
      {
        "key": "transaction.type-bad",
        "operator": "Equal-bad",
        "values": [
          "TRANSACTION_TYPE_REFUND"
        ]
      }
    ]
  },
  "spec": {
    "Schedules": {
      "0": {
        "TimeDelayOrigin": "now",
        "TimeDelaySync": "",
        "TimeDelay": 60000000000,
        "TimeSyncHour": 0,
        "TimeSyncZone": ""
      },
      "1": {
        "TimeDelayOrigin": "now",
        "TimeDelaySync": "",
        "TimeDelay": 60000000000,
        "TimeSyncHour": 0,
        "TimeSyncZone": ""
      },
      "2": {
        "TimeDelayOrigin": "now",
        "TimeDelaySync": "",
        "TimeDelay": 60000000000,
        "TimeSyncHour": 0,
        "TimeSyncZone": ""
      }
    }
  }
}`)

	jsonErr := json.Unmarshal(rawJson, obj)
	assert.Equal(t, nil, jsonErr)
	err := obj.Validate()
	vErrs, ok := err.(validator.ValidationErrors)
	assert.Equal(t, ok, true)
	assert.Equal(t, len(vErrs), 3)
	fmt.Println(vErrs)
}
