package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type StripeCredentials struct {
	AccountID          string  `json:"accountId" yaml:"accountId"`
	MerchantDescriptor string  `json:"merchantDescriptor" yaml:"merchantDescriptor" validate:"-"`
	APIKey             *string `json:"apiKey" yaml:"apiKey" validate:"required,gt=0"`
}

func (c *StripeCredentials) GetMID() string {
	return c.AccountID
}

func (c *StripeCredentials) GetLibrary() Library {
	return LibraryStripe
}

func (c *StripeCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *StripeCredentials) Validate() error {
	return nil
}

func (c *StripeCredentials) GetSecureFields() []*string {
	return []*string{c.APIKey}
}

func (c *StripeCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *StripeCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *StripeCredentials) SupportsSca() bool {
	return false
}

func (c *StripeCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *StripeCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *StripeCredentials) CanPlanModeUse(environment.Mode) bool {
	// todo will require updating when we have test credentials
	return true
}

func (c *StripeCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *StripeCredentials) Supports3RI() bool {
	return false
}
