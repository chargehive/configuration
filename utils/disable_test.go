package utils

import (
	"github.com/chargehive/configuration"
	"github.com/go-playground/assert/v2"
	"testing"
)

// Checks that a disabled config is enabled
func TestEnableConfig(t *testing.T) {
	rawJson := []byte(`{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":true},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"cybersource","configuration":"eyJtZXJjaGFudElEIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InRlc3QifQ=="}}`)
	configuration.Initialise()
	newJson, err := Disable(false, rawJson, "v1")
	assert.Equal(t, err, nil)
	assert.Equal(t, string(newJson), string([]byte(`{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"cybersource","configuration":"eyJtZXJjaGFudElEIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InRlc3QifQ=="}}`)))
}

// Checks that an enabled config is disabled
func TestDisableConfig(t *testing.T) {
	rawJson := []byte(`{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"cybersource","configuration":"eyJtZXJjaGFudElEIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InRlc3QifQ=="}}`)
	configuration.Initialise()
	newJson, err := Disable(true, rawJson, "v1")
	assert.Equal(t, err, nil)
	assert.Equal(t, string(newJson), string([]byte(`{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":true},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"cybersource","configuration":"eyJtZXJjaGFudElEIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InRlc3QifQ=="}}`)))
}

// Checks that a disabled config remains disabled
func TestDisableNoChangeConfig(t *testing.T) {
	rawJson := []byte(`{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":true},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"cybersource","configuration":"eyJtZXJjaGFudElEIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InRlc3QifQ=="}}`)
	configuration.Initialise()
	newJson, err := Disable(true, rawJson, "v1")
	assert.Equal(t, err, nil)
	assert.Equal(t, string(newJson), string([]byte(`{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":true},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"cybersource","configuration":"eyJtZXJjaGFudElEIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InRlc3QifQ=="}}`)))
}

// checks for an error when there is an issue with config
func TestDisableConfigError(t *testing.T) {
	rawJson := []byte(`{"kind":"GARY","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"cybersource","configuration":"eyJtZXJjaGFudElEIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InRlc3QifQ=="}}`)
	configuration.Initialise()
	newJson, err := Disable(true, rawJson, "v1")
	assert.Equal(t, err.Error(), "kind:`GARY`, version:`v1` has not been implemented")
	assert.Equal(t, string(newJson), "")
}
