package connectorconfig

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type ChargeHiveCredentials struct {
	GooglePay *GooglePayCredential `json:"googlePay,omitempty" yaml:"googlePay,omitempty"`
	ApplePay  *ApplePayCredential  `json:"applePay,omitempty" yaml:"applePay,omitempty"`
}

func (c *ChargeHiveCredentials) GetMID() string {
	return ""
}

func (c *ChargeHiveCredentials) GetLibrary() Library {
	return LibraryChargeHive
}

func (c *ChargeHiveCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *ChargeHiveCredentials) Validate() error {
	if c.GooglePay != nil {
		if !c.GooglePay.IsValid() {
			return errors.New("invalid google pay configuration")
		}
	}
	if c.ApplePay != nil {
		if !c.ApplePay.IsValid() {
			return errors.New("invalid apple pay configuration")
		}
	}
	return nil
}

func (c *ChargeHiveCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *ChargeHiveCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *ChargeHiveCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *ChargeHiveCredentials) SupportsSca() bool {
	return false
}

func (c *ChargeHiveCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *ChargeHiveCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *ChargeHiveCredentials) CanPlanModeUse(environment.Mode) bool {
	return true
}

func (c *ChargeHiveCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *ChargeHiveCredentials) Supports3RI() bool {
	return false
}

func (c *ChargeHiveCredentials) GetApplePay() *ApplePayCredential {
	return c.ApplePay
}

func (c *ChargeHiveCredentials) GetGooglePay() *GooglePayCredential {
	return c.GooglePay
}

func (c *ChargeHiveCredentials) GetGooglePayParams() map[string]string {
	return map[string]string{
		"gateway":           "chargehive",
		"gatewayMerchantId": c.GetGooglePay().GetGoogleCardMerchantId(),
	}
}
