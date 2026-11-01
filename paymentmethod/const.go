package paymentmethod

// Key is a key used to tag information about a payment method
type InfoKey string

// InfoKey constants
const (
	InfoKeyPaymentScheme      InfoKey = "payment_scheme"
	InfoKeyCardBrand          InfoKey = "card_brand"
	InfoKeyCardSubbrand       InfoKey = "card_subbrand" // a CardSubbrand
	InfoKeyIssuer             InfoKey = "issuer"
	InfoKeyIssueNumber        InfoKey = "issue_number"
	InfoKeyCardType           InfoKey = "card_type"
	InfoKeyCardBin            InfoKey = "card_bin"
	InfoKeyCardLastFour       InfoKey = "card_last_four"
	InfoKeyCardNumberLength   InfoKey = "card_number_length"
	InfoKeyDisplayName        InfoKey = "display_name"
	InfoKeyCountry            InfoKey = "country"
	InfoKeyExpiry             InfoKey = "expiry"
	InfoKeyValidFrom          InfoKey = "valid_from"
	InfoKeyLookupCode         InfoKey = "lookup_code"
	InfoKeyAccountHolder      InfoKey = "account_holder"
	InfoKeyMethodVersion      InfoKey = "payment_method_version"
	InfoKeyPayerReference     InfoKey = "payer_reference"
	InfoKeyAccountNumberLast2 InfoKey = "account_number_last_two"
	InfoKeySortCode           InfoKey = "sort_code"
	// InfoKeyConnectorToken token for this payment method on the connector
	InfoKeyConnectorToken InfoKey = "connector_token"
	// InfoKeyNetworkTransactionID deprecated - should be migrated into a sequence
	InfoKeyNetworkTransactionID     InfoKey = "network_transaction_id"
	InfoKeyAuthenticationCryptogram InfoKey = "authentication_cryptogram"
	InfoKeyAuthenticationECI        InfoKey = "authentication_eci"

	// InfoKeyCardNetworks is the networks a card can route over, comma separated, its
	// brand's own first, e.g. "VISA,ACCEL,STAR".
	InfoKeyCardNetworks   InfoKey = "card_networks"
	InfoKeyCardFunding    InfoKey = "card_funding"
	InfoKeyCardPrepaid    InfoKey = "card_prepaid"
	InfoKeyCardCommercial InfoKey = "card_commercial"
	InfoKeyCardConsumer   InfoKey = "card_consumer"
	InfoKeyCardSegment    InfoKey = "card_segment"
	// InfoKeyCardProduct and InfoKeyCardProductCode are the card's product within its brand
	// (its tier, e.g. "WORLD ELITE"), by name and by the brand's code (e.g. "MWE").
	InfoKeyCardProduct     InfoKey = "card_product"
	InfoKeyCardProductCode InfoKey = "card_product_code"
)

// Scheme us a value used to indicate a payment scheme type
type Scheme string

