package connectorconfig

import (
	"encoding/json"
	"strings"

	"github.com/chargehive/configuration/v1/connector"
)

// CreateCredentials create credentials from json
func GetCredentials(c *connector.Connector) (Credentials, error) {
	return getCreds(c, false)
}

func GetCredentialsStrict(c *connector.Connector) (Credentials, error) {
	return getCreds(c, true)
}

func getCreds(c *connector.Connector, strict bool) (Credentials, error) {
	credentials, err := Library(c.Library).GetCredential()

	if err != nil {
		return credentials, err
	}

	if c.Configuration != nil && string(c.Configuration) != "" {
		reader := strings.NewReader(string(c.Configuration))
		dec := json.NewDecoder(reader)
		if strict {
			dec.DisallowUnknownFields()
		}
		err = dec.Decode(credentials)
	}

	return credentials, err
}

type MerchantIdentifier interface {
	GetMID() string
}
