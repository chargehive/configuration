package tests

import (
	"fmt"
	"github.com/chargehive/configuration"
	"github.com/chargehive/configuration/object"
	"testing"
)

func TestValidation(t *testing.T) {
	configuration.Initialise()
	rawJson := []byte(`{"Kind":"ConnectorPool","metadata":{"projectId":"pcprotect","name":"pool-eur","uuid":"84ae2bc3-367e-4dc0-b007-e78a36ca53d5"},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","values":["EUR"]}]},"spec":{"restriction":"noRepeat","connectors":[{"connectorId":"chain-eur-1002417380","weighting":35},{"connectorId":"paysafe-eur-1002448944","weighting":65}]}}`)

	def, err := object.FromJson(rawJson)
	if err != nil {
		fmt.Println(err)
		return
	}

	if errs := def.Validate(); len(errs) > 0 {
		fmt.Println(errs)
	}
}
