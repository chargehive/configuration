package tests

import (
	"fmt"
	"github.com/chargehive/configuration"
	"github.com/chargehive/configuration/utils"
	"testing"
)

func TestValidation(t *testing.T) {
	configuration.Initialise()
	rawJson := []byte(`{"kind":"Connector","metadata":{"projectId":"test-project","name":"paypal-expresscheckout-connector"},"spec":{"Library":"paypal-expresscheckout","Type":"payment","Configuration":"ewogICJBUElVc2VybmFtZSI6ICJzYi1iaHF5dDE2OTYwX2FwaTEuYnVzaW5lc3MuZXhhbXBsZS5jb20iLAogICJBUElQYXNzd29yZCI6ICJOVVBGSjNBTjZYR1RCSzhTIiwKICAiQVBJU2lnbmF0dXJlIjogIkFmQll0Wkx0YnNOVUtLcFYuWHdaSjdOdUh6SGtBQmguYXBKRGk5MkRkcDRCWUNzT2NlRWVqckVqIiwKICAiU3VwcG9ydGVkQ3VycmVuY2llcyI6IFsKICAgICJBVUQiLAogICAgIkNBRCIsCiAgICAiRVVSIiwKICAgICJHQlAiLAogICAgIkpQWSIsCiAgICAiVVNEIiwKICAgICJCT0IiCiAgXSwKICAiRW52aXJvbm1lbnQiOiAic2FuZGJveCIKfQ=="}}`)

	if errs := utils.Validate(rawJson); len(errs) > 0 {
		fmt.Println(errs)
	}
}
