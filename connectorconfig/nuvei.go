package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type NuveiEnvironment string

const (
	NuveiEnvironmentSandbox    NuveiEnvironment = "sandbox"
	NuveiEnvironmentProduction NuveiEnvironment = "production"
)

type NuveiCredentials struct {
	MerchantID        *string          `json:"merchantID" yaml:"merchantID" validate:"required,gt=0"`
	MerchantSiteID    *string          `json:"merchantSiteID" yaml:"merchantSiteID" validate:"required,gt=0"`
	MerchantSecretKey *string          `json:"merchantSecretKey" yaml:"merchantSecretKey" validate:"required,gt=0"`
	Environment       NuveiEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
}

func (c *NuveiCredentials) GetMID() string {
	return *c.MerchantID
}

func (c *NuveiCredentials) GetLibrary() Library {
	return LibraryNuvei
}

func (c *NuveiCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *NuveiCredentials) Validate() error {
	return nil
}

func (c *NuveiCredentials) GetSecureFields() []*string {
	return []*string{c.MerchantSecretKey}
}

func (c *NuveiCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *NuveiCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *NuveiCredentials) SupportsSca() bool {
	return true
}

func (c *NuveiCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *NuveiCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *NuveiCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeSandbox && c.Environment == NuveiEnvironmentProduction {
		return false
	}
	return true
}

func (c *NuveiCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *NuveiCredentials) Supports3RI() bool {
	return false
}

func (c *NuveiCredentials) IsAccountUpdater() bool {
	return false
}
