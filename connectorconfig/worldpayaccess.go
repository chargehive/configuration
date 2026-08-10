package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/configuration/v1/scheduler"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

var (
	_ ApplePayEmbeddedCredential  = (*WorldpayAccessCredentials)(nil)
	_ GooglePayEmbeddedCredential = (*WorldpayAccessCredentials)(nil)
	_ Credentials                 = (*WorldpayAccessCredentials)(nil)
	_ MerchantIdentifier          = (*WorldpayAccessCredentials)(nil)
)

type WorldpayAccessEnvironment string

const (
	WorldpayAccessEnvironmentTry        WorldpayAccessEnvironment = "try"
	WorldpayAccessEnvironmentProduction WorldpayAccessEnvironment = "production"
)

type WorldpayAccessCredentials struct {
	Username           *string                   `json:"username" yaml:"username" validate:"required,gt=0"`
	Password           *string                   `json:"password" yaml:"password" validate:"required,gt=0"`
	Environment        WorldpayAccessEnvironment `json:"environment" yaml:"environment" validate:"oneof=try production"`
	Entity             string                    `json:"entity" yaml:"entity" validate:"gte=1,lte=32"`
	MerchantDescriptor string                    `json:"merchantDescriptor" yaml:"merchantDescriptor" validate:"gte=1,lte=24"`
	MCC                string                    `json:"mcc,omitempty" yaml:"mcc,omitempty" validate:"omitempty,len=4,numeric"`
	GooglePay          *GooglePayCredentials     `json:"googlePay,omitempty" yaml:"googlePay,omitempty"`
	ApplePay           *ApplePayCredentials      `json:"applePay,omitempty" yaml:"applePay,omitempty"`
}

func (c *WorldpayAccessCredentials) GetGooglePayParams() map[string]string {
	return map[string]string{
		"gateway":           "worldpay-access",
		"gatewayMerchantId": c.GetGooglePay().GetGoogleCardMerchantId(),
	}
}

func (c *WorldpayAccessCredentials) GetGooglePay() *GooglePayCredentials {
	return c.GooglePay
}

func (c *WorldpayAccessCredentials) GetApplePay() *ApplePayCredentials {
	return c.ApplePay
}

func (c *WorldpayAccessCredentials) GetMID() string {
	return c.Entity
}

func (c *WorldpayAccessCredentials) GetLibrary() Library {
	return LibraryWorldpayAccess
}

func (c *WorldpayAccessCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *WorldpayAccessCredentials) Validate() error {
	return nil
}

func (c *WorldpayAccessCredentials) GetSecureFields() []*string {
	fields := []*string{c.Username, c.Password}
	fields = append(fields, c.GetGooglePay().GetSecureFields()...)
	fields = append(fields, c.GetApplePay().GetSecureFields()...)
	return fields
}

func (c *WorldpayAccessCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *WorldpayAccessCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *WorldpayAccessCredentials) SupportsSca() bool {
	// The connector has no built-in 3DS; external SCA results are passed
	// through to Worldpay on authorization.
	return false
}

func (c *WorldpayAccessCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	if methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY {
		return c.GetApplePay().IsValid()
	}
	if methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY {
		return c.GetGooglePay().IsValid()
	}
	return true
}

func (c *WorldpayAccessCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *WorldpayAccessCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return mode == environment.ModeSandbox && c.Environment == WorldpayAccessEnvironmentTry ||
		mode == environment.ModeProduction && c.Environment == WorldpayAccessEnvironmentProduction
}

func (c *WorldpayAccessCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *WorldpayAccessCredentials) Supports3RI() bool {
	return false
}

func (c *WorldpayAccessCredentials) IsAccountUpdater() bool {
	return false
}

func (c *WorldpayAccessCredentials) SupportedTokenTypes() []scheduler.TokenSource {
	return []scheduler.TokenSource{scheduler.TokenSourcePan, scheduler.TokenSourceConnector, scheduler.TokenSourceNetworkToken}
}
