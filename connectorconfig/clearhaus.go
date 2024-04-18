package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type ClearhausEnvironment string

const (
	ClearhausEnvironmentTest ClearhausEnvironment = "test"
	ClearhausEnvironmentLive ClearhausEnvironment = "live"
)

type ClearhausCredentials struct {
	MerchantID         string               `json:"merchantId" yaml:"merchantId" validate:"required"`
	MerchantDescriptor string               `json:"merchantDescriptor" yaml:"merchantDescriptor" validate:"-"`
	Environment        ClearhausEnvironment `json:"environment" yaml:"environment" validate:"required,oneof=test live"`

	APIKey            *string `json:"apiKey" yaml:"apiKey" validate:"required"`
	SigningApiKey     *string `json:"signingApiKey" yaml:"signingApiKey" validate:"required"`
	SigningPrivateKey *string `json:"signingPrivateKey" yaml:"signingPrivateKey" validate:"required"`
}

func (c *ClearhausCredentials) GetMID() string {
	return c.MerchantID
}

func (c *ClearhausCredentials) GetLibrary() Library {
	return LibraryClearhaus
}

func (c *ClearhausCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *ClearhausCredentials) GetSecureFields() []*string {
	return []*string{c.APIKey, c.SigningApiKey, c.SigningPrivateKey}
}

func (c *ClearhausCredentials) Validate() error {
	return nil
}

func (c *ClearhausCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *ClearhausCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *ClearhausCredentials) SupportsSca() bool {
	return false
}

func (c *ClearhausCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *ClearhausCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *ClearhausCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeProduction && c.Environment == ClearhausEnvironmentLive {
		return true
	}

	if mode == environment.ModeSandbox && c.Environment == ClearhausEnvironmentTest {
		return true
	}

	return false
}

func (c *ClearhausCredentials) IsRecoveryAgent() bool {
	return false
}
