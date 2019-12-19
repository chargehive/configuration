package selector

type Key string

const (
	KeyNow           Key = "now"
	KeyRandomPercent Key = "randompercent"
)

const (
	// Assemble Properties
	KeyChargeLabel                    Key = "charge.label"
	KeyChargeCurrentTransactionNumber Key = "charge.currentTransactionNumber"
	KeyChargeCurrentAttemptNumber     Key = "charge.currentAttemptNumber"
	KeyChargeRenewalNumber            Key = "charge.renewalNumber"
	KeyChargeInitiatedTime            Key = "charge.initiatedTime"
	KeyChargeLastAttemptTime          Key = "charge.lastAttemptTime"
	KeyChargeScheduleAttempts         Key = "charge.scheduleAttempts"
	KeyChargeAttemptedTransactions    Key = "charge.attemptedTransactions"
	KeyChargeFailedTransactions       Key = "charge.failedTransactions"
	KeyChargeMerchantSubscriptionID   Key = "charge.merchantSubscriptionID"
	KeyChargeCountry                  Key = "charge.country"
	KeyChargeRegion                   Key = "charge.region"
	KeyChargePeriod                   Key = "charge.period"
	KeyChargePreferredMethodType      Key = "charge.preferredMethodType"

	// Charge Definition
	KeyChargeIntent            Key = "charge.intent"
	KeyChargeContract          Key = "charge.contract"
	KeyChargeAmount            Key = "charge.amount"
	KeyChargeAmountUnits       Key = "charge.amount.units"
	KeyChargeAmountCurrency    Key = "charge.amount.currency"
	KeyChargeExpiryTime        Key = "charge.expiryTime"
	KeyChargeMerchantReference Key = "charge.merchantReference"
	KeyChargeReferences        Key = "charge.references"
	KeyChargeUserLocale        Key = "charge.userLocale"
	KeyChargeUserLocation      Key = "charge.userLocation"
	KeyChargeEnvironment       Key = "charge.environment"

	// Charge Meta
	KeyChargeInvoiceDate                Key = "charge.meta.invoiceDate"
	KeyChargeDueDate                    Key = "charge.meta.dueDate"
	KeyChargeDiscountAmount             Key = "charge.meta.discountAmount"
	KeyChargeDiscountAmountUnits        Key = "charge.meta.discountAmount.units"
	KeyChargeDiscountAmountCurrency     Key = "charge.meta.discountAmount.currency"
	KeyChargeDeliveryAmount             Key = "charge.meta.deliveryAmount"
	KeyChargeDeliveryAmountUnits        Key = "charge.meta.deliveryAmount.units"
	KeyChargeDeliveryAmountCurrency     Key = "charge.meta.deliveryAmount.currency"
	KeyChargeTaxAmount                  Key = "charge.meta.taxAmount"
	KeyChargeTaxAmountUnits             Key = "charge.meta.taxAmount.units"
	KeyChargeTaxAmountCurrency          Key = "charge.meta.taxAmount.currency"
	KeyChargeTotalAmount                Key = "charge.meta.totalAmount"
	KeyChargeTotalAmountUnits           Key = "charge.meta.totalAmount.units"
	KeyChargeTotalAmountCurrency        Key = "charge.meta.totalAmount.currency"
	KeyChargeIpAddress                  Key = "charge.meta.ipAddress"
	KeyChargeBillingAddress             Key = "charge.meta.billingAddress"
	KeyChargeBillingAddressLineOne      Key = "charge.meta.billingAddress.lineOne"
	KeyChargeBillingAddressLineTwo      Key = "charge.meta.billingAddress.lineTwo"
	KeyChargeBillingAddressLineThree    Key = "charge.meta.billingAddress.lineThree"
	KeyChargeBillingAddressTown         Key = "charge.meta.billingAddress.town"
	KeyChargeBillingAddressCounty       Key = "charge.meta.billingAddress.county"
	KeyChargeBillingAddressCountry      Key = "charge.meta.billingAddress.country"
	KeyChargeBillingAddressPostalCode   Key = "charge.meta.billingAddress.postalCode"
	KeyChargeBillingAddressFao          Key = "charge.meta.billingAddress.fao"
	KeyChargeBillingAddressCompanyName  Key = "charge.meta.billingAddress.companyName"
	KeyChargeDeliveryAddress            Key = "charge.meta.deliveryAddress"
	KeyChargeDeliveryAddressLineOne     Key = "charge.meta.deliveryAddress.lineOne"
	KeyChargeDeliveryAddressLineTwo     Key = "charge.meta.deliveryAddress.lineTwo"
	KeyChargeDeliveryAddressLineThree   Key = "charge.meta.deliveryAddress.lineThree"
	KeyChargeDeliveryAddressTown        Key = "charge.meta.deliveryAddress.town"
	KeyChargeDeliveryAddressCounty      Key = "charge.meta.deliveryAddress.county"
	KeyChargeDeliveryAddressCountry     Key = "charge.meta.deliveryAddress.country"
	KeyChargeDeliveryAddressPostalCode  Key = "charge.meta.deliveryAddress.postalCode"
	KeyChargeDeliveryAddressFao         Key = "charge.meta.deliveryAddress.fao"
	KeyChargeDeliveryAddressCompanyName Key = "charge.meta.deliveryAddress.companyName"
	KeyChargePerson                     Key = "charge.meta.person"
	KeyChargePersonTitle                Key = "charge.meta.person.title"
	KeyChargePersonFirstName            Key = "charge.meta.person.firstName"
	KeyChargePersonLastName             Key = "charge.meta.person.lastName"
	KeyChargePersonFullName             Key = "charge.meta.person.fullName"
	KeyChargePersonEmail                Key = "charge.meta.person.email"
	KeyChargePersonPhoneNumber          Key = "charge.meta.person.phoneNumber"
	KeyChargePersonLanguage             Key = "charge.meta.person.language"
	KeyChargeCompany                    Key = "charge.meta.company"
	KeyChargeCompanyName                Key = "charge.meta.company.name"
	KeyChargeCompanyEmail               Key = "charge.meta.company.email"
	KeyChargeCompanyPhoneNumber         Key = "charge.meta.company.phoneNumber"
	KeyChargeDelivery                   Key = "charge.meta.delivery"
	KeyChargeDeliveryStandard           Key = "charge.meta.delivery.standard"
	KeyChargeDeliveryType               Key = "charge.meta.delivery.type"
	KeyChargeDeliveryTrackingCode       Key = "charge.meta.delivery.trackingCode"
	KeyChargeDeliveryCourier            Key = "charge.meta.delivery.courier"
	KeyChargeDevice                     Key = "charge.meta.device"
	KeyChargeDeviceColorDepth           Key = "charge.meta.device.colorDepth"
	KeyChargeDeviceLanguage             Key = "charge.meta.device.language"
	KeyChargeDeviceTimezone             Key = "charge.meta.device.timezone"
	KeyChargeDeviceTimezoneOffsetMins   Key = "charge.meta.device.timezoneOffsetMins"
	KeyChargeDeviceUserAgent            Key = "charge.meta.device.userAgent"
	KeyChargeDeviceIpAddress            Key = "charge.meta.device.ipAddress"
	KeyChargeDeviceOs                   Key = "charge.meta.device.os"
	KeyChargeDeviceOsVersion            Key = "charge.meta.device.osVersion"
	KeyChargeDeviceBrowser              Key = "charge.meta.device.browser"
	KeyChargeDeviceBrowserVersion       Key = "charge.meta.device.browserVersion"
	KeyChargeItems                      Key = "charge.meta.items"
	KeyChargeItemsCount                 Key = "charge.meta.items.count"
	KeyChargeItemSubscriptionId         Key = "charge.meta.item.subscriptionId"
	KeyChargeItemRenewalNumber          Key = "charge.meta.item.renewalNumber"
	KeyChargeItemDuration               Key = "charge.meta.item.duration"
	KeyChargeItemStartDate              Key = "charge.meta.item.startDate"
	KeyChargeItemEndDate                Key = "charge.meta.item.endDate"
	KeyChargeItemProductType            Key = "charge.meta.item.productType"
	KeyChargeItemSkuType                Key = "charge.meta.item.skuType"
	KeyChargeItemQuantity               Key = "charge.meta.item.quantity"
	KeyChargeItemUnitPrice              Key = "charge.meta.item.unitPrice"
	KeyChargeItemUnitPriceUnits         Key = "charge.meta.item.unitPrice.units"
	KeyChargeItemUnitPriceCurrency      Key = "charge.meta.item.unitPrice.currency"
	KeyChargeItemTaxAmount              Key = "charge.meta.item.taxAmount"
	KeyChargeItemTaxAmountUnits         Key = "charge.meta.item.taxAmount.units"
	KeyChargeItemTaxAmountCurrency      Key = "charge.meta.item.taxAmount.currency"
	KeyChargeItemDiscountAmount         Key = "charge.meta.item.discountAmount"
	KeyChargeItemDiscountAmountUnits    Key = "charge.meta.item.discountAmount.units"
	KeyChargeItemDiscountAmountCurrency Key = "charge.meta.item.discountAmount.currency"
	KeyChargeItemName                   Key = "charge.meta.item.name"
	KeyChargeItemDescription            Key = "charge.meta.item.description"
	KeyChargeItemProductCode            Key = "charge.meta.item.productCode"
	KeyChargeItemSkuCode                Key = "charge.meta.item.skuCode"
	KeyChargeItemDelivery               Key = "charge.meta.item.delivery"
	KeyChargeItemDeliveryStandard       Key = "charge.meta.item.delivery.standard"
	KeyChargeItemDeliveryType           Key = "charge.meta.item.delivery.type"
	KeyChargeItemDeliveryTrackingCode   Key = "charge.meta.item.delivery.trackingCode"
	KeyChargeItemDeliveryCourier        Key = "charge.meta.item.delivery.courier"

	// Transaction
	KeyTransactionType          Key = "transaction.type"
	KeyTransactionPaymentType   Key = "transaction.payment.type"
	KeyTransactionPaymentScheme Key = "transaction.payment.scheme"
	KeyTransactionContract      Key = "transaction.contract"
	KeyTransactionFailureType   Key = "transaction.failure.type"
	KeyTransactionVerified      Key = "transaction.verified"
	KeyTransactionLiability     Key = "transaction.liability"
	KeyTransactionSubType       Key = "transaction.sub.type"
	KeyTransactionResult        Key = "transaction.result"

	// Transaction response
	KeyTransactionResponseCode            Key = "transaction.response.code"
	KeyTransactionResponseFailureType     Key = "transaction.response.failure.type"
	KeyTransactionResponseCategory        Key = "transaction.response.category"
	KeyTransactionResponseErrorType       Key = "transaction.response.error.type"
	KeyTransactionResponseOriginatingCode Key = "transaction.response.originating.code"
)

const (
	KeyMethodName          Key = "method.name"
	KeyMethodPaymentScheme Key = "method.paymentScheme"
	KeyMethodValidFrom     Key = "method.validFrom"
	KeyMethodExpiry        Key = "method.expiry"
	KeyMethodType          Key = "method.type"
	KeyMethodInfo          Key = "method.info"
)
