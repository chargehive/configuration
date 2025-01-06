package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type TokenExApiAccountUpdaterCredentials struct {
	TokenExID  string        `json:"tokenExID" yaml:"tokenExID" validate:"required,gt=0"`
	ApiKey     *string       `json:"apiKey" yaml:"apiKey" validate:"required,gt=0"`
	MerchantID string        `json:"merchantID" yaml:"merchantID" validate:"required,gt=0"`
	Region     TokenExRegion `json:"region" yaml:"region" validate:"required,oneof=us eu"`

	Environment TokenExEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
}

func (c *TokenExApiAccountUpdaterCredentials) GetMID() string {
	return c.TokenExID
}

func (c *TokenExApiAccountUpdaterCredentials) GetLibrary() Library {
	return LibraryTokenExApiAccountUpdater
}

func (c *TokenExApiAccountUpdaterCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeMethodUpdater}
}

func (c *TokenExApiAccountUpdaterCredentials) Validate() error {
	return nil
}

func (c *TokenExApiAccountUpdaterCredentials) GetSecureFields() []*string {
	return []*string{c.ApiKey}
}

func (c *TokenExApiAccountUpdaterCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *TokenExApiAccountUpdaterCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *TokenExApiAccountUpdaterCredentials) SupportsSca() bool {
	return false
}

func (c *TokenExApiAccountUpdaterCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *TokenExApiAccountUpdaterCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *TokenExApiAccountUpdaterCredentials) CanPlanModeUse(environment.Mode) bool {
	return true
}

func (c *TokenExApiAccountUpdaterCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *TokenExApiAccountUpdaterCredentials) Supports3RI() bool {
	return false
}
