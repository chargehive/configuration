package connectorconfig

import (
	"encoding/json"
	"github.com/chargehive/proto/golang/chargehive/chtype"

	"github.com/chargehive/configuration/v1/connector"
)

type PayPalExpressCheckoutCredentials struct {
	APIUsername         *string           `json:"apiUsername" yaml:"apiUsername" validate:"required,gt=0"`
	APIPassword         *string           `json:"apiPassword" yaml:"apiPassword" validate:"required,gt=0"`
	APISignature        *string           `json:"apiSignature" yaml:"apiSignature" validate:"required,gt=0"`
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

func (c PayPalExpressCheckoutCredentials) SupportsSca() bool {
	return false
}

func (c PayPalExpressCheckoutCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_PAYPAL {
		return true
	}
	return false
}

func (c PayPalExpressCheckoutCredentials) CanSandboxPlanUse() bool {
	if c.Environment == PayPalEnvironmentLive {
		return false
	}
	return true
}
