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
	Username                 *string             `json:"username" yaml:"username" validate:"required,gt=0"`
	Password                 *string             `json:"password" yaml:"password" validate:"required,gt=0"`
	MerchantID               string              `json:"merchantID" yaml:"merchantID" validate:"gte=1,lte=50"`
	ReportGroup              string              `json:"reportGroup" yaml:"reportGroup" validate:"gte=1,lte=25"`
	Environment              WorldpayEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox postlive transactpostlive production productiontransact prelive transactprelive"`
	CardinalApiIdentifier    *string             `json:"cardinalApiIdentifier" yaml:"cardinalApiIdentifier" validate:"required"`
	CardinalApiKey           *string             `json:"cardinalApiKey" yaml:"cardinalApiKey" validate:"required"`
	CardinalOrgUnitId        *string             `json:"cardinalOrgUnitId" yaml:"cardinalOrgUnitId" validate:"required"`
	AppleMerchantIdentifier  string              `json:"appleMerchantIdentifier" yaml:"appleMerchantIdentifier"`
	AppleMerchantDisplayName string              `json:"appleMerchantDisplayName" yaml:"appleMerchantDisplayName" validate:"required_with=AppleMerchantIdentifier"`
	AppleInitiative          string              `json:"appleInitiative" yaml:"appleInitiative" validate:"required_with=AppleMerchantIdentifier,omitempty,oneof=web messaging"`
	AppleInitiativeContext   string              `json:"appleInitiativeContext" yaml:"appleInitiativeContext" validate:"required_with=AppleMerchantIdentifier"`
	AppleMerchantCertificate *string             `json:"appleMerchantPublicKey" yaml:"appleMerchantPublicKey" validate:"required"`
	AppleMerchantPrivateKey  *string             `json:"appleMerchantPrivateKey" yaml:"appleMerchantPrivateKey" validate:"required"`
}

func (c WorldpayCredentials) GetAppleMerchantPublicKey() string {
	if c.AppleMerchantCertificate == nil {
		return ""
	}
	return *c.AppleMerchantCertificate
}

func (c WorldpayCredentials) GetAppleMerchantPrivateKey() string {
	if c.AppleMerchantPrivateKey == nil {
		return ""
	}
	return *c.AppleMerchantPrivateKey
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
	return []*string{c.Username, c.Password, c.CardinalApiIdentifier, c.CardinalApiKey, c.AppleMerchantPrivateKey, c.AppleMerchantCertificate}
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
	if c.AppleMerchantIdentifier != "" &&
		c.AppleMerchantDisplayName != "" &&
		c.AppleInitiative != "" &&
		c.AppleInitiativeContext != "" &&
		(c.AppleMerchantCertificate != nil && c.AppleMerchantCertificate != new(string)) &&
		(c.AppleMerchantPrivateKey != nil && c.AppleMerchantPrivateKey != new(string)) {
		return true
	}
	return false
}
