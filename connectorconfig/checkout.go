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
	PublicKey              *string              `json:"publicKey" yaml:"publicKey" validate:"required,gt=0"`
	SecretKey              *string              `json:"secretKey" yaml:"secretKey" validate:"required,gt=0"`
	AuthorizationHeaderKey *string              `json:"authorizationHeaderKey" yaml:"authorizationHeaderKey" validate:"required,gt=0"`
	SignatureKey           *string              `json:"signatureKey" yaml:"signatureKey" validate:"required,gt=0"`
	Platform               *string              `json:"platform" yaml:"platform" validate:"required,oneof=default previous"`
	ProcessingChannelID    string               `json:"processingChannelId" yaml:"processingChannelId"`
	MerchantID             string               `json:"merchantID" yaml:"merchantID" validate:"required"`
	Currency               string               `json:"currency" yaml:"currency" validate:"oneof=AED AFN ALL AMD ANG AOA ARS AUD AWG AZN BAM BBD BDT BGN BHD BIF BMD BND BOB BRL BSD BTN BWP BYN BZD CAD CDF CHF CLF CLP CNY COP CRC CUP CVE CZK DJF DKK DOP DZD EEK EGP ERN ETB EUR FJD FKP GBP GEL GHS GIP GMD GNF GTQ GYD HKD HNL HRK HTG HUF IDR ILS INR IQD IRR ISK JMD JOD JPY KES KGS KHR KMF KPW KRW KWD KYD KZT LAK LBP LKR LRD LSL LTL LVL LYD MAD MDL MGA MKD MMK MNT MOP MRO MUR MVR MWK MXN MYR MZN NAD NGN NIO NOK NPR NZD OMR PAB PEN PGK PHP PKR PLN PYG QAR RON RSD RUB RWF SAR SBD SCR SDG SEK SGD SHP SLL SOS SRD STD SVC SYP SZL THB TJS TMT TND TOP TRY TTD TWD TZS UAH UGX USD UYU UZS VEF VND VUV WST XAF XCD XOF XPF YER ZAR ZMW ZWL"`
	Environment            CheckoutEnvironment  `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
	GooglePay              *GooglePayCredential `json:"googlePay,omitempty" yaml:"googlePay,omitempty"`
	ApplePay               *ApplePayCredential  `json:"applePay,omitempty" yaml:"applePay,omitempty"`
}

func (c *CheckoutCredentials) GetGooglePayParams() map[string]string {
	return map[string]string{
		"gateway":           "checkoutltd",
		"gatewayMerchantId": c.GetGooglePay().GetGoogleCardMerchantId(),
	}
}

func (c *CheckoutCredentials) GetMID() string {
	return c.MerchantID
}

func (c *CheckoutCredentials) GetGooglePay() *GooglePayCredential {
	return c.GooglePay
}

func (c *CheckoutCredentials) GetApplePay() *ApplePayCredential {
	return c.ApplePay
}

func (c *CheckoutCredentials) GetPublicKey() string {
	if c.PublicKey == nil {
		return ""
	}
	return *c.PublicKey
}

func (c *CheckoutCredentials) GetSecretKey() string {
	if c.SecretKey == nil {
		return ""
	}
	return *c.SecretKey
}

func (c *CheckoutCredentials) GetAuthorizationHeaderKey() string {
	if c.AuthorizationHeaderKey == nil {
		return ""
	}
	return *c.AuthorizationHeaderKey
}

func (c *CheckoutCredentials) GetSignatureKey() string {
	if c.SignatureKey == nil {
		return ""
	}
	return *c.SignatureKey
}

func (c *CheckoutCredentials) GetPlatform() string {
	if c.Platform == nil {
		return "previous"
	}
	return *c.Platform
}

func (c *CheckoutCredentials) GetLibrary() Library {
	return LibraryCheckout
}

func (c *CheckoutCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *CheckoutCredentials) Validate() error {
	return nil
}

func (c *CheckoutCredentials) GetSecureFields() []*string {
	fields := []*string{c.PublicKey, c.SecretKey, c.AuthorizationHeaderKey, c.SignatureKey}
	fields = append(fields, c.GetGooglePay().GetSecureFields()...)
	fields = append(fields, c.GetApplePay().GetSecureFields()...)
	return fields
}

func (c *CheckoutCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *CheckoutCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *CheckoutCredentials) SupportsSca() bool {
	return c.GetPublicKey() != "" && c.GetSecretKey() != "" && c.Environment != ""
}

func (c *CheckoutCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
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

func (c *CheckoutCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *CheckoutCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeSandbox && c.Environment == CheckoutEnvironmentProduction {
		return false
	}
	return true
}

func (c *CheckoutCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *CheckoutCredentials) Supports3RI() bool {
	return false
}
