package utils

import (
	"fmt"
	"testing"

	"github.com/chargehive/configuration"
	"github.com/go-playground/assert/v2"
)

// TestClean ensures that an additional fields in main struct and serialized config are removed
func TestClean(t *testing.T) {
	configuration.Initialise()
	rawJson := []byte(`{"Kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","cats":"dogs","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":{"apiLoginId":"CHANGE-ME","transactionKey":"CHANGE-ME","environment":"sandbox", "gary":"was here"}}}`)
	fmt.Printf("Before %v\n", string(rawJson))
	modified, output, err := Clean(rawJson, "v1", false)
	fmt.Printf("After  %v\nModified: %t, Error: %v\n", string(output), modified, err)
	assert.Equal(t, modified, true)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(output), `{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":{"apiLoginId":"CHANGE-ME","transactionKey":"CHANGE-ME","environment":"sandbox"}}}`)
}

// test that a good input returns modified = false
func TestCleanNoChanges(t *testing.T) {
	configuration.Initialise()
	rawJson := []byte(`{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":{"apiLoginId":"CHANGE-ME","transactionKey":"CHANGE-ME","environment":"sandbox"}}}`)
	fmt.Printf("Before %v\n", string(rawJson))
	modified, output, err := Clean(rawJson, "v1", false)
	fmt.Printf("After  %v\nModified: %t, Error: %v\n", string(output), modified, err)
	assert.Equal(t, modified, false)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(output), `{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":{"apiLoginId":"CHANGE-ME","transactionKey":"CHANGE-ME","environment":"sandbox"}}}`)
}

// test that a good input in a different order modified = false
func TestCleanReOrderB64(t *testing.T) {
	configuration.Initialise()
	rawJson := []byte(`{"kind":"Connector","metadata":{"name":"change-me","projectId":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":"eyJhcGlMb2dpbklkIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InNhbmRib3gifQ=="}}`)
	fmt.Printf("Before %v\n", string(rawJson))
	modified, output, err := Clean(rawJson, "v1", false)
	fmt.Printf("After  %v\nModified: %t, Error: %v\n", string(output), modified, err)
	assert.Equal(t, modified, true)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(output), `{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":{"apiLoginId":"CHANGE-ME","transactionKey":"CHANGE-ME","environment":"sandbox"}}}`)
}

// test that a good input in a different order modified = false
func TestCleanReOrderJson(t *testing.T) {
	configuration.Initialise()
	rawJson := []byte(`{"kind":"Connector","metadata":{"name":"change-me","projectId":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":{"apiLoginId":"CHANGE-ME","transactionKey":"CHANGE-ME","environment":"sandbox"}}}`)
	fmt.Printf("Before %v\n", string(rawJson))
	modified, output, err := Clean(rawJson, "v1", false)
	fmt.Printf("After  %v\nModified: %t, Error: %v\n", string(output), modified, err)
	assert.Equal(t, modified, false)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(output), `{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":{"apiLoginId":"CHANGE-ME","transactionKey":"CHANGE-ME","environment":"sandbox"}}}`)
}

// test for invalid json input
func TestCleanErrorB64(t *testing.T) {
	configuration.Initialise()
	rawJson := []byte(`{"Kind":::"Connector","metadasdfta":{"name":"change-me","displayName":"","description":"","annotaions":null,"labels":null},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":"eyJhcGlMb2dpbklkIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InNhbmRib3gifQ=="}}`)
	fmt.Printf("Before %v\n", string(rawJson))
	modified, output, err := Clean(rawJson, "v1", false)
	fmt.Printf("After  %v\nModified: %t, Error: %v\n", string(output), modified, err)
	assert.Equal(t, modified, false)
	assert.NotEqual(t, err, nil)
}

// test for invalid json input
func TestCleanErrorJson(t *testing.T) {
	configuration.Initialise()
	rawJson := []byte(`{"Kind":::"Connector","metadasdfta":{"name":"change-me","displayName":"","description":"","annotaions":null,"labels":null},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":{"apiLoginId":"CHANGE-ME","transactionKey":"CHANGE-ME","environment":"sandbox"}}}`)
	fmt.Printf("Before %v\n", string(rawJson))
	modified, output, err := Clean(rawJson, "v1", false)
	fmt.Printf("After  %v\nModified: %t, Error: %v\n", string(output), modified, err)
	assert.Equal(t, modified, false)
	assert.NotEqual(t, err, nil)
}
