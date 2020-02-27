package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type CyberSourceEnvironment string

const (
	CyberSourceEnvironmentTest CyberSourceEnvironment = "test"
	CyberSourceEnvironmentLive CyberSourceEnvironment = "live"
)

type CyberSourceCredentials struct {
	MerchantID     string                 `json:"merchantID" yaml:"merchantID" validate:"required"`
	TransactionKey *string                `json:"transactionKey" yaml:"transactionKey" validate:"required"`
	Environment    CyberSourceEnvironment `json:"environment" yaml:"environment" validate:"oneof=test live"`
}

func (c CyberSourceCredentials) GetLibrary() Library {
	return LibraryCyberSource
}

func (c *CyberSourceCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeFraud}
}

func (c *CyberSourceCredentials) Validate() error {
	return nil
}

func (c *CyberSourceCredentials) GetSecureFields() []*string {
	return []*string{c.TransactionKey}
}

func (c *CyberSourceCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *CyberSourceCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
