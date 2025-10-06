package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type WoodforestCredentials struct {
	MerchantID string `json:"merchantId,omitempty" yaml:"merchantId,omitempty"`
}

func (c *WoodforestCredentials) GetMID() string {
	return c.MerchantID
}

func (c *WoodforestCredentials) GetLibrary() Library {
	return LibraryWoodforest
}

func (c *WoodforestCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *WoodforestCredentials) Validate() error {
	return nil
}

func (c *WoodforestCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *WoodforestCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *WoodforestCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *WoodforestCredentials) SupportsSca() bool {
	return false
}

func (c *WoodforestCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *WoodforestCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *WoodforestCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (c *WoodforestCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *WoodforestCredentials) Supports3RI() bool {
	return false
}
