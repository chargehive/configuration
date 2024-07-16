package selector

type Key string

const (
	KeyNow           Key = "now"
	KeyRandomPercent Key = "randompercent"
)

const ( // Assemble Properties
	KeyChargeLabel                                Key = "charge.label"
	KeyChargeCurrentTransactionNumber             Key = "charge.currentTransactionNumber"
	KeyChargeCurrentAttemptNumber                 Key = "charge.currentAttemptNumber"
	KeyChargeCurrentAttemptMethodCascadeNumber    Key = "charge.currentAttemptMethodCascadeNumber"
	KeyChargeCurrentAttemptConnectorCascadeNumber Key = "charge.currentAttemptConnectorCascadeNumber"
	KeyChargeRenewalNumber                        Key = "charge.renewalNumber"
	KeyChargeInitiatedTime                        Key = "charge.initiatedTime"
	KeyChargeLastAttemptTime                      Key = "charge.lastAttemptTime"
	KeyChargeScheduleAttempts                     Key = "charge.scheduleAttempts"
	KeyChargeAttemptedTransactions                Key = "charge.attemptedTransactions"
	KeyChargeFailedTransactions                   Key = "charge.failedTransactions"
	KeyChargeMerchantSubscriptionID               Key = "charge.merchantSubscriptionID"
	KeyChargeCountry                              Key = "charge.country"
	KeyChargeRegion                               Key = "charge.region"
	KeyChargePeriod                               Key = "charge.period"
)

const ( // Charge Definition
	KeyChargeIntent              Key = "charge.intent"
	KeyChargeContract            Key = "charge.contract"
	KeyChargeAmount              Key = "charge.amount"
	KeyChargeAmountUnits         Key = "charge.amount.units"
	KeyChargeAmountCurrency      Key = "charge.amount.currency"
	KeyChargeExpiryTime          Key = "charge.expiryTime"
	KeyChargeMerchantReference   Key = "charge.merchantReference"
	KeyChargeReferences          Key = "charge.references"
	KeyChargeUserLocale          Key = "charge.userLocale"
	KeyChargeUserLocation        Key = "charge.userLocation"
	KeyChargeEnvironment         Key = "charge.environment"
	KeyChargePreferredMethodType Key = "charge.preferredMethodType"
)

const ( // Charge Meta
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
	KeyChargePlacementID                Key = "charge.meta.PlacementID"
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
	KeyChargeItemTermUnits              Key = "charge.meta.item.termUnits"
	KeyChargeItemTermType               Key = "charge.meta.item.termType"
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
)

const ( // Transaction
	KeyTransactionType             Key = "transaction.type"
	KeyTransactionPaymentType      Key = "transaction.payment.type"
	KeyTransactionPaymentProvider  Key = "transaction.payment.provider"
	KeyTransactionPaymentScheme    Key = "transaction.payment.scheme"
	KeyTransactionContract         Key = "transaction.contract"
	KeyTransactionFailureType      Key = "transaction.failure.type"
	KeyTransactionVerified         Key = "transaction.verified"
	KeyTransactionLiability        Key = "transaction.liability"
	KeyTransactionSubType          Key = "transaction.sub.type"
	KeyTransactionResult           Key = "transaction.result"
	KeyTransactionConnectorId      Key = "transaction.connector.id"
	KeyTransactionConnectorLibrary Key = "transaction.connector.library"
)

const ( // Transaction response
	KeyTransactionResponseCode            Key = "transaction.response.code"
	KeyTransactionResponseFailureType     Key = "transaction.response.failure.type"
	KeyTransactionResponseCategory        Key = "transaction.response.category"
	KeyTransactionResponseErrorType       Key = "transaction.response.error.type"
	KeyTransactionResponseOriginatingCode Key = "transaction.response.originating.code"
)

