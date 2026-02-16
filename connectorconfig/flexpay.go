package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/grpc/cht"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type FlexPayEnvironment string

const (
	FlexPayEnvironmentSandbox    FlexPayEnvironment = "sandbox"
	FlexPayEnvironmentProduction FlexPayEnvironment = "production"
)

type FlexPayCredentials struct {
	ApiKey      *string            `json:"apiKey" yaml:"apiKey" validate:"required"`
	Environment FlexPayEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
}

func (c *FlexPayCredentials) GetMID() string {
	return *c.ApiKey
}

func (c *FlexPayCredentials) GetLibrary() Library {
	return LibraryFlexPay
}

func (c *FlexPayCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeRecoveryAgent}
}

func (c *FlexPayCredentials) Validate() error {
	return nil
}

func (c *FlexPayCredentials) GetSecureFields() []*string {
	return []*string{c.ApiKey}
}

func (c *FlexPayCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *FlexPayCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *FlexPayCredentials) SupportsSca() bool {
	return false
}

func (c *FlexPayCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *FlexPayCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *FlexPayCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeSandbox && c.Environment == FlexPayEnvironmentProduction {
		return false
	}
	return true
}

func (c *FlexPayCredentials) IsRecoveryAgent() bool {
	return true
}

func (c *FlexPayCredentials) Supports3RI() bool {
	return false
}

func (c *FlexPayCredentials) IsAccountUpdater() bool {
	return false
}

func (c *FlexPayCredentials) SupportedTokenSources() []cht.TokenSource {
	return []cht.TokenSource{cht.TS_PAN}
}
