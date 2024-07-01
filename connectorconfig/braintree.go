package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type BraintreeEnvironment string

const (
	BraintreeEnvironmentSandbox    BraintreeEnvironment = "sandbox"
	BraintreeEnvironmentProduction BraintreeEnvironment = "production"

	braintreeSdkVersion = "3.97.3"
)

// https://articles.braintreepayments.com/control-panel/important-gateway-credentials
type BraintreeCredentials struct {
	PublicKey         *string               `json:"publicKey" yaml:"publicKey" validate:"required,gt=0"`
	PrivateKey        *string               `json:"privateKey" yaml:"privateKey" validate:"required,gt=0"`
	MerchantID        string                `json:"merchantID" yaml:"merchantID" validate:"required"`
	MerchantAccountID string                `json:"merchantAccountID" yaml:"merchantAccountID" validate:"-"`
	Currency          string                `json:"currency" yaml:"currency" validate:"oneof=AED AMD AOA ARS AUD AWG AZN BAM BBD BDT BGN BIF BMD BND BOB BRL BSD BWP BYN BZD CAD CHF CLP CNY COP CRC CVE CZK DJF DKK DOP DZD EGP ETB EUR FJD FKP GBP GEL GHS GIP GMD GNF GTQ GYD HKD HNL HRK HTG HUF IDR ILS INR ISK JMD JPY KES KGS KHR KMF KRW KYD KZT LAK LBP LKR LRD LSL LTL MAD MDL MKD MNT MOP MUR MVR MWK MXN MYR MZN NAD NGN NIO NOK NPR NZD PAB PEN PGK PHP PKR PLN PYG QAR RON RSD RUB RWF SAR SBD SCR SEK SGD SHP SLL SOS SRD STD SVC SYP SZL THB TJS TOP TRY TTD TWD TZS UAH UGX USD UYU UZS VES VND VUV WST XAF XCD XOF XPF YER ZAR ZMK ZWD"`
	Environment       BraintreeEnvironment  `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
	GooglePay         *GooglePayCredentials `json:"googlePay,omitempty" yaml:"googlePay,omitempty"`
	ApplePay          *ApplePayCredentials  `json:"applePay,omitempty" yaml:"applePay,omitempty"`
	TokenizationKey   string                `json:"tokenizationKey,omitempty" yaml:"tokenizationKey,omitempty" validate:"required_with=GooglePayEmbedded ApplePay,omitempty,gt=0"`
}

func (c *BraintreeCredentials) GetMID() string {
	return c.MerchantAccountID
}

func (c *BraintreeCredentials) GetPublicKey() string {
	if c.PublicKey == nil {
		return ""
	}
	return *c.PublicKey
}

func (c *BraintreeCredentials) GetPrivateKey() string {
	if c.PrivateKey == nil {
		return ""
	}
	return *c.PrivateKey
}

func (c *BraintreeCredentials) GetLibrary() Library {
	return LibraryBraintree
}

func (c *BraintreeCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *BraintreeCredentials) Validate() error {
	return nil
}

func (c *BraintreeCredentials) GetSecureFields() []*string {
	fields := []*string{c.PublicKey, c.PrivateKey}
	fields = append(fields, c.GetGooglePay().GetSecureFields()...)
	fields = append(fields, c.GetApplePay().GetSecureFields()...)
	return fields
}

func (c *BraintreeCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *BraintreeCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *BraintreeCredentials) SupportsSca() bool {
	return c.MerchantID != "" && c.GetPublicKey() != "" && c.GetPrivateKey() != "" && c.Environment != ""
}

func (c *BraintreeCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
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

func (c *BraintreeCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *BraintreeCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeSandbox && c.Environment == BraintreeEnvironmentProduction {
		return false
	}
	return true
}

func (c *BraintreeCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *BraintreeCredentials) GetGooglePay() *GooglePayCredentials {
	return c.GooglePay
}

func (c *BraintreeCredentials) GetGooglePayParams() map[string]string {
	return map[string]string{
		"gateway":              "braintree",
		"braintree:apiVersion": "v1",
		"braintree:sdkVersion": braintreeSdkVersion,
		"braintree:merchantId": c.MerchantID,
		"braintree:clientKey":  c.TokenizationKey,
	}
}

func (c *BraintreeCredentials) GetApplePay() *ApplePayCredentials {
	return c.ApplePay
}

func (c *BraintreeCredentials) Supports3RI() bool {
	return false
}