const ( // Payment Method
	KeyMethodName                 Key = "method.name"                    // e.g. "**** **** **** 0050"
	KeyMethodPaymentScheme        Key = "method.paymentScheme"           // e.g. "card_visa"
	KeyMethodValidFrom            Key = "method.validFrom"               // e.g. "1509984191" unix timestamp
	KeyMethodExpiry               Key = "method.expiry"                  // e.g. "1509984191" unix timestamp
	KeyMethodType                 Key = "method.type"                    // e.g. 0-4 PAYMENT_METHOD_TYPE_CARD
	KeyMethodProvider             Key = "method.provider"                // e.g. 0-5 PAYMENT_METHOD_PROVIDER_FORM
	KeyMethodInfoCardBin          Key = "method.info.card.bin"           // e.g. "998981" 6 digits
	KeyMethodInfoCardLastFour     Key = "method.info.card.last.four"     // e.g. "1234"
	KeyMethodInfoIssuer           Key = "method.info.issuer"             // e.g. "credit"
	KeyMethodInfoAccountHolder    Key = "method.info.account.holder"     // e.g. "John Smith"
	KeyMethodInfoCardBrand        Key = "method.info.card.brand"         // e.g. "VISA"
	KeyMethodInfoCardNumberLength Key = "method.info.card.number.length" // e.g. "16"
	KeyMethodInfoCountry          Key = "method.info.country"            // e.g. "GB"
)

