package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type PaySafeGooglePayCredentials struct {
	Acquirer               string
	AccountID              string
	APIUsername            *string
	APIPassword            *string
	Environment            PaySafeEnvironment
	Country                string
	Currency               string
	SingleUseTokenUsername *string
	SingleUseTokenPassword *string
	Locale                 PaysafeLocale

	// Pending implementation <additional parameters TBC>
}

func (c PaySafeGooglePayCredentials) GetLibrary() Library {
	return LibraryPaySafeGooglePay
}

func (c *PaySafeGooglePayCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *PaySafeGooglePayCredentials) Validate() error {
	return nil
}

func (c *PaySafeGooglePayCredentials) GetSecureFields() []*string {
	return []*string{c.APIUsername, c.APIPassword, c.SingleUseTokenUsername, c.SingleUseTokenPassword}
}

func (c *PaySafeGooglePayCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *PaySafeGooglePayCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
