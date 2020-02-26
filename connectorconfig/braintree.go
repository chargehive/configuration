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
	PublicKey         *string              `json:"publicKey" yaml:"publicKey" validate:"required"`
	PrivateKey        *string              `json:"privateKey" yaml:"privateKey" validate:"required"`
	MerchantID        string               `json:"merchantID" yaml:"merchantID" validate:"required"`
	MerchantAccountID string               `json:"merchantAccountID" yaml:"merchantAccountID" validate:"required"`
	Currency          string               `json:"currency" yaml:"currency" validate:"required"`
	Environment       BraintreeEnvironment `json:"environment" yaml:"environment" validate:"required"`
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
