package connectorconfig

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type ApplePayCredential interface {
	GetApplePay() *ApplePay
}

type ApplePay struct {
	// ConnectorID The ID of the connector that provides the ApplePay service
	ConnectorID string `json:"applePayConnectorID,omitempty" yaml:"applePayConnectorID,omitempty" validate:"-"`

	Region string `json:"region,omitempty" yaml:"region,omitempty" validate:"-"`

	// AppleMerchantIdentifier REQUIRED TO ENABLE APPLE PAY Merchant Identifier specified in the Apple Developer Merchant section
	AppleMerchantIdentifier string `json:"appleMerchantIdentifier,omitempty" yaml:"appleMerchantIdentifier,omitempty" validate:"-"`
	// AppleMerchantDisplayName Value to be displayed on the payment page
	AppleMerchantDisplayName string `json:"appleMerchantDisplayName,omitempty" yaml:"appleMerchantDisplayName,omitempty" validate:"required_with=AppleMerchantIdentifier"`

	// AppleSupportedNetworks Specifies which card networks will be accepted by ApplePay
	AppleSupportedNetworks []AppleSupportedNetwork `json:"appleSupportedNetworks,omitempty" yaml:"appleSupportedNetworks,omitempty" validate:"required_with=AppleMerchantIdentifier,dive,oneof=amex cartesBancaires chinaUnionPay discover eftpos electron elo interac jcb mada maestro masterCard privateLabel visa vPay"`
	// AppleMerchantCapabilities The payment capabilities supported by the merchant
	AppleMerchantCapabilities []AppleMerchantCapability `json:"appleMerchantCapabilities,omitempty" yaml:"appleMerchantCapabilities,omitempty" validate:"required_with=AppleMerchantIdentifier,dive,oneof=supports3DS supportsCredit supportsDebit supportsEMV"`

	// AppleMerchantCertificate Merchant certificate in the Apple Developer Merchant section (must be base64 encoded!)
	AppleMerchantCertificate *string `json:"appleMerchantCertificate" yaml:"appleMerchantCertificate" validate:"required_with=AppleMerchantIdentifier"`
	// AppleMerchantPrivateKey Merchant private key generated from the CSR in the Apple Developer Merchant section (must be base64 encoded!)
	AppleMerchantPrivateKey *string `json:"appleMerchantPrivateKey" yaml:"appleMerchantPrivateKey" validate:"required_with=AppleMerchantIdentifier"`
}

func (a *ApplePay) GetLibrary() Library {
	return LibraryApplePay
}

func (a *ApplePay) GetSupportedTypes() []LibraryType {
	return []LibraryType{}
}

func (a *ApplePay) Validate() error {
	if a.AppleMerchantIdentifier != "" {
		// ensure certificates are valid
		certData, _ := base64.StdEncoding.DecodeString(a.GetAppleMerchantCertificate())
		keyData, _ := base64.StdEncoding.DecodeString(a.GetAppleMerchantPrivateKey())
		_, err := tls.X509KeyPair(certData, keyData)
		return err
	}
	return nil
}

func (a *ApplePay) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(a.GetLibrary())}
	con.Configuration, _ = json.Marshal(a)
	return con
}

func (a *ApplePay) FromJson(input []byte) error {
	return json.Unmarshal(input, a)
}

func (a *ApplePay) SupportsSca() bool {
	return false
}

func (a *ApplePay) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	// this connector does not directly support any types, ApplePay tokens are processed through another connector
	return false
}

func (a *ApplePay) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (a *ApplePay) IsRecoveryAgent() bool {
	return false
}

func (a *ApplePay) GetSecureFields() []*string {
	if a == nil {
		return nil
	}
	return []*string{a.AppleMerchantCertificate, a.AppleMerchantPrivateKey}
}

func (a *ApplePay) IsValid() bool {
	return a == nil ||
		a.ConnectorID != "" ||
		(a.GetAppleMerchantIdentifier() != "" &&
			a.GetAppleMerchantDisplayName() != "" &&
			a.GetAppleMerchantCertificate() != "" &&
			a.GetAppleMerchantPrivateKey() != "" &&
			a.Validate() == nil)
}

type (
	AppleMerchantCapability string
	AppleSupportedNetwork   string
)

const (
	AppleMerchantCapabilitysupports3DS    AppleMerchantCapability = "supports3DS"
	AppleMerchantCapabilitysupportsCredit AppleMerchantCapability = "supportsCredit"
	AppleMerchantCapabilitysupportsDebit  AppleMerchantCapability = "supportsDebit"
	AppleMerchantCapabilitysupportsEMV    AppleMerchantCapability = "supportsEMV"

	AppleSupportedNetworkAmex            AppleSupportedNetwork = "amex"
	AppleSupportedNetworkCartesBancaires AppleSupportedNetwork = "cartesBancaires"
	AppleSupportedNetworkChinaUnionPay   AppleSupportedNetwork = "chinaUnionPay"
	AppleSupportedNetworkDiscover        AppleSupportedNetwork = "discover"
	AppleSupportedNetworkEFTPos          AppleSupportedNetwork = "eftpos"
	AppleSupportedNetworkElectron        AppleSupportedNetwork = "electron"
	AppleSupportedNetworkELO             AppleSupportedNetwork = "elo"
	AppleSupportedNetworkInterac         AppleSupportedNetwork = "interac"
	AppleSupportedNetworkJCB             AppleSupportedNetwork = "jcb"
	AppleSupportedNetworkMada            AppleSupportedNetwork = "mada"
	AppleSupportedNetworkMaestro         AppleSupportedNetwork = "maestro"
	AppleSupportedNetworkMasterCard      AppleSupportedNetwork = "masterCard"
	AppleSupportedNetworkPrivateLabel    AppleSupportedNetwork = "privateLabel"
	AppleSupportedNetworkVisa            AppleSupportedNetwork = "visa"
	AppleSupportedNetworkVPay            AppleSupportedNetwork = "vPay"
)

func (a *ApplePay) GetAppleMerchantPublicKey() string {
	if a == nil {
		return ""
	}
	if a.AppleMerchantCertificate == nil {
		return ""
	}
	return *a.AppleMerchantCertificate
}

func (a *ApplePay) GetAppleMerchantPrivateKey() string {
	if a == nil {
		return ""
	}
	if a.AppleMerchantPrivateKey == nil {
		return ""
	}
	return *a.AppleMerchantPrivateKey
}

func (a *ApplePay) GetAppleMerchantIdentifier() string {
	if a == nil {
		return ""
	}
	return a.AppleMerchantIdentifier
}

func (a *ApplePay) GetAppleMerchantDisplayName() string {
	if a == nil {
		return ""
	}
	return a.AppleMerchantDisplayName
}

func (a *ApplePay) GetAppleMerchantCertificate() string {
	if a == nil {
		return ""
	}
	return *a.AppleMerchantCertificate
}
