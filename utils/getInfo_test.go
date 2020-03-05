package utils

import (
	"github.com/chargehive/configuration"
	"github.com/go-playground/assert/v2"
	"testing"
)

// test for missing field
func TestGetInfo(t *testing.T) {
	rawJson := []byte(`{"kind":"Connector","metadata":{"projectId":"CHANGE-ME","name":"CHANGE-ME","displayName":"gary","description":"","annotations":null,"labels":null,"disabled":true},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"sandbox","configuration":"eyJtb2RlIjoiZHluYW1pYyIsInlyYW5zYWN0aW9uSURQcmVmaXgiOiIifQ=="}}`)
	configuration.Initialise()
	info, err := GetInfo(rawJson, "v1")
	assert.Equal(t, err, nil)
	assert.Equal(t, info, ConfigInfo{
		Kind:        "Connector",
		Version:     "v1",
		Name:        "CHANGE-ME",
		ProjectId:   "CHANGE-ME",
		Disabled:    true,
		DisplayName: "gary",
	})
}
