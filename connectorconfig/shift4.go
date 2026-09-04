package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/configuration/v1/scheduler"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

var (
	_ ApplePayEmbeddedCredential  = (*Shift4Credentials)(nil)
	_ GooglePayEmbeddedCredential = (*Shift4Credentials)(nil)
	_ Credentials                 = (*Shift4Credentials)(nil)
	_ MerchantIdentifier          = (*Shift4Credentials)(nil)
)

type Shift4Environment string

const (
	Shift4EnvironmentIntegration Shift4Environment = "integration"
	Shift4EnvironmentProduction  Shift4Environment = "production"
)

type Shift4Credentials struct {
	MerchantID         *string               `json:"merchantId" yaml:"merchantId" validate:"required,gt=0"`
	SignatureKey       *string               `json:"signatureKey" yaml:"signatureKey" validate:"required,gt=0"`
	Environment        Shift4Environment     `json:"environment" yaml:"environment" validate:"oneof=integration production"`
	MerchantDescriptor string                `json:"merchantDescriptor,omitempty" yaml:"merchantDescriptor,omitempty" validate:"omitempty,lte=25"`
	GooglePay          *GooglePayCredentials `json:"googlePay,omitempty" yaml:"googlePay,omitempty"`
	ApplePay           *ApplePayCredentials  `json:"applePay,omitempty" yaml:"applePay,omitempty"`
}

func (c *Shift4Credentials) GetGooglePayParams() map[string]string {
	return map[string]string{
		"gateway":           "shift4",
		"gatewayMerchantId": c.GetGooglePay().GetGoogleCardMerchantId(),
	}
}

func (c *Shift4Credentials) GetGooglePay() *GooglePayCredentials { return c.GooglePay }
func (c *Shift4Credentials) GetApplePay() *ApplePayCredentials   { return c.ApplePay }

func (c *Shift4Credentials) GetMID() string {
	if c.MerchantID == nil {
		return ""
	}
	return *c.MerchantID
}

func (c *Shift4Credentials) GetLibrary() Library { return LibraryShift4 }

func (c *Shift4Credentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *Shift4Credentials) Validate() error { return nil }

func (c *Shift4Credentials) GetSecureFields() []*string {
	fields := []*string{c.SignatureKey}
	fields = append(fields, c.GetGooglePay().GetSecureFields()...)
	fields = append(fields, c.GetApplePay().GetSecureFields()...)
	return fields
}

func (c *Shift4Credentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *Shift4Credentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *Shift4Credentials) SupportsSca() bool {
	// No built-in 3DS in v1; external SCA results pass through on i8.
	return false
}

func (c *Shift4Credentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
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

func (c *Shift4Credentials) SupportsCountry(country string) bool { return true }

func (c *Shift4Credentials) CanPlanModeUse(mode environment.Mode) bool {
	return mode == environment.ModeSandbox && c.Environment == Shift4EnvironmentIntegration ||
		mode == environment.ModeProduction && c.Environment == Shift4EnvironmentProduction
}

func (c *Shift4Credentials) IsRecoveryAgent() bool { return false }

func (c *Shift4Credentials) Supports3RI() bool { return false }

func (c *Shift4Credentials) IsAccountUpdater() bool { return false }

func (c *Shift4Credentials) SupportedTokenTypes() []scheduler.TokenSource {
	return []scheduler.TokenSource{scheduler.TokenSourcePan, scheduler.TokenSourceConnector, scheduler.TokenSourceNetworkToken}
}
