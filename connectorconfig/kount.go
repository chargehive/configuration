package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type KountEnvironment string

const (
	KountEnvironmentTest       KountEnvironment = "test"
	KountEnvironmentProduction KountEnvironment = "production"
)

type KountCredentials struct {
	SiteID                string           `json:"siteID" yaml:"siteID" validate:"required"`
	MerchantID            string           `json:"merchantID" yaml:"merchantID" validate:"required"`
	ConfigKey             string           `json:"configKey" yaml:"configKey" validate:"required"`
	APIKey                string           `json:"apiKey" yaml:"apiKey" validate:"required"`
	DataCollectorURL      string           `json:"dataCollectorURL" yaml:"dataCollectorURL" validate:"required"`
	RiskInquiryServiceURL string           `json:"riskInquiryServiceURL" yaml:"riskInquiryServiceURL" validate:"required"`
	Environment           KountEnvironment `json:"environment" yaml:"environment" validate:"oneof=test production"`
}

func (c *KountCredentials) GetMID() string {
	return c.MerchantID
}

func (c *KountCredentials) GetLibrary() Library {
	return LibraryKount
}

func (c *KountCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeFraud}
}

func (c *KountCredentials) Validate() error {
	return nil
}

func (c *KountCredentials) GetSecureFields() []*string {
	return nil
}

func (c *KountCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *KountCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
func (c *KountCredentials) SupportsSca() bool {
	return false
}

func (c *KountCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *KountCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *KountCredentials) CanPlanModeUse(environment.Mode) bool {
	return true
}

func (c *KountCredentials) IsRecoveryAgent() bool {
	return false
}
