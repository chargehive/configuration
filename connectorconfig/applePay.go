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
}

type ApplePayCredentials struct {
	// Region whether global(empty) or china
	Region string `json:"region,omitempty" yaml:"region,omitempty" validate:"oneof='' cn"`

	// AppleMerchantIdentifier REQUIRED TO ENABLE APPLE PAY Merchant Identifier specified in the Apple Developer Merchant section
	AppleMerchantIdentifier string `json:"appleMerchantIdentifier,omitempty" yaml:"appleMerchantIdentifier,omitempty" validate:"required"`
	// AppleMerchantDisplayName Value to be displayed on the payment page
	AppleMerchantDisplayName string `json:"appleMerchantDisplayName,omitempty" yaml:"appleMerchantDisplayName,omitempty" validate:"required"`
	// AppleSupportedNetworks Specifies which card networks will be accepted by ApplePayCredentials
	AppleSupportedNetworks []AppleSupportedNetwork `json:"appleSupportedNetworks,omitempty" yaml:"appleSupportedNetworks,omitempty" validate:"required,dive,oneof=amex cartesBancaires chinaUnionPay discover eftpos electron elo interac jcb mada maestro masterCard privateLabel visa vPay"`
	// AppleMerchantCapabilities The payment capabilities supported by the merchant
	AppleMerchantCapabilities []AppleMerchantCapability `json:"appleMerchantCapabilities,omitempty" yaml:"appleMerchantCapabilities,omitempty" validate:"required,dive,oneof=supports3DS supportsCredit supportsDebit supportsEMV"`

	// AppleIdentityCertificate Merchant certificate in the Apple Developer Merchant section (must be base64 encoded!)
	AppleIdentityCertificate *string `json:"appleIdentityCertificate" yaml:"appleIdentityCertificate" validate:"required"`
	// AppleIdentityPrivateKey Merchant private key generated from the CSR in the Apple Developer Merchant section (must be base64 encoded!)
	AppleIdentityPrivateKey *string `json:"appleIdentityPrivateKey" yaml:"AppleIdentityPrivateKey" validate:"required"`

	// AppleExistingMethodRequired Chargehive will not use this connector if the customer does not have a apple payment method already saved
	AppleExistingMethodRequired bool `json:"appleExistingMethodRequired,omitempty" yaml:"appleExistingMethodRequired,omitempty" validate:"-"`
	// AppleExistingMethodReport Chargehive will request the existing payment method information from ApplePay
	AppleExistingMethodReport bool `json:"appleExistingMethodReport,omitempty" yaml:"appleExistingMethodReport,omitempty" validate:"-"`

	// AppleCardAllowDebit Allow customer to pay with debit card
	AppleCardAllowDebit bool `json:"appleCardAllowDebit,omitempty" yaml:"appleCardAllowDebit,omitempty" validate:"-"`
	// AppleCardAllowCredit Allow customer to pay with credit card
	AppleCardAllowCredit bool `json:"appleCardAllowCredit,omitempty" yaml:"appleCardAllowCredit,omitempty" validate:"-"`

	// AppleEmailRequired (emailRequired) Set to true to request an email address.
	AppleEmailRequired bool `json:"appleEmailRequired,omitempty" yaml:"appleEmailRequired,omitempty" validate:"-"`

	// AppleCardBillingAddressReq Set to true if you require a billing address. A billing address should only be requested if it's required to process the transaction
	AppleCardBillingAddressReq bool `json:"appleCardBillingAddressReq,omitempty" yaml:"appleCardBillingAddressReq,omitempty" validate:"-"`
	// AppleCardBillingPhoneReq  Set to true if a phone number is required to process the transaction.
	AppleCardBillingPhoneReq bool `json:"appleCardBillingPhoneReq,omitempty" yaml:"appleCardBillingPhoneReq,omitempty" validate:"-"`

	// AppleCardShippingAddressReq Set to true if you require a shipping address. A shipping address should only be requested if it's required to process the transaction
	AppleCardShippingAddressReq bool `json:"appleCardShippingAddressReq,omitempty" yaml:"appleCardShippingAddressReq,omitempty" validate:"-"`
	// AppleCardShippingPhoneReq  Set to true if a phone number is required to process the transaction.
	AppleCardShippingPhoneReq bool `json:"appleCardShippingPhoneReq,omitempty" yaml:"appleCardShippingPhoneReq,omitempty" validate:"-"`
}

