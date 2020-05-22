package connectorconfig

type ApplePay struct {
	// AppleMerchantIdentifier Merchant Identifier specified in the Apple Developer Merchant section
	AppleMerchantIdentifier string `json:"appleMerchantIdentifier" yaml:"appleMerchantIdentifier"`
	// AppleMerchantDisplayName Value to be displayed on the payment page
	AppleMerchantDisplayName string `json:"appleMerchantDisplayName" yaml:"appleMerchantDisplayName" validate:"required_with=AppleMerchantIdentifier"`
	// AppleMerchantCertificate Merchant certificate in the Apple Developer Merchant section (must be base64 encoded!)
	AppleMerchantCertificate *string `json:"appleMerchantCertificate" yaml:"appleMerchantCertificate" validate:"required"`
	// AppleMerchantPrivateKey Merchant private key generated from the CSR in the Apple Developer Merchant section (must be base64 encoded!)
	AppleMerchantPrivateKey *string `json:"appleMerchantPrivateKey" yaml:"appleMerchantPrivateKey" validate:"required"`
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
