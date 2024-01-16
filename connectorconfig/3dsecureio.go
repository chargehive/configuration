package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type ThreeDSecureIOEnvironment string

const (
	ThreeDSecureIOEnvironmentSandbox    ThreeDSecureIOEnvironment = "sandbox"
	ThreeDSecureIOEnvironmentProduction ThreeDSecureIOEnvironment = "production"
)

type ThreeDSecureIOCredentials struct {
	ApiKey             *string                   `json:"apiKey" yaml:"apiKey" validate:"required,gt=0"`
	MerchantName       string                    `json:"merchantName" yaml:"merchantName"`
	MCC                string                    `json:"mcc" yaml:"mcc"`
	AcquirerBIN        string                    `json:"acquirerBIN" yaml:"acquirerBIN"`
	AcquirerMerchantID string                    `json:"acquirerMerchantID" yaml:"acquirerMerchantID"`
	Environment        ThreeDSecureIOEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
}

func (c *ThreeDSecureIOCredentials) GetMID() string {
	return c.MerchantName
}

func (c *ThreeDSecureIOCredentials) GetApiKey() string {
	if c.ApiKey == nil {
		return ""
	}
	return *c.ApiKey
}

func (c *ThreeDSecureIOCredentials) GetLibrary() Library {
	return LibraryThreeDSecureIO
}

func (c *ThreeDSecureIOCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeAuthentication}
}

func (c *ThreeDSecureIOCredentials) Validate() error {
	return nil
}

func (c *ThreeDSecureIOCredentials) GetSecureFields() []*string {
	return []*string{c.ApiKey}
}

func (c *ThreeDSecureIOCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *ThreeDSecureIOCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *ThreeDSecureIOCredentials) SupportsSca() bool {
	return true
}

func (c *ThreeDSecureIOCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *ThreeDSecureIOCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *ThreeDSecureIOCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeSandbox && c.Environment == ThreeDSecureIOEnvironmentProduction {
		return false
	}
	return true
}

func (c *ThreeDSecureIOCredentials) IsRecoveryAgent() bool {
	return false
}
