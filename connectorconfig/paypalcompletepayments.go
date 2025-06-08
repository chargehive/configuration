package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type PayPalCompletePaymentsCredentials struct {
	MerchantID string `json:"merchantId,omitempty" yaml:"merchantId,omitempty"`
}

func (c *PayPalCompletePaymentsCredentials) GetMID() string {
	return c.MerchantID
}

func (c *PayPalCompletePaymentsCredentials) GetLibrary() Library {
	return LibraryPayPalCompletePayments
}

func (c *PayPalCompletePaymentsCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *PayPalCompletePaymentsCredentials) Validate() error {
	return nil
}

func (c *PayPalCompletePaymentsCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *PayPalCompletePaymentsCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *PayPalCompletePaymentsCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *PayPalCompletePaymentsCredentials) SupportsSca() bool {
	return false
}

func (c *PayPalCompletePaymentsCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *PayPalCompletePaymentsCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *PayPalCompletePaymentsCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (c *PayPalCompletePaymentsCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *PayPalCompletePaymentsCredentials) Supports3RI() bool {
	return false
}
