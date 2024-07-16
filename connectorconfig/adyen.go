package connectorconfig

import (
	"encoding/json"
	"strings"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type AdyenEnvironment string

const (
	AdyenEnvironmentSandbox    AdyenEnvironment = "sandbox"
	AdyenEnvironmentProduction AdyenEnvironment = "production"
)

type AdyenCredentials struct {
	Environment     AdyenEnvironment      `json:"environment" yaml:"environment" validate:"required,oneof=sandbox production"`
	MerchantAccount string                `json:"merchantAccount" yaml:"merchantAccount" validate:"required"`
	ApiKey          *string               `json:"apiKey" yaml:"apiKey" validate:"required"`
	ApiPrefix       string                `json:"apiPrefix" yaml:"apiPrefix" validate:"required"`
	HMACKey         *string               `json:"hmacKey" yaml:"hmacKey" validate:"required"`
	GooglePay       *GooglePayCredentials `json:"googlePay,omitempty" yaml:"googlePay,omitempty"`
	ApplePay        *ApplePayCredentials  `json:"applePay,omitempty" yaml:"applePay,omitempty"`
}

func (c *AdyenCredentials) GetMID() string {
	return c.MerchantAccount
}

func (c *AdyenCredentials) GetLibrary() Library {
	return LibraryAdyen
}

func (c *AdyenCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *AdyenCredentials) GetSecureFields() []*string {
	return []*string{c.ApiKey}
}

func (c *AdyenCredentials) Validate() error {
	return nil
}

func (c *AdyenCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *AdyenCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *AdyenCredentials) SupportsSca() bool {
	return false
}

func (c *AdyenCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

var adyenAllowedCountires = []string{"AT", "AU", "BE", "CA", "CH", "CY", "CZ", "DE", "DK", "EE", "ES", "FI", "FR", "GB", "GR", "IE", "IT", "LI", "LT", "LU", "LV", "MC", "MT", "NL", "NO", "PL", "PT", "SE", "SI", "SK", "US"}

func (c *AdyenCredentials) SupportsCountry(country string) bool {
	if country == "" {
		return true
	}
	for _, v := range adyenAllowedCountires {
		if strings.EqualFold(v, country) {
			return true
		}
	}
	return false
}

func (c *AdyenCredentials) CanPlanModeUse(mode environment.Mode) bool {

	if mode == environment.ModeSandbox && c.Environment == AdyenEnvironmentSandbox {
		return true
	}

	if mode == environment.ModeProduction && c.Environment == AdyenEnvironmentProduction {
		return true
	}

	return false
}

func (c *AdyenCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *AdyenCredentials) Supports3RI() bool {
	return false
}

func (c *AdyenCredentials) GetGooglePayParams() map[string]string {
	return map[string]string{
		"gateway":           "adyen",
		"gatewayMerchantId": c.GetGooglePay().GetGoogleCardMerchantId(),
	}
}

func (c *AdyenCredentials) GetGooglePay() *GooglePayCredentials {
	return c.GooglePay
}

func (c *AdyenCredentials) GetApplePay() *ApplePayCredentials {
	return c.ApplePay
}
