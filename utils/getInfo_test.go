package utils

import (
	"github.com/chargehive/configuration"
	"github.com/chargehive/configuration/object"
	"github.com/go-playground/assert/v2"
	"testing"
)

// test for missing field
func TestGetInfo(t *testing.T) {
	rawJson := []byte(`{"kind":"Project","metadata":{"projectId":"change-me","name":"change-me","displayName":"gary","description":"","annotations":null,"labels":null,"disabled":true},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"sandbox","configuration":"eyJtb2RlIjoiZHluYW1pYyIsInlyYW5zYWN0aW9uSURQcmVmaXgiOiIifQ=="}}`)
	configuration.Initialise()
	info, err := GetInfo(rawJson)
	assert.Equal(t, err, nil)
	assert.Equal(t, info, ConfigInfo{
		Kind: "Project",
		MetaData: object.MetaData{
			ProjectID:   "change-me",
			Name:        "change-me",
			DisplayName: "gary",
			Description: "",
			Annotations: nil,
			Labels:      nil,
			Disabled:    true,
		},
		SpecVersion: "v1",
	})
}
