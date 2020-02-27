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
	Username              *string             `json:"username" yaml:"username" validate:"required"`
	Password              *string             `json:"password" yaml:"password" validate:"required"`
	MerchantID            string              `json:"merchantID" yaml:"merchantID" validate:"required"`
	ReportGroup           string              `json:"reportGroup" yaml:"reportGroup" validate:"-"`
	Environment           WorldpayEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox postlive transactpostlive production productiontransact prelive transactprelive"`
	CardinalApiIdentifier *string             `json:"cardinalApiIdentifier" yaml:"cardinalApiIdentifier" validate:"required_with=CardinalApiKey CardinalOrgUnitId"`
	CardinalApiKey        *string             `json:"cardinalApiKey" yaml:"cardinalApiKey" validate:"required_with=CardinalApiIdentifier CardinalOrgUnitId"`
	CardinalOrgUnitId     *string             `json:"cardinalOrgUnitId" yaml:"cardinalOrgUnitId" validate:"required_with=CardinalApiIdentifier CardinalApiKey"`
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
	return []*string{c.Username, c.Password}
}

func (c *WorldpayCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *WorldpayCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
