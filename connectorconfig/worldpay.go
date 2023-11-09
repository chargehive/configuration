package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
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

// assert interface compliance
var _ Credentials = (*WorldpayCredentials)(nil)

type WorldpayCredentials struct {
	Username              *string             `json:"username" yaml:"username" validate:"required,gt=0"`
	Password              *string             `json:"password" yaml:"password" validate:"required,gt=0"`
	MerchantID            string              `json:"merchantID" yaml:"merchantID" validate:"gte=1,lte=50"`
	ReportGroup           string              `json:"reportGroup" yaml:"reportGroup" validate:"gte=1,lte=25"`
	Environment           WorldpayEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox postlive transactpostlive production productiontransact prelive transactprelive"`
	CardinalApiIdentifier *string             `json:"cardinalApiIdentifier" yaml:"cardinalApiIdentifier" validate:"required"`
	CardinalApiKey        *string             `json:"cardinalApiKey" yaml:"cardinalApiKey" validate:"required"`
	CardinalOrgUnitId     *string             `json:"cardinalOrgUnitId" yaml:"cardinalOrgUnitId" validate:"required"`
	GooglePayPageId       string              `json:"googlePayPageId"` // vantiv:merchantPayPageId
	GooglePay             *GooglePay          `json:"googlePay,omitempty" yaml:"googlePay,omitempty"`
	ApplePay              *ApplePayEmbedded   `json:"applePay,omitempty" yaml:"applePay,omitempty"`
}

func (c *WorldpayCredentials) GetGooglePayParams() map[string]string {
	return map[string]string{
		"gateway":                      "worldpay",
		"gatewayMerchantId":            c.GetGooglePay().GetGoogleCardMerchantId(),
		"vantiv:merchantPayPageId":     c.GooglePayPageId,
		"vantiv:merchantOrderId":       "YOUR_ORDER_ID",
		"vantiv:merchantTransactionId": "YOUR_TRANSACTION_ID",
		"vantiv:merchantReportGroup":   c.ReportGroup,
	}
}

func (c *WorldpayCredentials) GetMID() string {
	return c.MerchantID
}

func (c *WorldpayCredentials) GetGooglePay() *GooglePay {
	return c.GooglePay
}

func (c *WorldpayCredentials) GetApplePay() *ApplePayEmbedded {
	return c.ApplePay
}

func (c *WorldpayCredentials) GetCardinalApiIdentifier() string {
	if c.CardinalApiIdentifier == nil {
		return ""
	}
	return *c.CardinalApiIdentifier
}

func (c *WorldpayCredentials) GetCardinalApiKey() string {
	if c.CardinalApiKey == nil {
		return ""
	}
	return *c.CardinalApiKey
}

func (c *WorldpayCredentials) GetCardinalOrgUnitId() string {
	if c.CardinalOrgUnitId == nil {
		return ""
	}
	return *c.CardinalOrgUnitId
}

func (c *WorldpayCredentials) GetLibrary() Library {
	return LibraryWorldpay
}

func (c *WorldpayCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *WorldpayCredentials) Validate() error {
	return nil
}

func (c *WorldpayCredentials) GetSecureFields() []*string {
	fields := []*string{c.Username, c.Password, c.CardinalApiIdentifier, c.CardinalApiKey}
	if c.ApplePay != nil {
		fields = append(fields, c.ApplePay.AppleMerchantPrivateKey, c.ApplePay.AppleMerchantCertificate)
	}

	return fields
}

func (c *WorldpayCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *WorldpayCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *WorldpayCredentials) SupportsSca() bool {
	return c.GetCardinalApiIdentifier() != "" && c.GetCardinalApiKey() != "" && c.GetCardinalOrgUnitId() != ""
}

func (c *WorldpayCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	if methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY {
		return c.GetApplePay().IsValid()
	}
	if methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY {
		return c.GetGooglePay().IsValid()
	}
	return true
}

func (c *WorldpayCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeSandbox {
		if c.Environment == WorldpayEnvironmentProduction || c.Environment == WorldpayEnvironmentProductionTransact {
			return false
		}
	}
	return true
}

func (c *WorldpayCredentials) IsRecoveryAgent() bool {
	return false
}
