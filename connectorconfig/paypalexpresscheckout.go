package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type PayPalExpressCheckoutCredentials struct {
	APIUsername         *string           `json:"apiUsername" yaml:"apiUsername"`
	APIPassword         *string           `json:"apiPassword" yaml:"apiPassword"`
	APISignature        *string           `json:"apiSignature" yaml:"apiSignature"`
	SupportedCurrencies []string          `json:"supportedCurrencies" yaml:"supportedCurrencies"`
	Environment         PayPalEnvironment `json:"environment" yaml:"environment"`
}

func (c PayPalExpressCheckoutCredentials) GetLibrary() Library {
	return LibraryPayPalExpressCheckout
}

func (c *PayPalExpressCheckoutCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *PayPalExpressCheckoutCredentials) Validate() error {
	return nil
}

func (c *PayPalExpressCheckoutCredentials) GetSecureFields() []*string {
	return []*string{c.APIUsername, c.APIPassword, c.APISignature}
}

func (c *PayPalExpressCheckoutCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *PayPalExpressCheckoutCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
