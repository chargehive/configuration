package connectorconfig

import (
	"encoding/json"
	"github.com/chargehive/configuration/v1/connector"
)

type AuthorizeEnvironment string

const (
	AuthorizeEnvironmentSandbox    AuthorizeEnvironment = "sandbox"
	AuthorizeEnvironmentProduction AuthorizeEnvironment = "production"
)

type AuthorizeCredentials struct {
	APILoginID     *string
	TransactionKey *string
	Environment    AuthorizeEnvironment
}

func (c AuthorizeCredentials) GetLibrary() Library {
	return LibraryAuthorize
}
func (c AuthorizeCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *AuthorizeCredentials) Validate() error {
	return nil
}

func (c *AuthorizeCredentials) GetSecureFields() []*string {
	return []*string{c.APILoginID, c.TransactionKey}
}

func (c *AuthorizeCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *AuthorizeCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}