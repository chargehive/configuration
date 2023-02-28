package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type CWAMSCredentials struct {
	TestMode    bool       `json:"testMode" yaml:"testMode"`
	SecurityKey *string    `json:"securityKey" yaml:"securityKey" validate:"required"`
	GooglePay   *GooglePay `json:"googlePay,omitempty" yaml:"googlePay,omitempty"`
	ApplePay    *ApplePay  `json:"applePay,omitempty" yaml:"applePay,omitempty"`
}

func (c *CWAMSCredentials) GetGooglePay() *GooglePay {
	return c.GooglePay
}

func (c *CWAMSCredentials) GetApplePay() *ApplePay {
	return c.ApplePay
}

func (c *CWAMSCredentials) GetLibrary() Library {
	return LibraryCWAMS
}

func (c *CWAMSCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *CWAMSCredentials) GetSecureFields() []*string {
	return []*string{c.SecurityKey}
}

func (c *CWAMSCredentials) Validate() error {
	return nil
}

func (c *CWAMSCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *CWAMSCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *CWAMSCredentials) SupportsSca() bool {
	return true
}

func (c *CWAMSCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
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

func (c *CWAMSCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeProduction && !c.TestMode {
		return true
	}

	if mode == environment.ModeSandbox && c.TestMode {
		return true
	}

	return false
}

func (c *CWAMSCredentials) IsRecoveryAgent() bool {
	return false
}
