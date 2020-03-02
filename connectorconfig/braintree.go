package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type BraintreeEnvironment string

const (
	BraintreeEnvironmentSandbox    BraintreeEnvironment = "sandbox"
	BraintreeEnvironmentProduction BraintreeEnvironment = "production"
)

type BraintreeCredentials struct {
	PublicKey         *string              `json:"publicKey" yaml:"publicKey" validate:"required"`
	PrivateKey        *string              `json:"privateKey" yaml:"privateKey" validate:"required"`
	MerchantID        string               `json:"merchantID" yaml:"merchantID" validate:"required"`
	MerchantAccountID string               `json:"merchantAccountID" yaml:"merchantAccountID" validate:"required"`
	Currency          string               `json:"currency" yaml:"currency" validate:"oneof=AED AMD AOA ARS AUD AWG AZN BAM BBD BDT BGN BIF BMD BND BOB BRL BSD BWP BYN BZD CAD CHF CLP CNY COP CRC CVE CZK DJF DKK DOP DZD EGP ETB EUR FJD FKP GBP GEL GHS GIP GMD GNF GTQ GYD HKD HNL HRK HTG HUF IDR ILS INR ISK JMD JPY KES KGS KHR KMF KRW KYD KZT LAK LBP LKR LRD LSL LTL MAD MDL MKD MNT MOP MUR MVR MWK MXN MYR MZN NAD NGN NIO NOK NPR NZD PAB PEN PGK PHP PKR PLN PYG QAR RON RSD RUB RWF SAR SBD SCR SEK SGD SHP SLL SOS SRD STD SVC SYP SZL THB TJS TOP TRY TTD TWD TZS UAH UGX USD UYU UZS VES VND VUV WST XAF XCD XOF XPF YER ZAR ZMK ZWD"`
	Environment       BraintreeEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
}

func (c BraintreeCredentials) GetLibrary() Library {
	return LibraryBraintree
}

func (c BraintreeCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *BraintreeCredentials) Validate() error {
	return nil
}

func (c *BraintreeCredentials) GetSecureFields() []*string {
	return []*string{c.PublicKey, c.PrivateKey}
}

func (c *BraintreeCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *BraintreeCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
