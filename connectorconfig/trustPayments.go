package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type TrustPaymentsCredentials struct {
}

func (c *TrustPaymentsCredentials) GetLibrary() Library {
	return LibraryTrustPayments
}

func (c *TrustPaymentsCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *TrustPaymentsCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *TrustPaymentsCredentials) Validate() error {
	return nil
}

func (c *TrustPaymentsCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *TrustPaymentsCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *TrustPaymentsCredentials) SupportsSca() bool {
	return false
}

func (c *TrustPaymentsCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}

	return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
}

func (c *TrustPaymentsCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (c *TrustPaymentsCredentials) IsRecoveryAgent() bool {
	return false
}
