package connectorconfig

import (
	"encoding/json"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type (
	ThreeDSecureEnvironment string
	ThreeDSecureUrl         string
)

const (
	ThreeDSecureEnvironmentLive    ThreeDSecureEnvironment = "live"
	ThreeDSecureEnvironmentSandbox ThreeDSecureEnvironment = "sandbox"
	ThreeDSecureUrlLive            ThreeDSecureUrl         = "https://service.3dsecure.io"
	ThreeDSecureUrlSandbox         ThreeDSecureUrl         = "https://service.sandbox.3dsecure.io"
)

type ThreeDSecureIoCredentials struct {
	APIKey      *string                 `json:"apiKey" yaml:"apiKey" validate:"required,gt=0"`
	Supports210 bool                    `json:"supports210"` // supports 3ds version 2.1.0
	Supports220 bool                    `json:"supports220"` // supports 3ds version 2.2.0
	Environment ThreeDSecureEnvironment `json:"environment"`
}

func (c ThreeDSecureIoCredentials) GetLibrary() Library {
	return LibraryThreeDSecureIo
}

func (c *ThreeDSecureIoCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeSCA}
}

func (c *ThreeDSecureIoCredentials) Validate() error {
	return nil
}

func (c *ThreeDSecureIoCredentials) GetSecureFields() []*string {
	return []*string{c.APIKey}
}

func (c *ThreeDSecureIoCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *ThreeDSecureIoCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c ThreeDSecureIoCredentials) SupportsSca() bool {
	return true
}

func (c ThreeDSecureIoCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if methodType == chtype.PAYMENT_METHOD_TYPE_CARD {
		return true
	}
	return false
}
