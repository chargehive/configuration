package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/proto/golang/chargehive/chtype"

	"github.com/chargehive/configuration/v1/connector"
)

type BottomlineEnvironment string

const (
	BottomlineEnvironmentUAT        BottomlineEnvironment = "uat"
	BottomlineEnvironmentSandbox    BottomlineEnvironment = "sandbox"
	BottomlineEnvironmentProduction BottomlineEnvironment = "production"
)

type BottomlineCredentials struct {
	Username      *string               `json:"username" yaml:"username" validate:"required,gt=0"`
	Password      *string               `json:"password" yaml:"password" validate:"required,gt=0"`
	ClientID      *string               `json:"clientID" yaml:"clientID" validate:"required,gt=0"`
	ClientSUN     *string               `json:"clientSUN" yaml:"clientSUN" validate:"required,gt=0,numeric"`
	PaymentPlanID *string               `json:"paymentPlanID" yaml:"paymentPlanID" validate:"required,gt=0"`
	Environment   BottomlineEnvironment `json:"environment" yaml:"environment" validate:"oneof=uat sandbox production"`
}

func (c BottomlineCredentials) GetLibrary() Library {
	return LibraryBottomline
}

func (c *BottomlineCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *BottomlineCredentials) Validate() error {
	return nil
}

func (c *BottomlineCredentials) GetSecureFields() []*string {
	return []*string{c.Username, c.Password, c.ClientID, c.ClientSUN}
}

func (c *BottomlineCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *BottomlineCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c BottomlineCredentials) SupportsSca() bool {
	return false
}

func (c BottomlineCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if methodType == chtype.PAYMENT_METHOD_TYPE_DIRECTDEBIT {
		return true
	}
	return false
}
