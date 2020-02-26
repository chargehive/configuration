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

type PaysafeLocale string

const (
	PaysafeLocaleENGB PaysafeLocale = "en_GB"
	PaysafeLocaleENUS PaysafeLocale = "en_US"
	PaysafeLocaleFRCA PaysafeLocale = "fr_CA"
)

type PaySafeCredentials struct {
	Acquirer               string             `json:"acquirer" yaml:"acquirer" validate:"required"`
	AccountID              string             `json:"accountID" yaml:"accountID" validate:"required"`
	APIUsername            *string            `json:"apiUsername" yaml:"apiUsername" validate:"required"`
	APIPassword            *string            `json:"apiPassword" yaml:"apiPassword" validate:"required"`
	Environment            PaySafeEnvironment `json:"environment" yaml:"environment" validate:"required"`
	Country                string             `json:"country" yaml:"country"`
	Currency               string             `json:"currency" yaml:"currency"`
	UseVault               bool               `json:"useVault" yaml:"useVault"`
	SingleUseTokenPassword *string            `json:"singleUseTokenPassword" yaml:"singleUseTokenPassword"`
	SingleUseTokenUsername string             `json:"singleUseTokenUsername" yaml:"singleUseTokenUsername"`
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
