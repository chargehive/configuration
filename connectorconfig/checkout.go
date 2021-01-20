package connectorconfig

import (
	"encoding/json"
	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type CheckoutEnvironment string

const (
	CheckoutEnvironmentSandbox    CheckoutEnvironment = "sandbox"
	CheckoutEnvironmentProduction CheckoutEnvironment = "production"
)

type CheckoutCredentials struct {
	PublicKey   *string             `json:"publicKey" yaml:"publicKey" validate:"required,gt=0"`
	SecretKey   *string             `json:"secretKey" yaml:"secretKey" validate:"required,gt=0"`
	Environment CheckoutEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
}

func (c CheckoutCredentials) GetPublicKey() string {
	if c.PublicKey == nil {
		return ""
	}
	return *c.PublicKey
}

func (c CheckoutCredentials) GetSecretKey() string {
	if c.SecretKey == nil {
		return ""
	}
	return *c.SecretKey
}

func (c CheckoutCredentials) GetLibrary() Library {
	return LibraryCheckout
}

func (c CheckoutCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *CheckoutCredentials) Validate() error {
	return nil
}

func (c *CheckoutCredentials) GetSecureFields() []*string {
	return []*string{c.PublicKey, c.SecretKey}
}

func (c *CheckoutCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *CheckoutCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c CheckoutCredentials) SupportsSca() bool {
	return c.GetPublicKey() != "" && c.GetSecretKey() != "" && c.Environment != ""
}

func (c CheckoutCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if methodType == chtype.PAYMENT_METHOD_TYPE_CARD {
		return true
	}
	return false
}

func (c CheckoutCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeSandbox && c.Environment == CheckoutEnvironmentProduction {
		return false
	}
	return true
}
