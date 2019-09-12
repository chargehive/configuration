package connectorconfig

import (
	"encoding/json"
	"github.com/chargehive/configuration/v1/connector"
)

type StripeCredentials struct {
	APIKey *string
}

func (c StripeCredentials) GetLibrary() Library {
	return LibraryStripe
}

func (c *StripeCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *StripeCredentials) Validate() error {
	return nil
}

func (c *StripeCredentials) GetSecureFields() []*string {
	return []*string{c.APIKey}
}

func (c *StripeCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *StripeCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
