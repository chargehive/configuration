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
	APILoginID     *string              `json:"apiLoginId" yaml:"apiLoginId" validate:"required,gt=0"`
	TransactionKey *string              `json:"transactionKey" yaml:"transactionKey" validate:"required,gt=0"`
	Environment    AuthorizeEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
}

func (c AuthorizeCredentials) GetAPILoginID() string {
	if c.APILoginID == nil {
		return ""
	}
	return *c.APILoginID
}

func (c AuthorizeCredentials) GetTransactionKey() string {
	if c.TransactionKey == nil {
		return ""
	}
	return *c.TransactionKey
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

func (c AuthorizeCredentials) SupportsSca() bool {
	return c.GetAPILoginID() != "" && c.GetTransactionKey() != ""
}
