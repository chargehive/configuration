package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

// MaxMindMinFraudServiceType is the maxmind minFraud service type
type MaxMindMinFraudServiceType int

// MaxMindMinFraudServiceTypes minFraud Score, minFraud Insights, and minFraud Factors
const (
	MaxMindMinFraudServiceTypeScore MaxMindMinFraudServiceType = iota
	MaxMindMinFraudServiceTypeInsights
	MaxMindMinFraudServiceTypeFactors
)

type MaxMindCredentials struct {
	AccountID   string                     `json:"accountID" yaml:"accountID" validate:"required"`
	LicenceKey  *string                    `json:"licenceKey" yaml:"licenceKey" validate:"required"`
	ServiceType MaxMindMinFraudServiceType `json:"serviceType" yaml:"serviceType" validate:"min=0,max=2"`
}

func (c MaxMindCredentials) GetLibrary() Library {
	return LibraryMaxMind
}

func (c *MaxMindCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeFraud}
}

func (c *MaxMindCredentials) Validate() error {
	return nil
}

func (c *MaxMindCredentials) GetSecureFields() []*string {
	return []*string{c.LicenceKey}
}

func (c *MaxMindCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *MaxMindCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
