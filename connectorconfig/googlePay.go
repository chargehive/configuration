package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type GooglePayCredential interface {
	GetGooglePay() *GooglePayEmbedded
	GetGooglePayParams() map[string]string
}

type GooglePayCredentials struct {
	GooglePayEmbedded
}

func (g GooglePayCredentials) GetGooglePay() *GooglePayEmbedded {
	return &g.GooglePayEmbedded
}

func (g GooglePayCredentials) GetGooglePayParams() map[string]string {
	return map[string]string{
		"gateway":           g.GetGoogleCardGateway(),
		"gatewayMerchantId": g.GetGoogleCardMerchantId(),
	}
}

func (g *GooglePayCredentials) GetMID() string {
	return g.GoogleMerchantId
}

func (g *GooglePayCredentials) GetLibrary() Library {
	return LibraryApplePay
}

func (g *GooglePayCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{}
}

func (g *GooglePayCredentials) Validate() error {
	return nil
}

func (g *GooglePayCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(g.GetLibrary())}
	con.Configuration, _ = json.Marshal(g)
	return con
}

func (g *GooglePayCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, g)
}

func (g *GooglePayCredentials) SupportsSca() bool {
	return false
}

func (g *GooglePayCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	// this connector does not directly support any types, GooglePayEmbedded tokens are processed through another connector
	return false
}

func (g *GooglePayCredentials) SupportsCountry(country string) bool {
	return true
}

func (g *GooglePayCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (g *GooglePayCredentials) IsRecoveryAgent() bool {
	return false
}

func (g *GooglePayCredentials) Supports3RI() bool {
	return false
}

func (g *GooglePayCredentials) GetSecureFields() []*string {
	if g == nil {
		return nil
	}
	return []*string{}
}

func (g *GooglePayCredentials) IsValid() bool {
	if g == nil {
		return false
	}
	return g.Validate() == nil
}

type GooglePayOptions interface {
	GetGoogleCardGateway() string
	GetGoogleCardMerchantId() string
	GetGoogleAcquirerCountry() string
	GetGoogleExistingMethodRequired() bool
	GetGoogleExistingMethodReport() bool
	GetGoogleEmailReq() bool
	GetGoogleAcceptCard() bool
	GetGoogleCardAllowPrepaid() bool
	GetGoogleCardAllowCredit() bool
	GetGoogleCardBillingAddressReq() bool
	GetGoogleCardBillingPhoneReq() bool
	GetGoogleCardShippingAddressReq() bool
	GetGoogleCardShippingPhoneReq() bool
}
