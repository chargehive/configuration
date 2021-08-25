# Configuration Selectors

Selectors are optional configurations that can be defined against all objects in ChargeHive and allow setting a series of Rules that must match against the
charge.

A Selector has a `priority` which is used to determine a config to use if there are multiple matches. The higher the priority, then the more likely the config
will be used.

A Selector also has a list of one or more `expressions`. These act as rules and ensure the config will only be applied to charges which match the expression
rules.

If two or more configurations have Selector Rules in the expressions section which are matched to the Charge, the most specific Configuration will be selected (
the one with the most Rules defined which match the charge).

```json
{
  "priority": 0,
  "expressions": [
    {
      "key": "NameOfTheKey",
      "operator": "Equal",
      "conversion": "",
      "values": [
        "value1",
        "value2"
      ]
    }
  ]
}
```

## Expression Definition

FieldName | Required | Definition
---:|---|:---
[key](#key-values)|true|The key is the field to be compared to the value to find a match. List of available keys are below
[operator](#operator-values)|true|One of a fixed list of operators listed below
[conversion](#conversion-values)|false|Can be used to convert the value from the key into another value. This is typically used to convert the default time value into a specific time or date format. See list of conversions below
values|false|Simple array of strings to represent one or more values to compare

## Key Values

Special Keys | Definition
---:|:---
"now"| Key returned will be the timestamp now
"randompercent"| Key will be a random integer between 0-100

Assemble Keys | Definition
---:|:---
"charge.label"|
"charge.currentTransactionNumber"|(int64)
"charge.currentAttemptNumber"|(int64)
"charge.currentAttemptMethodCascadeNumber"|(int64) Cascade number for payment methods
"charge.currentAttemptConnectorCascadeNumber"|(int64) Cascade number for connectors
"charge.renewalNumber"| (int64) Number of renewals this charge has completed.
"charge.initiatedTime"| (time.Time)
"charge.lastAttemptTime"| (time.Time)
"charge.scheduleAttempts"| (int32)
"charge.attemptedTransactions"| (int32)
"charge.failedTransactions"| (int32)
"charge.merchantSubscriptionID"| (string)
"charge.country"| (int32 - ISO 3166)
"charge.region"| (string)
"charge.period"| (int64)

Charge Definition Keys | Definition
---:|:---
"charge.intent"| (int32 1-5) 1 = Add, 2 = Refresh, 3 = Verify, 4 = Capture, 5 = Refund
"charge.contract"| (int32 1-5) 1 = None, 2 = Payment, 3 = Subscription Initial, 4 = Subscription Renewal, 5 = OneClick
"charge.amount"| exists or not exists
"charge.amount.units"| (int64)
"charge.amount.currency"| (string)
"charge.expiryTime"| (time.Time)
"charge.merchantReference"| (string)
"charge.references"| (string [array])
"charge.userLocale"| (string)
"charge.userLocation"| (string)
"charge.environment"| (int32 0-4) 0 = Invalid, 1 = Retail, 2 = Ecommerce, 3 = Moto, 4 = Renewal
"charge.preferredMethodType"| (int32 0-7) 0 = Invalid, 1 = Card, 2 = PayPal, 3 = Direct Debit, 4 = Crypto Currency, 5 = Paysafe Apple Pay, 6 = Paysafe Google Pay, 7 = None

Charge Meta Keys | Definition
---:|:---
"charge.meta.invoiceDate"| (time.Time)
"charge.meta.dueDate"| (time.Time)
"charge.meta.discountAmount"| exists or not exists
"charge.meta.discountAmount.units"| (int64)
"charge.meta.discountAmount.currency"| (string)
"charge.meta.deliveryAmount"| exists or not exists
"charge.meta.deliveryAmount.units"| (int64)
"charge.meta.deliveryAmount.currency"| (string)
"charge.meta.taxAmount"| exists or not exists
"charge.meta.taxAmount.units"| (int64)
"charge.meta.taxAmount.currency"| (string)
"charge.meta.totalAmount"| exists or not exists
"charge.meta.totalAmount.units"| (int64)
"charge.meta.totalAmount.currency"| (string)
"charge.meta.ipAddress"| (string)
"charge.meta.billingAddress"| exists or not exists
"charge.meta.billingAddress.lineOne"| (string)
"charge.meta.billingAddress.lineTwo"| (string)
"charge.meta.billingAddress.lineThree"| (string)
"charge.meta.billingAddress.town"| (string)
"charge.meta.billingAddress.county"| (string)
"charge.meta.billingAddress.country"| (string)
"charge.meta.billingAddress.postalCode"| (string)
"charge.meta.billingAddress.fao"| (string)
"charge.meta.billingAddress.companyName"| (string)
"charge.meta.deliveryAddress"| exists or not exists
"charge.meta.deliveryAddress.lineOne"| (string)
"charge.meta.deliveryAddress.lineTwo"| (string)
"charge.meta.deliveryAddress.lineThree"| (string)
"charge.meta.deliveryAddress.town"| (string)
"charge.meta.deliveryAddress.county"| (string)
"charge.meta.deliveryAddress.country"| (string)
"charge.meta.deliveryAddress.postalCode"| (string)
"charge.meta.deliveryAddress.fao"| (string)
"charge.meta.deliveryAddress.companyName"| (string)
"charge.meta.person"| ()
"charge.meta.person.title"|
"charge.meta.person.firstName"|
"charge.meta.person.lastName"|
"charge.meta.person.fullName"|
"charge.meta.person.email"|
"charge.meta.person.phoneNumber"|
"charge.meta.person.language"|
"charge.meta.company"|
"charge.meta.company.name"|
"charge.meta.company.email"|
"charge.meta.company.phoneNumber"|
"charge.meta.delivery"|
"charge.meta.delivery.standard"|
"charge.meta.delivery.type"|
"charge.meta.delivery.trackingCode"|
"charge.meta.delivery.courier"|
"charge.meta.device"|
"charge.meta.device.colorDepth"|
"charge.meta.device.language"|
"charge.meta.device.timezone"|
"charge.meta.device.timezoneOffsetMins"|
"charge.meta.device.userAgent"|
"charge.meta.device.ipAddress"|
"charge.meta.device.os"|
"charge.meta.device.osVersion"|
"charge.meta.device.browser"|
"charge.meta.device.browserVersion"|
"charge.meta.items"|
"charge.meta.items.count"|
"charge.meta.item.subscriptionId"|
"charge.meta.item.renewalNumber"|
"charge.meta.item.duration"|
"charge.meta.item.startDate"|
"charge.meta.item.endDate"|
"charge.meta.item.productType"|
"charge.meta.item.skuType"|
"charge.meta.item.quantity"|
"charge.meta.item.unitPrice"|
"charge.meta.item.unitPrice.units"|
"charge.meta.item.unitPrice.currency"|
"charge.meta.item.taxAmount"| exists or not exists
"charge.meta.item.taxAmount.units"| (int64)
"charge.meta.item.taxAmount.currency"| (string)
"charge.meta.item.discountAmount"|
"charge.meta.item.discountAmount.units"|
"charge.meta.item.discountAmount.currency"|
"charge.meta.item.name"|
"charge.meta.item.description"|
"charge.meta.item.productCode"|
"charge.meta.item.skuCode"|
"charge.meta.item.delivery"|
"charge.meta.item.delivery.standard"|
"charge.meta.item.delivery.type"|
"charge.meta.item.delivery.trackingCode"|
"charge.meta.item.delivery.courier"|

Transaction Keys | Definition
---:|:---
"transaction.type"|
"transaction.payment.type"|
"transaction.payment.provider"|
"transaction.payment.scheme"|
"transaction.contract"|
"transaction.failure.type"|
"transaction.verified"|
"transaction.liability"|
"transaction.sub.type"|
"transaction.result"|
"transaction.connector.id"| (string) the id of the connector
"transaction.connector.library"| (string) the name of the library (e.g. "paysafe", "worldpay")

Transaction Response Keys | Definition
---:|:---
"transaction.response.code"|
"transaction.response.failure.type"|
"transaction.response.category"|
"transaction.response.error.type"|
"transaction.response.originating.code"|

Method Keys | Definition
---:|:---
"method.name"                    | Display Name for payment method e.g. "**** **** **** 0050"
"method.paymentScheme"           | Payment Scheme enum e.g. "card_visa" see [Payment Schemes](#Payment-Schemes)
"method.validFrom"               | Valid from date e.g. "1509984191" unix timestamp
"method.expiry"                  | Expiry date e.g. "1509984191" unix timestamp
"method.type"                    | Payment method type e.g. 0-4 PAYMENT_METHOD_TYPE_CARD
"method.provider"                | Payment method provider e.g. 0-5 PAYMENT_METHOD_PROVIDER_FORM
"method.info.card.last.four"     | Last four digits of payment method e.g. "1234"
"method.info.issuer"             | Issuing bank (useful for selecting sandbox cards) e.g. "ChargeHive SandBanx""
"method.info.account.holder"     | Name on card e.g. "John Smith"
"method.info.card.brand"         | Brand of card e.g. "VISA"
"method.info.card.number.length" | Number of digits on card e.g. "16"
"method.info.country"            | Country of issuance e.g. "GB"

### Payment Schemes

Scheme Key | Definition
---:|:---
"card_unknown"  | Unknown
"card_air_plus" | AirPlus
"card_american_express" | American Express
"card_aurore" |  Aurore
"card_carte_bancaire" | Carte Bancaire
"card_carte_bleue" |  Carte Blue
"card_dankort" |  DanKort
"card_diners_club" |  Diners Club
"card_discover" |  Discover
"card_ge_capital" |  GE Captial
"card_japanese_credit_bank" |  JCB
"card_maestro" |  Maestro
"card_master_card" | MasterCard
"card_uatp" | UATP
"card_visa" | Visa
"card_visa_debit" | Visa Debit
"card_visa_electron" | Visa Electron
"card_unionpay" | China UnionPay
"card_bancomat" | Bancomat
"card_bc_card" |  BC Card
"card_bca_card" |  BCA Card
"card_cabcharge" | Cab Charge
"card_eftpos" |  EFT POS
"card_eps" | EPS
"card_elo" | ELO
"card_forbrugsforeningen" |  FBF (Denmark loyalty club)
"card_girocard" | Girocard (German)
"card_interac" | Interac (Canadian)
"card_isracard" | Isracard (Israeli)
"card_mir" | Mir (Russia)
"card_meps" | MepsPay.com (Middle Eastern Payment Services)
"card_nets" | Nets (Denmark)
"card_pay_pak" |  PayPak (Pakistan)
"card_ru_pay" | RuPay (Russia)
"card_troy" | Troy (Turkish)
"card_v_pay" | VPay (visa Europe)
"card_verve" | Verve (Mastercard)
"paypal_personal" | Paypal Personal account payment
"paypal_business" | Paypal Business account payment
"bitcoin" | Bitcoin
"directdebit" | Direct Debit

## Operator Values

Value | Definition
---:|:---
"Equal"|Match if the key is exactly the same as first value
"NotEqual"|Match if the key is anything but the first value
"In"|Match if key is in the list of values (should only be used when there are multiple values)
"NotIn"|Match if the key is not in the list of values (should only be used when there are multiple values)
"Exists"|Match if the key is an array, and the value is in that array
"DoesNotExists"|Match if the key is an array, and the value is not in that array
"Gt"|Match if the key is greater than the value (only works if the key and value are numeric)
"Lt"|Match if the key is less than the value (only works if the key and value are numeric)

## Conversion Values

Value | Definition
---:|:---
""| No conversion
"TimeDow"|Day Of Week - Converts timestamp to "Mon","Tues"...
"TimeMonth"|Month - Converts timestamps to a month "Jan","Feb"...
"DurationSeconds"|Duration in Seconds
"DurationHours"|Duration in Hours
"DurationDays"|Duration in Days

#### Converting Time fields

The time fields on ChargeHive are in time.Time format, so when entering a time value into a Selector, you must use the format *Year-Month-DayTHour:Min:SecZ*

