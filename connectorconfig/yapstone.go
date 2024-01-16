package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type YapstoneEnvironment string

const (
	YapstoneEnvironmentTest YapstoneEnvironment = "test"
	YapstoneEnvironmentLive YapstoneEnvironment = "live"
)

type YapstoneCredentials struct {
	ClientID     string              `json:"clientID" yaml:"clientID" validate:"required"`
	ClientSecret string              `json:"clientSecret" yaml:"clientSecret" validate:"required"`
	Environment  YapstoneEnvironment `json:"environment" yaml:"environment" validate:"oneof=test live"`
}

func (c *YapstoneCredentials) GetMID() string {
	return c.ClientID
}

func (c *YapstoneCredentials) GetLibrary() Library {
	return LibraryYapstone
}

func (c *YapstoneCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *YapstoneCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *YapstoneCredentials) Validate() error {
	return nil
}

func (c *YapstoneCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *YapstoneCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *YapstoneCredentials) SupportsSca() bool {
	return false
}

func (c *YapstoneCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *YapstoneCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *YapstoneCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeProduction && c.Environment == YapstoneEnvironmentLive {
		return true
	}

	if mode == environment.ModeSandbox && c.Environment == YapstoneEnvironmentTest {
		return true
	}

	return false
}

func (c *YapstoneCredentials) IsRecoveryAgent() bool {
	return false
}
