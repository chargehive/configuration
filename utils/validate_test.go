package utils

import (
	"fmt"
	"github.com/chargehive/configuration"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestValidation(t *testing.T) {
	rawJson := []byte(`{"kind":"Connector","metadata":{"projectId":"test-project","name":"paypal-expresscheckout-connector"},"spec":{"Library":"paypal-expresscheckout","Type":"payment","Configuration":"ewogICJBUElVc2VybmFtZSI6ICJzYi1iaHF5dDE2OTYwX2FwaTEuYnVzaW5lc3MuZXhhbXBsZS5jb20iLAogICJBUElQYXNzd29yZCI6ICJOVVBGSjNBTjZYR1RCSzhTIiwKICAiQVBJU2lnbmF0dXJlIjogIkFmQll0Wkx0YnNOVUtLcFYuWHdaSjdOdUh6SGtBQmguYXBKRGk5MkRkcDRCWUNzT2NlRWVqckVqIiwKICAiU3VwcG9ydGVkQ3VycmVuY2llcyI6IFsKICAgICJBVUQiLAogICAgIkNBRCIsCiAgICAiRVVSIiwKICAgICJHQlAiLAogICAgIkpQWSIsCiAgICAiVVNEIiwKICAgICJCT0IiCiAgXSwKICAiRW52aXJvbm1lbnQiOiAic2FuZGJveCIKfQ=="}}`)

	configuration.Initialise()
	if errs := Validate(rawJson, "v1"); len(errs) > 0 {
		fmt.Println(errs)
	}
}

// Added "cats" to config to check if unknown fields are correctly identified
func TestAdditionalUnknownVars(t *testing.T) {
	rawJson := []byte(`{"Kind":"Connector","metadata":{"projectId":"CHANGE-ME","name":"CHANGE-ME","uuid":"","displayName":"","description":"","annotaions":null,"labels":null},"specVersion":"v1","cats":"dogs","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":"eyJhcGlMb2dpbklkIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InNhbmRib3gifQ=="}}`)

	configuration.Initialise()
	if errs := Validate(rawJson, "v1"); len(errs) > 0 {
		assert.Equal(t, 1, len(errs))
	}
}
