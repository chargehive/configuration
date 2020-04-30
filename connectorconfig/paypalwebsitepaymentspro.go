package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type PayPalWebsitePaymentsProCredentials struct {
	APIUsername            *string           `json:"apiUsername" yaml:"apiUsername" validate:"required,gt=0"`
	APIPassword            *string           `json:"apiPassword" yaml:"apiPassword" validate:"required,gt=0"`
	APISignature           *string           `json:"apiSignature" yaml:"apiSignature" validate:"required,gt=0"`
	SupportedCurrencies    []string          `json:"supportedCurrencies" yaml:"supportedCurrencies" validate:"gt=0,dive,oneof=AUD BRL CAD CZK DKK EUR HKD HUF INR ILS JPY MYR MXN TWD NZD NOK PHP PLN GBP RUB SGD SEK CHF THB USD"`
	CardinalProcessorID    *string           `json:"cardinalProcessorID" yaml:"cardinalProcessorID" validate:"required"`
	CardinalMerchantID     *string           `json:"cardinalMerchantID" yaml:"cardinalMerchantID" validate:"required"`
	CardinalTransactionPw  *string           `json:"cardinalTransactionPw" yaml:"cardinalTransactionPw" validate:"required"`
	CardinalTransactionURL *string           `json:"cardinalTransactionURL" yaml:"cardinalTransactionURL" validate:"required"`
	CardinalAPIIdentifier  *string           `json:"cardinalAPIIdentifier" yaml:"cardinalAPIIdentifier" validate:"required"`
	CardinalAPIKey         *string           `json:"cardinalAPIKey" yaml:"cardinalAPIKey" validate:"required"`
	CardinalOrgUnitID      *string           `json:"cardinalOrgUnitID" yaml:"cardinalOrgUnitID" validate:"required"`
	Environment            PayPalEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox live"`
}

func (c PayPalWebsitePaymentsProCredentials) GetCardinalProcessorID() string {
	if c.CardinalProcessorID == nil {
		return ""
	}
	return *c.CardinalProcessorID
}

func (c PayPalWebsitePaymentsProCredentials) GetCardinalMerchantID() string {
	if c.CardinalMerchantID == nil {
		return ""
	}
	return *c.CardinalMerchantID
}

func (c PayPalWebsitePaymentsProCredentials) GetCardinalTransactionPw() string {
	if c.CardinalTransactionPw == nil {
		return ""
	}
	return *c.CardinalTransactionPw
}

func (c PayPalWebsitePaymentsProCredentials) GetCardinalTransactionURL() string {
	if c.CardinalTransactionURL == nil {
		return ""
	}
	return *c.CardinalTransactionURL
}

func (c PayPalWebsitePaymentsProCredentials) GetCardinalAPIIdentifier() string {
	if c.CardinalAPIIdentifier == nil {
		return ""
	}
	return *c.CardinalAPIIdentifier
}

func (c PayPalWebsitePaymentsProCredentials) GetCardinalAPIKey() string {
	if c.CardinalAPIKey == nil {
		return ""
	}
	return *c.CardinalAPIKey
}

func (c PayPalWebsitePaymentsProCredentials) GetCardinalOrgUnitID() string {
	if c.CardinalOrgUnitID == nil {
		return ""
	}
	return *c.CardinalOrgUnitID
}

func (c PayPalWebsitePaymentsProCredentials) GetLibrary() Library {
	return LibraryPayPalWebsitePaymentsPro
}

func (c *PayPalWebsitePaymentsProCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *PayPalWebsitePaymentsProCredentials) Validate() error {
	return nil
}

func (c *PayPalWebsitePaymentsProCredentials) GetSecureFields() []*string {
	return []*string{c.APIUsername, c.APIPassword, c.APISignature, c.CardinalTransactionPw, c.CardinalAPIIdentifier, c.CardinalAPIKey}
}

func (c *PayPalWebsitePaymentsProCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *PayPalWebsitePaymentsProCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *PayPalWebsitePaymentsProCredentials) SupportsSca() bool {
	return c.GetCardinalProcessorID() != "" &&
		c.GetCardinalMerchantID() != "" &&
		c.GetCardinalTransactionPw() != "" &&
		c.GetCardinalTransactionURL() != "" &&
		c.GetCardinalAPIIdentifier() != "" &&
		c.GetCardinalAPIKey() != "" &&
		c.GetCardinalOrgUnitID() != ""
}

func (c *PayPalWebsitePaymentsProCredentials) SupportsApplePay() bool {
	return false
}
