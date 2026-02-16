package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/grpc/cht"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type PagosRealtimeAccountUpdaterCredentials struct {
	MerchantId string `json:"merchantId" yaml:"merchantId"`
}

func (c *PagosRealtimeAccountUpdaterCredentials) GetMID() string {
	return c.MerchantId
}

func (c *PagosRealtimeAccountUpdaterCredentials) GetLibrary() Library {
	return LibraryPagosRealtimeAccountUpdater
}

func (c *PagosRealtimeAccountUpdaterCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeMethodUpdater}
}

func (c *PagosRealtimeAccountUpdaterCredentials) Validate() error {
	return nil
}

func (c *PagosRealtimeAccountUpdaterCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *PagosRealtimeAccountUpdaterCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *PagosRealtimeAccountUpdaterCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *PagosRealtimeAccountUpdaterCredentials) SupportsSca() bool {
	return false
}

func (c *PagosRealtimeAccountUpdaterCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *PagosRealtimeAccountUpdaterCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *PagosRealtimeAccountUpdaterCredentials) CanPlanModeUse(environment.Mode) bool {
	return true
}

func (c *PagosRealtimeAccountUpdaterCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *PagosRealtimeAccountUpdaterCredentials) Supports3RI() bool {
	return false
}

func (c *PagosRealtimeAccountUpdaterCredentials) IsAccountUpdater() bool {
	return true
}

func (c *PagosRealtimeAccountUpdaterCredentials) SupportedTokenSources() []cht.TokenSource {
	return []cht.TokenSource{cht.TS_PAN}
}
