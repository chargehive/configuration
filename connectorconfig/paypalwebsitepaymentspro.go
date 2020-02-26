package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type PayPalWebsitePaymentsProCredentials struct {
	APIUsername            *string           `json:"apiUsername" yaml:"apiUsername" validate:"required"`
	APIPassword            *string           `json:"apiPassword" yaml:"apiPassword" validate:"required"`
	APISignature           *string           `json:"apiSignature" yaml:"apiSignature" validate:"required"`
	SupportedCurrencies    []string          `json:"supportedCurrencies" yaml:"supportedCurrencies" validate:"required"`
	CardinalProcessorID    *string           `json:"cardinalProcessorID" yaml:"cardinalProcessorID"`
	CardinalMerchantID     *string           `json:"cardinalMerchantID" yaml:"cardinalMerchantID"`
	CardinalTransactionPw  *string           `json:"cardinalTransactionPw" yaml:"cardinalTransactionPw"`
	CardinalTransactionURL *string           `json:"cardinalTransactionURL" yaml:"cardinalTransactionURL"`
	CardinalAPIIdentifier  *string           `json:"cardinalAPIIdentifier" yaml:"cardinalAPIIdentifier"`
	CardinalAPIKey         *string           `json:"cardinalAPIKey" yaml:"cardinalAPIKey"`
	CardinalOrgUnitID      *string           `json:"cardinalOrgUnitID" yaml:"cardinalOrgUnitID"`
	Environment            PayPalEnvironment `json:"environment" yaml:"environment" validate:"required"`
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
