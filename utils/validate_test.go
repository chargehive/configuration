package utils

import (
	"encoding/json"
	"fmt"
	"github.com/chargehive/configuration"
	"github.com/go-playground/assert/v2"
	"testing"
)

// test for accidental capitalisation ("type"->"Type")
func TestAdditionalUnknownVars(t *testing.T) {
	rawJson := []byte(`{"kind":"Connector","metadata":{"projectId":"test-project","name":"paypal-expresscheckout-connector"},"spec":{"Library":"paypal-expresscheckout","Type":"payment","Configuration":"ewogICJBUElVc2VybmFtZSI6ICJzYi1iaHF5dDE2OTYwX2FwaTEuYnVzaW5lc3MuZXhhbXBsZS5jb20iLAogICJBUElQYXNzd29yZCI6ICJOVVBGSjNBTjZYR1RCSzhTIiwKICAiQVBJU2lnbmF0dXJlIjogIkFmQll0Wkx0YnNOVUtLcFYuWHdaSjdOdUh6SGtBQmguYXBKRGk5MkRkcDRCWUNzT2NlRWVqckVqIiwKICAiU3VwcG9ydGVkQ3VycmVuY2llcyI6IFsKICAgICJBVUQiLAogICAgIkNBRCIsCiAgICAiRVVSIiwKICAgICJHQlAiLAogICAgIkpQWSIsCiAgICAiVVNEIiwKICAgICJCT0IiCiAgXSwKICAiRW52aXJvbm1lbnQiOiAic2FuZGJveCIKfQ=="}}`)
	configuration.Initialise()
	errs := Validate(rawJson, "v1")
	_ = PrettyPrint(errs)
	assert.Equal(t, errs["json"], "json: unknown field \"Type\"")
}

// test for missing field
func TestValidation(t *testing.T) {
	rawJson := []byte(`{"kind":"Connector","metadata":{"projectId":"test-project","name":"paypal-expresscheckout-connector"},"spec":{"Library":"paypal-expresscheckout","type":"payment","Configuration":"ewogICJBUElVc2VybmFtZSI6ICJzYi1iaHF5dDE2OTYwX2FwaTEuYnVzaW5lc3MuZXhhbXBsZS5jb20iLAogICJBUElQYXNzd29yZCI6ICJOVVBGSjNBTjZYR1RCSzhTIiwKICAiQVBJU2lnbmF0dXJlIjogIkFmQll0Wkx0YnNOVUtLcFYuWHdaSjdOdUh6SGtBQmguYXBKRGk5MkRkcDRCWUNzT2NlRWVqckVqIiwKICAiU3VwcG9ydGVkQ3VycmVuY2llcyI6IFsKICAgICJBVUQiLAogICAgIkNBRCIsCiAgICAiRVVSIiwKICAgICJHQlAiLAogICAgIkpQWSIsCiAgICAiVVNEIiwKICAgICJCT0IiCiAgXSwKICAiRW52aXJvbm1lbnQiOiAic2FuZGJveCIKfQ=="}}`)
	configuration.Initialise()
	if errs := Validate(rawJson, "v1"); len(errs) > 0 {
		_ = PrettyPrint(errs)
		assert.Equal(t, 1, len(errs))
	}
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
