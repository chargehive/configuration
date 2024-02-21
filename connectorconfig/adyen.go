package connectorconfig

import (
	"encoding/json"
	"strings"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type AdyenEnvironment string

const (
	AdyenEnvironmentSandbox    AdyenEnvironment = "sandbox"
	AdyenEnvironmentProduction AdyenEnvironment = "production"
)

type AdyenCredentials struct {
	Environment AdyenEnvironment `json:"environment" yaml:"environment" validate:"required,oneof=sandbox production"`
	MID         string           `json:"mid" yaml:"mid" validate:"required"`
}

func (c *AdyenCredentials) GetMID() string {
	return c.MID
}

func (c *AdyenCredentials) GetLibrary() Library {
	return LibraryAdyen
}

func (c *AdyenCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *AdyenCredentials) GetSecureFields() []*string {
	return nil
}

func (c *AdyenCredentials) Validate() error {
	return nil
}

func (c *AdyenCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *AdyenCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *AdyenCredentials) SupportsSca() bool {
	return false
}

func (c *AdyenCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

var adyenAllowedCountires = []string{}

func (c *AdyenCredentials) SupportsCountry(country string) bool {
	for _, v := range adyenAllowedCountires {
		if strings.EqualFold(v, country) {
			return true
		}
	}
	return false
}

func (c *AdyenCredentials) CanPlanModeUse(mode environment.Mode) bool {

	if mode == environment.ModeSandbox && c.Environment == AdyenEnvironmentSandbox {
		return true
	}

	if mode == environment.ModeProduction && c.Environment == AdyenEnvironmentProduction {
		return true
	}

	return false
}

func (c *AdyenCredentials) IsRecoveryAgent() bool {
	return false
}
