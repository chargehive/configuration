package connectorconfig

type ApplePay struct {
	// AppleMerchantIdentifier Merchant Identifier specified in the Apple Developer Merchant section
	AppleMerchantIdentifier string `json:"appleMerchantIdentifier,omitempty" yaml:"appleMerchantIdentifier,omitempty"`
	// AppleMerchantDisplayName Value to be displayed on the payment page
	AppleMerchantDisplayName string `json:"appleMerchantDisplayName,omitempty" yaml:"appleMerchantDisplayName,omitempty" validate:"required_with=AppleMerchantIdentifier"`
	// AppleMerchantCertificate Merchant certificate in the Apple Developer Merchant section (must be base64 encoded!)
	AppleMerchantCertificate *string `json:"appleMerchantCertificate" yaml:"appleMerchantCertificate" validate:"required"`
	// AppleMerchantPrivateKey Merchant private key generated from the CSR in the Apple Developer Merchant section (must be base64 encoded!)
	AppleMerchantPrivateKey *string `json:"appleMerchantPrivateKey" yaml:"appleMerchantPrivateKey" validate:"required"`
	// AppleSupportedNetworks Specifies which card networks will be accepted by ApplePay
	AppleSupportedNetworks []AppleNetwork `json:"appleSupportedNetworks" yaml:"appleSupportedNetworks"`
	// AppleMerchantCapabilities The payment capabilities supported by the merchant
	AppleMerchantCapabilities []AppleMerchantCapability `json:"appleMerchantCapabilities" yaml:"appleMerchantCapabilities"`
}

type (
	AppleNetwork            string
	AppleMerchantCapability string
)

const (
	AppleNetworkAmex            AppleNetwork = "amex"
	AppleNetworkCartesBancaires AppleNetwork = "cartesBancaires"
	AppleNetworkChinaUnionPay   AppleNetwork = "chinaUnionPay"
	AppleNetworkDiscover        AppleNetwork = "discover"
	AppleNetworkEftpos          AppleNetwork = "eftpos"
	AppleNetworkElectron        AppleNetwork = "electron"
	AppleNetworkElo             AppleNetwork = "elo"
	AppleNetworkInterac         AppleNetwork = "interac"
	AppleNetworkJcb             AppleNetwork = "jcb"
	AppleNetworkMada            AppleNetwork = "mada"
	AppleNetworkMaestro         AppleNetwork = "maestro"
	AppleNetworkMasterCard      AppleNetwork = "masterCard"
	AppleNetworkPrivateLabel    AppleNetwork = "privateLabel"
	AppleNetworkVisa            AppleNetwork = "visa"
	AppleNetworkVPay            AppleNetwork = "vPay"

	AppleMerchantCapability3DS    AppleMerchantCapability = "supports3DS"    // Required. This value must be supplied.
	AppleMerchantCapabilityCredit AppleMerchantCapability = "supportsCredit" // Optional. If present, only transactions that are categorized as credit cards are allowed.
	AppleMerchantCapabilityDebit  AppleMerchantCapability = "supportsDebit"  // Optional. If present, only transactions that are categorized as debit cards are allowed.
	AppleMerchantCapabilityEMV    AppleMerchantCapability = "supportsEMV"    // Include this value only if you support China Union Pay transactions.
)

func (c ApplePay) GetAppleMerchantPublicKey() string {
	if c.AppleMerchantCertificate == nil {
		return ""
	}
	return *c.AppleMerchantCertificate
}

func (c ApplePay) GetAppleMerchantPrivateKey() string {
	if c.AppleMerchantPrivateKey == nil {
		return ""
	}
	return *c.AppleMerchantPrivateKey
}
