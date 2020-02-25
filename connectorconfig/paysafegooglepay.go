package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type PaySafeGooglePayCredentials struct {
	Acquirer               string             `json:"acquirer" yaml:"acquirer"`
	AccountID              string             `json:"accountID" yaml:"accountID"`
	APIUsername            *string            `json:"apiUsername" yaml:"apiUsername"`
	APIPassword            *string            `json:"apiPassword" yaml:"apiPassword"`
	Environment            PaySafeEnvironment `json:"environment" yaml:"environment"`
	Country                string             `json:"country" yaml:"country"`
	Currency               string             `json:"currency" yaml:"currency"`
	SingleUseTokenUsername *string            `json:"singleUseTokenUsername" yaml:"singleUseTokenUsername"`
	SingleUseTokenPassword *string            `json:"singleUseTokenPassword" yaml:"singleUseTokenPassword"`
	Locale                 PaysafeLocale      `json:"locale" yaml:"locale"`

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