var KeyRegister = map[Key]bool{
	KeyNow:                            true,
	KeyRandomPercent:                  true,
	KeyChargeLabel:                    true,
	KeyChargeCurrentTransactionNumber: true,
	KeyChargeCurrentAttemptNumber:     true,
	KeyChargeCurrentAttemptMethodCascadeNumber:    true,
	KeyChargeCurrentAttemptConnectorCascadeNumber: true,
	KeyChargeRenewalNumber:                        true,
	KeyChargeInitiatedTime:                        true,
	KeyChargeLastAttemptTime:                      true,
	KeyChargeScheduleAttempts:                     true,
	KeyChargeAttemptedTransactions:                true,
	KeyChargeFailedTransactions:                   true,
	KeyChargeMerchantSubscriptionID:               true,
	KeyChargeCountry:                              true,
	KeyChargeRegion:                               true,
	KeyChargePeriod:                               true,
	KeyChargeIntent:                               true,
	KeyChargeContract:                             true,
	KeyChargeAmount:                               true,
	KeyChargeAmountUnits:                          true,
	KeyChargeAmountCurrency:                       true,
	KeyChargeExpiryTime:                           true,
	KeyChargeMerchantReference:                    true,
	KeyChargeReferences:                           true,
	KeyChargeUserLocale:                           true,
	KeyChargeUserLocation:                         true,
	KeyChargeEnvironment:                          true,
	KeyChargePreferredMethodType:                  true,
	KeyChargeInvoiceDate:                          true,
	KeyChargeDueDate:                              true,
	KeyChargeDiscountAmount:                       true,
	KeyChargeDiscountAmountUnits:                  true,
	KeyChargeDiscountAmountCurrency:               true,
	KeyChargeDeliveryAmount:                       true,
	KeyChargeDeliveryAmountUnits:                  true,
	KeyChargeDeliveryAmountCurrency:               true,
	KeyChargeTaxAmount:                            true,
	KeyChargeTaxAmountUnits:                       true,
	KeyChargeTaxAmountCurrency:                    true,
	KeyChargeTotalAmount:                          true,
	KeyChargeTotalAmountUnits:                     true,
	KeyChargeTotalAmountCurrency:                  true,
	KeyChargeIpAddress:                            true,
	KeyChargeBillingAddress:                       true,
	KeyChargeBillingAddressLineOne:                true,
	KeyChargeBillingAddressLineTwo:                true,
	KeyChargeBillingAddressLineThree:              true,
	KeyChargeBillingAddressTown:                   true,
	KeyChargeBillingAddressCounty:                 true,
	KeyChargeBillingAddressCountry:                true,
	KeyChargeBillingAddressPostalCode:             true,
	KeyChargeBillingAddressFao:                    true,
	KeyChargeBillingAddressCompanyName:            true,
	KeyChargeDeliveryAddress:                      true,
	KeyChargeDeliveryAddressLineOne:               true,
	KeyChargeDeliveryAddressLineTwo:               true,
	KeyChargeDeliveryAddressLineThree:             true,
	KeyChargeDeliveryAddressTown:                  true,
	KeyChargeDeliveryAddressCounty:                true,
	KeyChargeDeliveryAddressCountry:               true,
	KeyChargeDeliveryAddressPostalCode:            true,
	KeyChargeDeliveryAddressFao:                   true,
	KeyChargeDeliveryAddressCompanyName:           true,
	KeyChargePerson:                               true,
	KeyChargePersonTitle:                          true,
	KeyChargePersonFirstName:                      true,
	KeyChargePersonLastName:                       true,
	KeyChargePersonFullName:                       true,
	KeyChargePersonEmail:                          true,
	KeyChargePersonPhoneNumber:                    true,
	KeyChargePersonLanguage:                       true,
	KeyChargeCompany:                              true,
	KeyChargeCompanyName:                          true,
	KeyChargeCompanyEmail:                         true,
	KeyChargeCompanyPhoneNumber:                   true,
	KeyChargeDelivery:                             true,
	KeyChargeDeliveryStandard:                     true,
	KeyChargeDeliveryType:                         true,
	KeyChargeDeliveryTrackingCode:                 true,
	KeyChargeDeliveryCourier:                      true,
	KeyChargeDevice:                               true,
	KeyChargeDeviceColorDepth:                     true,
	KeyChargeDeviceLanguage:                       true,
	KeyChargeDeviceTimezone:                       true,
	KeyChargeDeviceTimezoneOffsetMins:             true,
	KeyChargeDeviceUserAgent:                      true,
	KeyChargeDeviceIpAddress:                      true,
	KeyChargeDeviceOs:                             true,
	KeyChargeDeviceOsVersion:                      true,
	KeyChargeDeviceBrowser:                        true,
	KeyChargeDeviceBrowserVersion:                 true,
	KeyChargePlacementID:                          true,
	KeyChargeItems:                                true,
	KeyChargeItemsCount:                           true,
	KeyChargeItemSubscriptionId:                   true,
	KeyChargeItemRenewalNumber:                    true,
	KeyChargeItemTermUnits:                        true,
	KeyChargeItemTermType:                         true,
	KeyChargeItemDuration:                         true,
	KeyChargeItemStartDate:                        true,
	KeyChargeItemEndDate:                          true,
	KeyChargeItemProductType:                      true,
	KeyChargeItemSkuType:                          true,
	KeyChargeItemQuantity:                         true,
	KeyChargeItemUnitPrice:                        true,
	KeyChargeItemUnitPriceUnits:                   true,
	KeyChargeItemUnitPriceCurrency:                true,
	KeyChargeItemTaxAmount:                        true,
	KeyChargeItemTaxAmountUnits:                   true,
	KeyChargeItemTaxAmountCurrency:                true,
	KeyChargeItemDiscountAmount:                   true,
	KeyChargeItemDiscountAmountUnits:              true,
	KeyChargeItemDiscountAmountCurrency:           true,
	KeyChargeItemName:                             true,
	KeyChargeItemDescription:                      true,
	KeyChargeItemProductCode:                      true,
	KeyChargeItemSkuCode:                          true,
	KeyChargeItemDelivery:                         true,
	KeyChargeItemDeliveryStandard:                 true,
	KeyChargeItemDeliveryType:                     true,
	KeyChargeItemDeliveryTrackingCode:             true,
	KeyChargeItemDeliveryCourier:                  true,
	KeyTransactionType:                            true,
	KeyTransactionPaymentType:                     true,
	KeyTransactionPaymentProvider:                 true,
	KeyTransactionPaymentScheme:                   true,
	KeyTransactionContract:                        true,
	KeyTransactionFailureType:                     true,
	KeyTransactionVerified:                        true,
	KeyTransactionLiability:                       true,
	KeyTransactionSubType:                         true,
	KeyTransactionResult:                          true,
	KeyTransactionConnectorId:                     true,
	KeyTransactionConnectorLibrary:                true,
	KeyTransactionResponseCode:                    true,
	KeyTransactionResponseFailureType:             true,
	KeyTransactionResponseCategory:                true,
	KeyTransactionResponseErrorType:               true,
	KeyTransactionResponseOriginatingCode:         true,
	KeyMethodName:                                 true,
	KeyMethodPaymentScheme:                        true,
	KeyMethodValidFrom:                            true,
	KeyMethodExpiry:                               true,
	KeyMethodType:                                 true,
	KeyMethodProvider:                             true,
	KeyMethodInfoCardBin:                          true,
	KeyMethodInfoCardLastFour:                     true,
	KeyMethodInfoIssuer:                           true,
	KeyMethodInfoAccountHolder:                    true,
	KeyMethodInfoCardBrand:                        true,
	KeyMethodInfoCardNumberLength:                 true,
	KeyMethodInfoCountry:                          true,
}
