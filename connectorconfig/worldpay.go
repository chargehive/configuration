package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type WorldpayEnvironment string

const (
	WorldpayEnvironmentSandbox            WorldpayEnvironment = "sandbox"
	WorldpayEnvironmentPostLive           WorldpayEnvironment = "postlive"
	WorldpayEnvironmentTransactPostLive   WorldpayEnvironment = "transactpostlive"
	WorldpayEnvironmentProduction         WorldpayEnvironment = "production"
	WorldpayEnvironmentProductionTransact WorldpayEnvironment = "productiontransact"
	WorldpayEnvironmentPrelive            WorldpayEnvironment = "prelive"
	WorldpayEnvironmentTransactPreLive    WorldpayEnvironment = "transactprelive"
)

type WorldpayCredentials struct {
	Username              *string             `json:"username" yaml:"username" validate:"required,gt=0"`
	Password              *string             `json:"password" yaml:"password" validate:"required,gt=0"`
	MerchantID            string              `json:"merchantID" yaml:"merchantID" validate:"gte=1,lte=50"`
	ReportGroup           string              `json:"reportGroup" yaml:"reportGroup" validate:"gte=1,lte=25"`
	Environment           WorldpayEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox postlive transactpostlive production productiontransact prelive transactprelive"`
	CardinalApiIdentifier *string             `json:"cardinalApiIdentifier" yaml:"cardinalApiIdentifier" validate:"required"`
	CardinalApiKey        *string             `json:"cardinalApiKey" yaml:"cardinalApiKey" validate:"required"`
	CardinalOrgUnitId     *string             `json:"cardinalOrgUnitId" yaml:"cardinalOrgUnitId" validate:"required"`
}

func (c WorldpayCredentials) GetCardinalApiIdentifier() string {
	if c.CardinalApiIdentifier == nil {
		return ""
	}
	return *c.CardinalApiIdentifier
}

func (c WorldpayCredentials) GetCardinalApiKey() string {
	if c.CardinalApiKey == nil {
		return ""
	}
	return *c.CardinalApiKey
}

func (c WorldpayCredentials) GetCardinalOrgUnitId() string {
	if c.CardinalOrgUnitId == nil {
		return ""
	}
	return *c.CardinalOrgUnitId
}

func (c WorldpayCredentials) GetLibrary() Library {
	return LibraryWorldpay
}

func (c *WorldpayCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *WorldpayCredentials) Validate() error {
	return nil
}

func (c *WorldpayCredentials) GetSecureFields() []*string {
	return []*string{c.Username, c.Password, c.CardinalApiIdentifier, c.CardinalApiKey}
}

func (c *WorldpayCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *WorldpayCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c WorldpayCredentials) SupportsSca() bool {
	return c.GetCardinalApiIdentifier() != "" && c.GetCardinalApiKey() != "" && c.GetCardinalOrgUnitId() != ""
}

func (c WorldpayCredentials) SupportsApplePay() bool {
	return true
}
