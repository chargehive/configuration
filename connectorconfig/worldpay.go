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
	Username              *string             `json:"username" yaml:"username"`
	Password              *string             `json:"password" yaml:"password"`
	MerchantID            string              `json:"merchantID" yaml:"merchantID"`
	ReportGroup           string              `json:"reportGroup" yaml:"reportGroup"`
	Environment           WorldpayEnvironment `json:"environment" yaml:"environment"`
	CardinalApiIdentifier *string             `json:"cardinalApiIdentifier" yaml:"cardinalApiIdentifier"`
	CardinalApiKey        *string             `json:"cardinalApiKey" yaml:"cardinalApiKey"`
	CardinalOrgUnitId     *string             `json:"cardinalOrgUnitId" yaml:"cardinalOrgUnitId"`
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
