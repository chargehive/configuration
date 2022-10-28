package connectorconfig

import (
	"encoding/base64"
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

	str, isStr := c.Configuration.(string)
	if !isStr {
		s, _ := json.Marshal(c.Configuration)
		str = string(s)
	} else {
		// check base64
		s, err := base64.StdEncoding.DecodeString(str)
		if err == nil {
			str = string(s)
		}
	}
	reader := strings.NewReader(str)
	dec := json.NewDecoder(reader)
	if strict {
		dec.DisallowUnknownFields()
	}

	return credentials, dec.Decode(credentials)
}
