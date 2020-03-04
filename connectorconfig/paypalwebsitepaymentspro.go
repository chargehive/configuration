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
	CardinalProcessorID    *string           `json:"cardinalProcessorID" yaml:"cardinalProcessorID" validate:"required_with=cardinalMerchantID cardinalTransactionPw cardinalTransactionURL cardinalAPIIdentifier cardinalAPIKey cardinalOrgUnitID"`
	CardinalMerchantID     *string           `json:"cardinalMerchantID" yaml:"cardinalMerchantID" validate:"required_with=cardinalProcessorID cardinalTransactionPw cardinalTransactionURL cardinalAPIIdentifier cardinalAPIKey cardinalOrgUnitID"`
	CardinalTransactionPw  *string           `json:"cardinalTransactionPw" yaml:"cardinalTransactionPw" validate:"required_with=cardinalProcessorID cardinalMerchantID cardinalTransactionURL cardinalAPIIdentifier cardinalAPIKey cardinalOrgUnitID"`
	CardinalTransactionURL *string           `json:"cardinalTransactionURL" yaml:"cardinalTransactionURL" validate:"required_with=cardinalProcessorID cardinalMerchantID cardinalTransactionPw cardinalAPIIdentifier cardinalAPIKey cardinalOrgUnitID"`
	CardinalAPIIdentifier  *string           `json:"cardinalAPIIdentifier" yaml:"cardinalAPIIdentifier" validate:"required_with=cardinalProcessorID cardinalMerchantID cardinalTransactionPw cardinalTransactionURL cardinalAPIKey cardinalOrgUnitID"`
	CardinalAPIKey         *string           `json:"cardinalAPIKey" yaml:"cardinalAPIKey" validate:"required_with=cardinalProcessorID cardinalMerchantID cardinalTransactionPw cardinalTransactionURL cardinalAPIIdentifier cardinalOrgUnitID"`
	CardinalOrgUnitID      *string           `json:"cardinalOrgUnitID" yaml:"cardinalOrgUnitID" validate:"required_with=cardinalProcessorID cardinalMerchantID cardinalTransactionPw cardinalTransactionURL cardinalAPIIdentifier cardinalAPIKey"`
	Environment            PayPalEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox live"`
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
