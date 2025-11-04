package connectorconfig

import (
	"encoding/json"
	"strings"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type PaySafeEnvironment string

const (
	PaySafeEnvironmentMock PaySafeEnvironment = "MOCK"
	PaySafeEnvironmentTest PaySafeEnvironment = "TEST"
	PaySafeEnvironmentLive PaySafeEnvironment = "LIVE"
)

type PaysafeLocale string

const (
	PaysafeLocaleENGB PaysafeLocale = "en_GB"
	PaysafeLocaleENUS PaysafeLocale = "en_US"
	PaysafeLocaleFRCA PaysafeLocale = "fr_CA"
)

type PaySafeCredentials struct {
	Acquirer               string                `json:"acquirer" yaml:"acquirer" validate:"-"`
	MerchantURL            string                `json:"merchantURL" yaml:"merchantURL" validate:"required"`
	MerchantDescriptor     string                `json:"merchantDescriptor" yaml:"merchantDescriptor" validate:"-"`
	AccountID              string                `json:"accountID" yaml:"accountID" validate:"required"`
	APIUsername            *string               `json:"apiUsername" yaml:"apiUsername" validate:"required,gt=0"`
	APIPassword            *string               `json:"apiPassword" yaml:"apiPassword" validate:"required,gt=0"`
	Environment            PaySafeEnvironment    `json:"environment" yaml:"environment" validate:"oneof=MOCK TEST LIVE"`
	Country                string                `json:"country" yaml:"country" validate:"omitempty,oneof=AF AX AL DZ AS AD AO AI AQ AG AR AM AW AU AT AZ BS BH BD BB BY BE BZ BJ BM BT BO BQ BA BW BV BR IO BN BG BF BI KH CM CA CV KY CF TD CL CN CX CC CO KM CG CD CK CR CI HR CU CW CY CZ DK DJ DM DO EC EG SV GQ ER EE ET FK FO FJ FI FR GF PF TF GA GM GE DE GH GI GR GL GD GP GU GT GG GN GW GY HT HM HN HK HU IS IN ID IR IQ IE IM IL IT JM JP JE JO KZ KE KI KP KR KW KG LA LV LB LS LR LY LI LT LU MO MK MG MW MY MV ML MT MH MQ MR MU YT MX FM MD MC MN ME MS MA MZ MM NA NR NP NC NZ NI NE NG NU NF MP NO OM PK PW PS PA PG PY PE PH PN PL PT PR QA RE RO RU RW BL SH KN LC MF VC WS SM ST SA SN RS SC SL SG SX SK SI SB SO ZA GS SS ES LK PM SD SR SJ SZ SE CH SY TW TJ TZ TH NL TL TG TK TO TT TN TR TM TC TV UG UA AE GB US UM UY UZ VU VA VE VN VG VI WF EH YE ZM ZW"`
	Currency               string                `json:"currency" yaml:"currency" validate:"oneof=ARS AUD AZN BHD BYR BOB BAM BRL BGN CAD CLP CNY COP CRC HRK CZK DKK DOP XCD EGP ETB EUR FJD GEL GTQ HTG HNL HKD HUF ISK INR IDR IRR JMD JPY JOD KZT KES KRW KWD LVL LBP LYD LTL MWK MYR MUR MXN MDL MAD ILS NZD NGN NOK OMR PKR PAB PYG PEN PHP PLN GBP QAR RON RUB SAR RSD SGD ZAR LKR SEK CHF SYP TWD THB TTD TND TRY UAH AED UYU USD VEF VND"`
	UseVault               *bool                 `json:"useVault" yaml:"useVault" validate:"required"`
	SingleUseTokenPassword *string               `json:"singleUseTokenPassword" yaml:"singleUseTokenPassword" validate:"required"` // string* needs "required" to ensure nil is never returned
	SingleUseTokenUsername string                `json:"singleUseTokenUsername" yaml:"singleUseTokenUsername" validate:"-"`        // string will default to empty string
	GooglePay              *GooglePayCredentials `json:"googlePay,omitempty" yaml:"googlePay,omitempty"`
	ApplePay               *ApplePayCredentials  `json:"applePay,omitempty" yaml:"applePay,omitempty"`
}

func (c *PaySafeCredentials) GetGooglePayParams() map[string]string {
	return map[string]string{
		"gateway":           "paysafe",
		"gatewayMerchantId": c.GetGooglePay().GetGoogleCardMerchantId(),
	}
}

func (c *PaySafeCredentials) GetMID() string {
	return c.AccountID
}

func (c *PaySafeCredentials) GetGooglePay() *GooglePayCredentials {
	return c.GooglePay
}

func (c *PaySafeCredentials) GetApplePay() *ApplePayCredentials {
	return c.ApplePay
}

func (c *PaySafeCredentials) GetUseVault() bool {
	if c.UseVault == nil {
		return false
	}
	return *c.UseVault
}

func (c *PaySafeCredentials) GetLibrary() Library {
	return LibraryPaySafe
}

func (c *PaySafeCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *PaySafeCredentials) Validate() error {
	return nil
}

func (c *PaySafeCredentials) GetSecureFields() []*string {
	fields := []*string{c.APIUsername, c.APIPassword, c.SingleUseTokenPassword}
	if c.ApplePay != nil {
		fields = append(fields, c.ApplePay.AppleMerchantPrivateKey, c.ApplePay.AppleMerchantCertificate)
	}
	return fields
}

func (c *PaySafeCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *PaySafeCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *PaySafeCredentials) SupportsSca() bool {
	return c.Environment != "" && c.AccountID != "" && *c.SingleUseTokenPassword != "" && c.SingleUseTokenUsername != ""
}

func (c *PaySafeCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
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

var paysafeAllowedCountires = []string{"AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT", "AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ", "CA", "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU", "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE", "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB", "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS", "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT", "HU", "ID", "IE", "IL", "IM", "IN", "IO", "IQ", "IR", "IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG", "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI", "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG", "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP", "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN", "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW", "SA", "SB", "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR", "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF", "TG", "TH", "TJ", "TK", "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG", "UM", "US", "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "YE", "YT", "ZA", "ZM", "ZW"}

func (c *PaySafeCredentials) SupportsCountry(country string) bool {
	if country == "" {
		return true
	}
	for _, v := range paysafeAllowedCountires {
		if strings.EqualFold(v, country) {
			return true
		}
	}
	return false
}

func (c *PaySafeCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeSandbox && c.Environment == PaySafeEnvironmentLive {
		return false
	}
	return true
}

func (c *PaySafeCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *PaySafeCredentials) Supports3RI() bool {
	return false
}

func (c *PaySafeCredentials) IsAccountUpdater() bool {
	return false
}
