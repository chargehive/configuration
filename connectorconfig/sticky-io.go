package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/grpc/cht"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type StickyIOCredentials struct {
	ApiKey string `json:"apiKey" yaml:"apiKey" validate:"required,gt=0"`
}

func (c *StickyIOCredentials) GetApiKey() string {
	return c.ApiKey
}

func (c *StickyIOCredentials) GetLibrary() Library {
	return LibraryStickyIO
}

func (c *StickyIOCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeScheduler}
}

func (c *StickyIOCredentials) Validate() error {
	return nil
}

func (c *StickyIOCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *StickyIOCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *StickyIOCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *StickyIOCredentials) SupportsSca() bool {
	return true
}

func (c *StickyIOCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *StickyIOCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *StickyIOCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (c *StickyIOCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *StickyIOCredentials) Supports3RI() bool {
	return false
}

func (c *StickyIOCredentials) IsAccountUpdater() bool {
	return false
}

func (c *StickyIOCredentials) SupportedTokenSources() []cht.TokenSource {
	return []cht.TokenSource{cht.TS_PAN}
}
