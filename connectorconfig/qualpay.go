package connectorconfig

import (
	"encoding/json"
	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type QualPayEnvironment string

const (
	QualPayEnvironmentTest QualPayEnvironment = "test"
	QualPayEnvironmentLive QualPayEnvironment = "live"
)

type QualpayCredentials struct {
	APIKey      *string            `json:"apiKey" yaml:"apiKey" validate:"required,gt=0"`
	MerchantID  int64              `json:"merchantID" yaml:"merchantID" validate:"min=1"`
	Environment QualPayEnvironment `json:"environment" yaml:"environment" validate:"oneof=test live"`
}

func (c QualpayCredentials) GetLibrary() Library {
	return LibraryQualPay
}

func (c QualpayCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *QualpayCredentials) Validate() error {
	return nil
}

func (c *QualpayCredentials) GetSecureFields() []*string {
	return []*string{c.APIKey}
}

func (c *QualpayCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *QualpayCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c QualpayCredentials) SupportsSca() bool {
	return false
}

func (c QualpayCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if methodType == chtype.PAYMENT_METHOD_TYPE_CARD {
		return true
	}
	return false
}

func (c QualpayCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeSandbox && c.Environment == QualPayEnvironmentLive {
		return false
	}
	return true
}

func (c QualpayCredentials) IsRecoveryAgent() bool {
	return false
}
