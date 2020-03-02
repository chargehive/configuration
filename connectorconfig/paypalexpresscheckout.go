package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type PayPalExpressCheckoutCredentials struct {
	APIUsername         *string           `json:"apiUsername" yaml:"apiUsername" validate:"required"`
	APIPassword         *string           `json:"apiPassword" yaml:"apiPassword" validate:"required"`
	APISignature        *string           `json:"apiSignature" yaml:"apiSignature" validate:"required"`
	SupportedCurrencies []string          `json:"supportedCurrencies" yaml:"supportedCurrencies" validate:"gt=0,dive,oneof=AUD BRL CAD CZK DKK EUR HKD HUF INR ILS JPY MYR MXN TWD NZD NOK PHP PLN GBP RUB SGD SEK CHF THB USD"`
	Environment         PayPalEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox live"`
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
