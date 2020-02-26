package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type QualPayEnvironment string

const (
	QualPayEnvironmentTest QualPayEnvironment = "test"
	QualPayEnvironmentLive QualPayEnvironment = "live"
)

type QualpayCredentials struct {
	APIKey      *string            `json:"apiKey" yaml:"apiKey" validate:"required"`
	MerchantID  int64              `json:"merchantID" yaml:"merchantID" validate:"required"`
	Environment QualPayEnvironment `json:"environment" yaml:"environment" validate:"required"`
}

func (c QualpayCredentials) GetLibrary() Library {
	return LibraryQualPay
}

func (c QualpayCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *QualpayCredentials) Validate() error {
	return nil
}

func (c *QualpayCredentials) GetSecureFields() []*string {
	return []*string{c.APIKey}
}

func (c *QualpayCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *QualpayCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
