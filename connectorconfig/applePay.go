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
	AppleSupportedNetworks []string `json:"appleSupportedNetworks" yaml:"appleSupportedNetworks" validate:"omitempty,dive,oneof=amex cartesBancaires chinaUnionPay discover eftpos electron elo interac jcb mada maestro masterCard privateLabel visa vPay"`
	// AppleMerchantCapabilities The payment capabilities supported by the merchant
	AppleMerchantCapabilities []string `json:"appleMerchantCapabilities" yaml:"appleMerchantCapabilities" validate:"omitempty,dive,oneof=supports3DS supportsCredit supportsDebit supportsEMV"`
}

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