// Schemes
const (
	// Card
	SchemeCardUnknown            Scheme = "card_unknown"
	SchemeCardUnbranded          Scheme = "card_unbranded" // issued without a scheme: EBT, private-label and similar rails only
	SchemeCardAirPlus            Scheme = "card_air_plus"
	SchemeCardAmericanExpress    Scheme = "card_american_express"
	SchemeCardAurore             Scheme = "card_aurore"
	SchemeCardCarteBancaire      Scheme = "card_carte_bancaire"
	SchemeCardCarteBleue         Scheme = "card_carte_bleue"
	SchemeCardDankort            Scheme = "card_dankort"
	SchemeCardDinersClub         Scheme = "card_diners_club"
	SchemeCardDiscover           Scheme = "card_discover"
	SchemeCardGECapital          Scheme = "card_ge_capital"
	SchemeCardJCB                Scheme = "card_japanese_credit_bank"
	SchemeCardMaestro            Scheme = "card_maestro"
	SchemeCardMasterCard         Scheme = "card_master_card"
	SchemeCardUATP               Scheme = "card_uatp"
	SchemeCardVisa               Scheme = "card_visa"
	SchemeCardVisaDebit          Scheme = "card_visa_debit"
	SchemeCardVisaElectron       Scheme = "card_visa_electron"
	SchemeCardUnionPay           Scheme = "card_unionpay"
	SchemeCardBancomat           Scheme = "card_bancomat"
	SchemeCardBCCard             Scheme = "card_bc_card"
	SchemeCardBCACard            Scheme = "card_bca_card"
	SchemeCardCabcharge          Scheme = "card_cabcharge"
	SchemeCardEftpos             Scheme = "card_eftpos"
	SchemeCardEps                Scheme = "card_eps"
	SchemeCardElo                Scheme = "card_elo"
	SchemeCardForbrugsforeningen Scheme = "card_forbrugsforeningen"
	SchemeCardGirocard           Scheme = "card_girocard"
	SchemeCardInterac            Scheme = "card_interac"
	SchemeCardIsracard           Scheme = "card_isracard"
	SchemeCardMir                Scheme = "card_mir"
	SchemeCardMEPS               Scheme = "card_meps"
	SchemeCardNETS               Scheme = "card_nets"
	SchemeCardPayPak             Scheme = "card_pay_pak"
	SchemeCardRuPay              Scheme = "card_ru_pay"
	SchemeCardTroy               Scheme = "card_troy"
	SchemeCardVpay               Scheme = "card_v_pay"
	SchemeCardVerve              Scheme = "card_verve"

	// Paypal
	SchemePaypalPersonal Scheme = "paypal_personal"
	SchemePaypalBusiness Scheme = "paypal_business"

	// Crypto
	SchemeCryptoBitcoin Scheme = "bitcoin"

	// Direct Debit
	SchemeDirectDebit Scheme = "directdebit"
)

// CardType is a value used to indicate a payment card type
type CardType string

// CardTypes
const (
	CardTypeUnknown     CardType = "unknown"
	CardTypeCredit      CardType = "credit"
	CardTypeCharge      CardType = "charge"
	CardTypeDebit       CardType = "debit"
	CardTypePrepaid     CardType = "prepaid"
	CardTypeGift        CardType = "gift"
	CardTypeCorporate   CardType = "corporate"
	CardTypeGovernment  CardType = "government"
	CardTypeProprietary CardType = "proprietary"
)

type FundingType string

const (
	FundingTypeUnknown              FundingType = "unknown"
	FundingTypeCredit               FundingType = "credit"
	FundingTypeDebit                FundingType = "debit"
	FundingTypeCharge               FundingType = "charge"
	FundingTypeDeferredDebit        FundingType = "deferred_debit"
	FundingTypePrepaid              FundingType = "prepaid"
	FundingTypePrepaidReloadable    FundingType = "prepaid_reloadable"
	FundingTypePrepaidNonReloadable FundingType = "prepaid_non_reloadable"
)

// CardSubbrand is an acceptance brand within a card's brand: a mark of its own that a
// merchant can accept or decline apart from the brand, identified by the brand's product
// codes. Maestro and Cirrus (ATM only) within Mastercard, Visa Electron and V PAY within
// Visa. Tiers such as Classic, Gold or Platinum are the card's product, not a subbrand.
type CardSubbrand string

const (
	CardSubbrandMaestro      CardSubbrand = "MAESTRO"
	CardSubbrandCirrus       CardSubbrand = "CIRRUS"
	CardSubbrandVisaElectron CardSubbrand = "VISA_ELECTRON"
	CardSubbrandVPay         CardSubbrand = "V_PAY"
)

// BACSReportType indicate a type of BACS direct debit report
type BACSReportType string

// BACSReportType
const (
	AUDDIS BACSReportType = "AUDDIS" // rejected registrations
	DDICA  BACSReportType = "DDICA"  // chargebacks
	ADDACS BACSReportType = "ADDACS" // account updates/changes
	ARUDD  BACSReportType = "ARUDD"  // failed payments
)
