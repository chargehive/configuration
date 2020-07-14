package connectorconfig

type ApplePay struct {
	// AppleMerchantIdentifier REQUIRED TO ENABLE APPLE PAY Merchant Identifier specified in the Apple Developer Merchant section
	AppleMerchantIdentifier string `json:"appleMerchantIdentifier,omitempty" yaml:"appleMerchantIdentifier,omitempty" validate:"-"`
	// AppleMerchantDisplayName Value to be displayed on the payment page
	AppleMerchantDisplayName string `json:"appleMerchantDisplayName,omitempty" yaml:"appleMerchantDisplayName,omitempty" validate:"required_with=AppleMerchantIdentifier"`
	// AppleMerchantCertificate Merchant certificate in the Apple Developer Merchant section (must be base64 encoded!)
	AppleMerchantCertificate *string `json:"appleMerchantCertificate" yaml:"appleMerchantCertificate" validate:"required"`
	// AppleMerchantPrivateKey Merchant private key generated from the CSR in the Apple Developer Merchant section (must be base64 encoded!)
	AppleMerchantPrivateKey *string `json:"appleMerchantPrivateKey" yaml:"appleMerchantPrivateKey" validate:"required"`
	// AppleSupportedNetworks Specifies which card networks will be accepted by ApplePay
	AppleSupportedNetworks []AppleSupportedNetwork `json:"appleSupportedNetworks,omitempty" yaml:"appleSupportedNetworks,omitempty" validate:"required_with=AppleMerchantIdentifier,dive,oneof=amex cartesBancaires chinaUnionPay discover eftpos electron elo interac jcb mada maestro masterCard privateLabel visa vPay"`
	// AppleMerchantCapabilities The payment capabilities supported by the merchant
	AppleMerchantCapabilities []AppleMerchantCapability `json:"appleMerchantCapabilities,omitempty" yaml:"appleMerchantCapabilities,omitempty" validate:"required_with=AppleMerchantIdentifier,dive,oneof=supports3DS supportsCredit supportsDebit supportsEMV"`
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