func (a *ApplePayCredentials) GetMID() string {
	return a.AppleMerchantIdentifier
}

func (a *ApplePayCredentials) GetLibrary() Library {
	return LibraryApplePay
}

func (a *ApplePayCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{}
}

func (a *ApplePayCredentials) Validate() error {
	if a.AppleMerchantIdentifier != "" {
		// ensure certificates are valid
		certData, _ := base64.StdEncoding.DecodeString(a.GetAppleIdentityCertificate())
		keyData, _ := base64.StdEncoding.DecodeString(a.GetAppleIdentityPrivateKey())
		_, err := tls.X509KeyPair(certData, keyData)
		return err
	}
	return nil
}

func (a *ApplePayCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(a.GetLibrary())}
	con.Configuration, _ = json.Marshal(a)
	return con
}

func (a *ApplePayCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, a)
}

func (a *ApplePayCredentials) SupportsSca() bool {
	return false
}

func (a *ApplePayCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	// this connector does not directly support any types, ApplePay tokens are processed through another connector
	return false
}

func (a *ApplePayCredentials) SupportsCountry(country string) bool {
	return true
}

func (a *ApplePayCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (a *ApplePayCredentials) IsRecoveryAgent() bool {
	return false
}

func (a *ApplePayCredentials) GetSecureFields() []*string {
	if a == nil {
		return nil
	}
	return []*string{a.AppleIdentityPrivateKey, a.AppleIdentityCertificate}
}

func (a *ApplePayCredentials) IsValid() bool {
	if a == nil {
		return false
	}
	return a.GetAppleMerchantIdentifier() != "" &&
		a.GetAppleMerchantDisplayName() != "" &&
		a.GetAppleIdentityCertificate() != "" &&
		a.GetAppleIdentityPrivateKey() != "" &&
		a.Validate() == nil
}

type (
	AppleMerchantCapability string
	AppleSupportedNetwork   string
)

//goland:noinspection GoUnusedConst
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

func (a *ApplePayCredentials) GetAppleIdentityCertificate() string {
	if a == nil {
		return ""
	}
	if a.AppleIdentityCertificate == nil {
		return ""
	}
	return *a.AppleIdentityCertificate
}

func (a *ApplePayCredentials) GetAppleIdentityPrivateKey() string {
	if a == nil {
		return ""
	}
	if a.AppleIdentityPrivateKey == nil {
		return ""
	}
	return *a.AppleIdentityPrivateKey
}

func (a *ApplePayCredentials) GetAppleMerchantIdentifier() string {
	if a == nil {
		return ""
	}
	return a.AppleMerchantIdentifier
}

func (a *ApplePayCredentials) GetAppleMerchantDisplayName() string {
	if a == nil {
		return ""
	}
	return a.AppleMerchantDisplayName
}

func (a *ApplePayCredentials) GetAppleExistingMethodRequired() bool {
	if a == nil {
		return false
	}
	return a.AppleExistingMethodRequired
}
func (a *ApplePayCredentials) GetAppleExistingMethodReport() bool {
	if a == nil {
		return false
	}
	return a.AppleExistingMethodReport
}
func (a *ApplePayCredentials) GetAppleCardAllowDebit() bool {
	if a == nil {
		return false
	}
	return a.AppleCardAllowDebit
}
func (a *ApplePayCredentials) GetAppleCardAllowCredit() bool {
	if a == nil {
		return false
	}
	return a.AppleCardAllowCredit
}
func (a *ApplePayCredentials) GetAppleEmailRequired() bool {
	if a == nil {
		return false
	}
	return a.AppleEmailRequired
}
func (a *ApplePayCredentials) GetAppleCardBillingAddressReq() bool {
	if a == nil {
		return false
	}
	return a.AppleCardBillingAddressReq
}
func (a *ApplePayCredentials) GetAppleCardBillingPhoneReq() bool {
	if a == nil {
		return false
	}
	return a.AppleCardBillingPhoneReq
}
func (a *ApplePayCredentials) GetAppleCardShippingAddressReq() bool {
	if a == nil {
		return false
	}
	return a.AppleCardShippingAddressReq
}
func (a *ApplePayCredentials) GetAppleCardShippingPhoneReq() bool {
	if a == nil {
		return false
	}
	return a.AppleCardShippingPhoneReq
}
