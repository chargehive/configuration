package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type PaySafeEnvironment string

const (
	PaySafeEnvironmentMock PaySafeEnvironment = "MOCK"
	PaySafeEnvironmentTest PaySafeEnvironment = "TEST"
	PaySafeEnvironmentLive PaySafeEnvironment = "LIVE"
)

type PaySafeCredentials struct {
	Acquirer               string
	AccountID              string
	APIUsername            *string
	APIPassword            *string
	Environment            PaySafeEnvironment
	Country                string
	Currency               string
	UseVault               bool
	SingleUseTokenPassword *string
	SingleUseTokenUsername string
}

func (c PaySafeCredentials) GetLibrary() Library {
	return LibraryPaySafe
}

func (c *PaySafeCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *PaySafeCredentials) Validate() error {
	return nil
}

func (c *PaySafeCredentials) GetSecureFields() []*string {
	return []*string{c.APIUsername, c.APIPassword, c.SingleUseTokenPassword}
}

func (c *PaySafeCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *PaySafeCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
