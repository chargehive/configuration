package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type TokenExNetworkTokenizationCredentials struct {
	Environment TokenExEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
	Region      TokenExRegion      `json:"region" yaml:"region" validate:"oneof=us eu"`
	TokenExID   string             `json:"tokenExID,omitempty"`
	APIKey      *string            `json:"apiKey,omitempty"`
}

func (c *TokenExNetworkTokenizationCredentials) GetLibrary() Library {
	return LibraryTokenExNetworkTokenization
}

func (c *TokenExNetworkTokenizationCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeNetworkTokenization}
}

func (c *TokenExNetworkTokenizationCredentials) Validate() error {
	return nil
}

func (c *TokenExNetworkTokenizationCredentials) GetSecureFields() []*string {
	return []*string{c.APIKey}
}

func (c *TokenExNetworkTokenizationCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *TokenExNetworkTokenizationCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *TokenExNetworkTokenizationCredentials) SupportsSca() bool {
	return false
}

func (c *TokenExNetworkTokenizationCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *TokenExNetworkTokenizationCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *TokenExNetworkTokenizationCredentials) CanPlanModeUse(environment.Mode) bool {
	return true
}

func (c *TokenExNetworkTokenizationCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *TokenExNetworkTokenizationCredentials) Supports3RI() bool {
	return false
}

func (c *TokenExNetworkTokenizationCredentials) IsAccountUpdater() bool {
	return false
}
