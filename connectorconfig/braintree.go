package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type BraintreeEnvironment string

const (
	BraintreeEnvironmentSandbox    BraintreeEnvironment = "sandbox"
	BraintreeEnvironmentProduction BraintreeEnvironment = "production"
)

type BraintreeCredentials struct {
	PublicKey         *string
	PrivateKey        *string
	MerchantID        string
	MerchantAccountID string
	Currency          string
	Environment       BraintreeEnvironment
}

func (c BraintreeCredentials) GetLibrary() Library {
	return LibraryBraintree
}

func (c BraintreeCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *BraintreeCredentials) Validate() error {
	return nil
}

func (c *BraintreeCredentials) GetSecureFields() []*string {
	return []*string{c.PublicKey, c.PrivateKey}
}

func (c *BraintreeCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *BraintreeCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
