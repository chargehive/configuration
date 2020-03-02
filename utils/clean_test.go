package utils

import (
	"fmt"
	"github.com/chargehive/configuration"
	"github.com/go-playground/assert/v2"
	"testing"
)

// TestClean ensures that an additional fields in main struct and serialized config are removed
func TestClean(t *testing.T) {
	configuration.Initialise()
	rawJson := []byte(`{"Kind":"Connector","metadata":{"projectId":"CHANGE-ME","name":"CHANGE-ME","uuid":"","displayName":"","description":"","annotations":null,"labels":null},"specVersion":"v1","cats":"dogs","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":"eyJhcGlMb2dpbklkIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InNhbmRib3giLCAiZ2FyeSI6IndhcyBoZXJlIn0="}}`)
	fmt.Printf("Before %v\n", string(rawJson))
	isCleaned, output, err := Clean(rawJson, "v1", false)
	fmt.Printf("After  %v\nCleaned: %t, Error: %v\n", string(output), isCleaned, err)
	assert.Equal(t, isCleaned, true)
	assert.Equal(t, err, nil)
	assert.Equal(t, output, []byte(`{"Kind":"Connector","metadata":{"projectId":"CHANGE-ME","name":"CHANGE-ME","uuid":"","displayName":"","description":"","annotations":null,"labels":null},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":"eyJhcGlMb2dpbklkIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InNhbmRib3gifQ=="}}`))
}

// test that a good input returns isCleaned = false
func TestCleanNoChanges(t *testing.T) {
	configuration.Initialise()
	rawJson := []byte(`{"Kind":"Connector","metadata":{"projectId":"CHANGE-ME","name":"CHANGE-ME","uuid":"","displayName":"","description":"","annotations":null,"labels":null},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":"eyJhcGlMb2dpbklkIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InNhbmRib3gifQ=="}}`)
	fmt.Printf("Before %v\n", string(rawJson))
	isCleaned, output, err := Clean(rawJson, "v1", false)
	fmt.Printf("After  %v\nCleaned: %t, Error: %v\n", string(output), isCleaned, err)
	assert.Equal(t, isCleaned, false)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(output), string([]byte(`{"Kind":"Connector","metadata":{"projectId":"CHANGE-ME","name":"CHANGE-ME","uuid":"","displayName":"","description":"","annotations":null,"labels":null},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":"eyJhcGlMb2dpbklkIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InNhbmRib3gifQ=="}}`)))
}

// test that a good input in a different order isCleaned = false
func TestCleanReOrder(t *testing.T) {
	configuration.Initialise()
	rawJson := []byte(`{"Kind":"Connector","metadata":{"name":"CHANGE-ME","projectId":"CHANGE-ME","uuid":"","displayName":"","description":"","annotations":null,"labels":null},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":"eyJhcGlMb2dpbklkIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InNhbmRib3gifQ=="}}`)
	fmt.Printf("Before %v\n", string(rawJson))
	isCleaned, output, err := Clean(rawJson, "v1", false)
	fmt.Printf("After  %v\nCleaned: %t, Error: %v\n", string(output), isCleaned, err)
	assert.Equal(t, isCleaned, false)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(output), string([]byte(`{"Kind":"Connector","metadata":{"projectId":"CHANGE-ME","name":"CHANGE-ME","uuid":"","displayName":"","description":"","annotations":null,"labels":null},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":"eyJhcGlMb2dpbklkIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InNhbmRib3gifQ=="}}`)))
}

// test for invalid json input
func TestCleanError(t *testing.T) {
	configuration.Initialise()
	rawJson := []byte(`{"Kind":::"Connector","metadasdfta":{"name":"CHANGE-ME","uuid":"","displayName":"","description":"","annotaions":null,"labels":null},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":"eyJhcGlMb2dpbklkIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InNhbmRib3gifQ=="}}`)
	fmt.Printf("Before %v\n", string(rawJson))
	isCleaned, output, err := Clean(rawJson, "v1", false)
	fmt.Printf("After  %v\nCleaned: %t, Error: %v\n", string(output), isCleaned, err)
	assert.Equal(t, isCleaned, false)
	assert.NotEqual(t, err, nil)
}
