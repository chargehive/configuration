package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type PaySafeGooglePayCredentials struct {
	Acquirer               string             `json:"acquirer" yaml:"acquirer" validate:"required"`
	AccountID              string             `json:"accountID" yaml:"accountID" validate:"required"`
	APIUsername            *string            `json:"apiUsername" yaml:"apiUsername" validate:"required"`
	APIPassword            *string            `json:"apiPassword" yaml:"apiPassword" validate:"required"`
	Environment            PaySafeEnvironment `json:"environment" yaml:"environment" validate:"oneof=MOCK TEST LIVE"`
	Country                string             `json:"country" yaml:"country" validate:"omitempty,oneof=AF AX AL DZ AS AD AO AI AQ AG AR AM AW AU AT AZ BS BH BD BB BY BE BZ BJ BM BT BO BQ BA BW BV BR IO BN BG BF BI KH CM CA CV KY CF TD CL CN CX CC CO KM CG CD CK CR CI HR CU CW CY CZ DK DJ DM DO EC EG SV GQ ER EE ET FK FO FJ FI FR GF PF TF GA GM GE DE GH GI GR GL GD GP GU GT GG GN GW GY HT HM HN HK HU IS IN ID IR IQ IE IM IL IT JM JP JE JO KZ KE KI KP KR KW KG LA LV LB LS LR LY LI LT LU MO MK MG MW MY MV ML MT MH MQ MR MU YT MX FM MD MC MN ME MS MA MZ MM NA NR NP NC NZ NI NE NG NU NF MP NO OM PK PW PS PA PG PY PE PH PN PL PT PR QA RE RO RU RW BL SH KN LC MF VC WS SM ST SA SN RS SC SL SG SX SK SI SB SO ZA GS SS ES LK PM SD SR SJ SZ SE CH SY TW TJ TZ TH NL TL TG TK TO TT TN TR TM TC TV UG UA AE GB US UM UY UZ VU VA VE VN VG VI WF EH YE ZM ZW"`
	Currency               string             `json:"currency" yaml:"currency" validate:"oneof=ARS AUD AZN BHD BYR BOB BAM BRL BGN CAD CLP CNY COP CRC HRK CZK DKK DOP XCD EGP ETB EUR FJD GEL GTQ HTG HNL HKD HUF ISK INR IDR IRR JMD JPY JOD KZT KES KRW KWD LVL LBP LYD LTL MWK MYR MUR MXN MDL MAD ILS NZD NGN NOK OMR PKR PAB PYG PEN PHP PLN GBP QAR RON RUB SAR RSD SGD ZAR LKR SEK CHF SYP TWD THB TTD TND TRY UAH AED UYU USD VEF VND"`
	SingleUseTokenPassword *string            `json:"singleUseTokenPassword" yaml:"singleUseTokenPassword" validate:"required_with=singleUseTokenUsername"`
	SingleUseTokenUsername *string            `json:"singleUseTokenUsername" yaml:"singleUseTokenUsername" validate:"required_with=singleUseTokenPassword"`
	Locale                 PaysafeLocale      `json:"locale" yaml:"locale" validate:"oneof=en_GB en_US fr_CA"`
	// Pending implementation <additional parameters TBC>
}

func (c PaySafeGooglePayCredentials) GetLibrary() Library {
	return LibraryPaySafeGooglePay
}

func (c *PaySafeGooglePayCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *PaySafeGooglePayCredentials) Validate() error {
	return nil
}

func (c *PaySafeGooglePayCredentials) GetSecureFields() []*string {
	return []*string{c.APIUsername, c.APIPassword, c.SingleUseTokenUsername, c.SingleUseTokenPassword}
}

func (c *PaySafeGooglePayCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *PaySafeGooglePayCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}