package connectorconfig

import (
	"encoding/json"
	"strings"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type TrustPaymentsRegion string

const (
	TrustPaymentsRegionUS TrustPaymentsRegion = "us"
	TrustPaymentsRegionEU TrustPaymentsRegion = "eu"
)

type TrustPaymentsCredentials struct {
	Username string              `json:"username" yaml:"username" validate:"required"`
	Password string              `json:"password" yaml:"password" validate:"required"`
	SiteRef  string              `json:"siteRef" yaml:"siteRef" validate:"required"`
	Region   TrustPaymentsRegion `json:"region" yaml:"region" validate:"oneof=us eu"`
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

	if mode == environment.ModeProduction && strings.HasPrefix(c.SiteRef, "test_") {
		return false
	}

	if mode == environment.ModeSandbox && !strings.HasPrefix(c.SiteRef, "test_") {
		return false
	}

	return true
}

func (c *TrustPaymentsCredentials) IsRecoveryAgent() bool {
	return false
}
