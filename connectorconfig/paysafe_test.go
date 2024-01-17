package connectorconfig

import (
	"testing"
)

func TestSupportsCountry(t *testing.T) {

	credential := &PaySafeCredentials{}

	tests := map[string]bool{
		"GB": true,
		"XK": false,
		"XX": false,
		"":   false,
	}

	for k, v := range tests {
		if credential.SupportsCountry(k) != v {
			t.Errorf("Expected %v, got %v", v, !v)
		}
	}
}
