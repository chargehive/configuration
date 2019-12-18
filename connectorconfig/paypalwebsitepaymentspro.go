package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type PayPalWebsitePaymentsProCredentials struct {
	APIUsername            *string
	APIPassword            *string
	APISignature           *string
	SupportedCurrencies    []string
	CardinalProcessorID    *string
	CardinalMerchantID     *string
	CardinalTransactionPw  *string
	CardinalTransactionURL *string
	CardinalAPIIdentifier  *string
	CardinalAPIKey         *string
	CardinalOrgUnitID      *string
	Environment            PayPalEnvironment
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
