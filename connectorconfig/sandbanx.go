package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type SandbanxCredentials struct {
	MerchantID             string `json:"merchantId" yaml:"merchantID" validate:"-"`
	ProcessingDelaySeconds int    `json:"processingDelaySeconds" yaml:"processingDelaySeconds" validate:"-"` // seconds to delay processing
	Offline                bool   `json:"offline" yaml:"offline" validate:"-"`                               // Return service unavailable
	FailAuth               bool   `json:"failAuth" yaml:"failAuth" validate:"-"`                             // Fail to authorize
	ChaosLevel             int    `json:"chaosLevel" yaml:"chaosLevel" validate:"-"`                         // Percent of errors
}

func (c *SandbanxCredentials) GetMID() string {
	return c.MerchantID
}

func (c *SandbanxCredentials) GetLibrary() Library {
	return LibrarySandbanx
}

func (c *SandbanxCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *SandbanxCredentials) Validate() error {
	return nil
}

func (c *SandbanxCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *SandbanxCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *SandbanxCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *SandbanxCredentials) SupportsSca() bool {
	return true
}

func (c *SandbanxCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *SandbanxCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *SandbanxCredentials) CanPlanModeUse(environment.Mode) bool {
	return true
}

func (c *SandbanxCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *SandbanxCredentials) Supports3RI() bool {
	return false
}
