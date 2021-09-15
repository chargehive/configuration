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
	Currency    string              `json:"currency" yaml:"currency" validate:"oneof=AED AFN ALL AMD ANG AOA ARS AUD AWG AZN BAM BBD BDT BGN BHD BIF BMD BND BOB BRL BSD BTN BWP BYN BZD CAD CDF CHF CLF CLP CNY COP CRC CUP CVE CZK DJF DKK DOP DZD EEK EGP ERN ETB EUR FJD FKP GBP GEL GHS GIP GMD GNF GTQ GYD HKD HNL HRK HTG HUF IDR ILS INR IQD IRR ISK JMD JOD JPY KES KGS KHR KMF KPW KRW KWD KYD KZT LAK LBP LKR LRD LSL LTL LVL LYD MAD MDL MGA MKD MMK MNT MOP MRO MUR MVR MWK MXN MYR MZN NAD NGN NIO NOK NPR NZD OMR PAB PEN PGK PHP PKR PLN PYG QAR RON RSD RUB RWF SAR SBD SCR SDG SEK SGD SHP SLL SOS SRD STD SVC SYP SZL THB TJS TMT TND TOP TRY TTD TWD TZS UAH UGX USD UYU UZS VEF VND VUV WST XAF XCD XOF XPF YER ZAR ZMW ZWL"`
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
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}

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

func (c CheckoutCredentials) IsRecoveryAgent() bool {
	return false
}
