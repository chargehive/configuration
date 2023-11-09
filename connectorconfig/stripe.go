package connectorconfig

import (
	"encoding/json"
	"github.com/chargehive/configuration/paymentmethod"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

// assert interface compliance
var _ Credentials = (*StripeCredentials)(nil)

type StripeCredentials struct {
	AccountID string  `json:"accountId" yaml:"accountId"`
	APIKey    *string `json:"apiKey" yaml:"apiKey" validate:"required,gt=0"`
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

func (c *StripeCredentials) CanPlanModeUse(environment.Mode) bool {
	// todo will require updating when we have test credentials
	return true
}

func (c *StripeCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *StripeCredentials) SupportedSchemes() []paymentmethod.Scheme {
	return []paymentmethod.Scheme{
		paymentmethod.SchemeCardVisa,
		paymentmethod.SchemeCardMasterCard,
		paymentmethod.SchemeCardAmericanExpress,
		paymentmethod.SchemeCardDiscover,
		paymentmethod.SchemeCardDinersClub,
		paymentmethod.SchemeCardUnionPay,
		paymentmethod.SchemeCardJCB,
		paymentmethod.SchemeCardCarteBancaire,
	}
}

func (c *StripeCredentials) SupportsNetworkToken() bool {
	return false
}
