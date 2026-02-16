package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type PagosNetworkTokenizationCredentials struct {
	MerchantID string `json:"merchantId,omitempty" yaml:"merchantId,omitempty"`
}

func (c *PagosNetworkTokenizationCredentials) GetMID() string {
	return c.MerchantID
}

func (c *PagosNetworkTokenizationCredentials) GetLibrary() Library {
	return LibraryPagosNetworkTokenization
}

func (c *PagosNetworkTokenizationCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeNetworkTokenization}
}

func (c *PagosNetworkTokenizationCredentials) Validate() error {
	return nil
}

func (c *PagosNetworkTokenizationCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *PagosNetworkTokenizationCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *PagosNetworkTokenizationCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *PagosNetworkTokenizationCredentials) SupportsSca() bool {
	return false
}

func (c *PagosNetworkTokenizationCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *PagosNetworkTokenizationCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *PagosNetworkTokenizationCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (c *PagosNetworkTokenizationCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *PagosNetworkTokenizationCredentials) Supports3RI() bool {
	return false
}

func (c *PagosNetworkTokenizationCredentials) IsAccountUpdater() bool {
	return false
}
