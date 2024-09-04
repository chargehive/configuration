package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type RecaptchaCredentials struct {
	SiteKey   string `json:"siteKey" yaml:"siteKey" validate:"required"`
	ProjectID string `json:"projectId" yaml:"projectId" validate:"required"`
}

func (c *RecaptchaCredentials) GetLibrary() Library {
	return LibraryRecaptcha
}

func (c *RecaptchaCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeFraud}
}

func (c *RecaptchaCredentials) Validate() error {
	return nil
}

func (c *RecaptchaCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *RecaptchaCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *RecaptchaCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *RecaptchaCredentials) SupportsSca() bool {
	return true
}

func (c *RecaptchaCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *RecaptchaCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *RecaptchaCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (c *RecaptchaCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *RecaptchaCredentials) Supports3RI() bool {
	return false
}
