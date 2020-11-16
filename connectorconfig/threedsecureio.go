package connectorconfig

import (
	"encoding/json"
	"github.com/LucidCube/chargehive-transport-config/plans"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type (
	ThreeDSecureEnvironment string
	ThreeDSecureUrl         string
	ThreeDVersion           string
)

const (
	ThreeDSecureEnvironmentLive    ThreeDSecureEnvironment = "live"
	ThreeDSecureEnvironmentSandbox ThreeDSecureEnvironment = "sandbox"

	ThreeDSecureUrlLive    ThreeDSecureUrl = "https://service.3dsecure.io"
	ThreeDSecureUrlSandbox ThreeDSecureUrl = "https://service.sandbox.3dsecure.io"

	ThreeDVersion200 ThreeDVersion = "2.0.0"
	ThreeDVersion210 ThreeDVersion = "2.1.0"
	ThreeDVersion220 ThreeDVersion = "2.2.0"
	ThreeDVersion230 ThreeDVersion = "2.3.0"
)

type ThreeDSecureIoCredentials struct {
	APIKey             *string                 `json:"apiKey" yaml:"apiKey" validate:"required,gt=0"`                                         // Api key supplied by 3dsecure.io
	SupportsMinVersion ThreeDVersion           `json:"supportsMinVersion" yaml:"supportsMinVersion" validate:"oneof=2.0.0 2.1.0 2.2.0 2.3.0"` // lowest supported version
	SupportsMaxVersion ThreeDVersion           `json:"supportsMaxVersion" yaml:"supportsMaxVersion" validate:"oneof=2.0.0 2.1.0 2.2.0 2.3.0"` // highest supported version
	Environment        ThreeDSecureEnvironment `json:"environment" yaml:"environment" validate:"oneof=live sandbox"`                          // live or sandbox environment
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

func (c ThreeDSecureIoCredentials) CanPlanModeUse(plans.Mode) bool {
	return true
}
