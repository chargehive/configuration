package selector

type Key string

const (
	KeyNow           Key = "now"
	KeyRandomPercent Key = "randompercent"
)

type RequestKey Key

const (
	RequestKeyType             RequestKey = "type"
	RequestKeyLabel            RequestKey = "label"
	RequestKeyContract         RequestKey = "contract"
	RequestKeyScheduleAttempt  RequestKey = "scheduleattempt"
	RequestKeyFailedAttempts   RequestKey = "failedattempts"
	RequestKeyRenewalNumber    RequestKey = "renewalnumber"
	RequestKeyAmount           RequestKey = "amount"
	RequestKeyCurrency         RequestKey = "currency"
	RequestKeyTaxAmount        RequestKey = "taxamount"
	RequestKeyTaxCurrency      RequestKey = "taxcurrency"
	RequestKeyDeliveryAmount   RequestKey = "deliveryamount"
	RequestKeyDeliveryCurrency RequestKey = "deliverycurrency"
	RequestKeyProductType      RequestKey = "producttype"
	RequestKeySkuType          RequestKey = "skutype"
	RequestKeySkuCode          RequestKey = "skucode"
	RequestKeyDeliveryType     RequestKey = "deliverytype"
	RequestKeyDeliveryStandard RequestKey = "deliverystandard"
	RequestKeyCountry          RequestKey = "country"
	RequestKeyRegion           RequestKey = "region"
	RequestKeyPeriod           RequestKey = "period"
	RequestKeyInitiated        RequestKey = "initiated"
	RequestKeyExpiry           RequestKey = "expiry"
	RequestKeyLastAttempted    RequestKey = "lastattempted"
)
