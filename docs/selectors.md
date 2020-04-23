# Configuration Selectors

Selectors are optional configurations that can be defined against all objects in ChargeHive and allow setting a series of Rules that much match against charge.

Selectors allow each config to be applied to all (if the selector is empty), or a specific subset of charges based on whether or not they match the Rules 
provided in the expressions section. 

A Selector has a `priority` which is used to determine a config to use if there are multiple matches. The higher the priority, then the more
likely that config will be used. 
A Selector also has a list of one or more `expressions`. These act as rules and ensure the config will only be applied to charges which match the expression rules.

If two or more configurations have Selector Rules in the expressions section which are matched to the Charge, the most specific Configuration will be selected (the one with the most Rules defined which match the charge).

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
[conversion](#conversion-values)|false|Can be used to convert the value from the key into another value. See list of conversions below
values|false|Simple array of strings to represent one or more values to compare


## Key Values
Special Keys | Definition 
---:|:---
"now"| Key returned will be the timestamp now
"randompercent"| Key will be a random integer between 0-100

Assemble Keys | Definition 
---:|:---
"charge.label"|
"charge.currentTransactionNumber"|
"charge.currentAttemptNumber"|
"charge.renewalNumber"| (int) Number of renewals this charge has completed.
"charge.initiatedTime"|
"charge.lastAttemptTime"|
"charge.scheduleAttempts"|
"charge.attemptedTransactions"|
"charge.failedTransactions"|
"charge.merchantSubscriptionID"|
"charge.country"|
"charge.region"|
"charge.period"|

Charge Definition Keys | Definition 
---:|:---
"charge.intent"|
"charge.contract"|
"charge.amount"|
"charge.amount.units"|
"charge.amount.currency"|
"charge.expiryTime"|
"charge.merchantReference"|
"charge.references"|
"charge.userLocale"|
"charge.userLocation"|
"charge.environment"|
"charge.preferredMethodType"|
	
Charge Meta Keys | Definition 
---:|:---
"charge.meta.invoiceDate"|
"charge.meta.dueDate"|
"charge.meta.discountAmount"|
"charge.meta.discountAmount.units"|
"charge.meta.discountAmount.currency"|
"charge.meta.deliveryAmount"|
"charge.meta.deliveryAmount.units"|
"charge.meta.deliveryAmount.currency"|
"charge.meta.taxAmount"|
"charge.meta.taxAmount.units"|
"charge.meta.taxAmount.currency"|
"charge.meta.totalAmount"|
"charge.meta.totalAmount.units"|
"charge.meta.totalAmount.currency"|
"charge.meta.ipAddress"|
"charge.meta.billingAddress"|
"charge.meta.billingAddress.lineOne"|
"charge.meta.billingAddress.lineTwo"|
"charge.meta.billingAddress.lineThree"|
"charge.meta.billingAddress.town"|
"charge.meta.billingAddress.county"|
"charge.meta.billingAddress.country"|
"charge.meta.billingAddress.postalCode"|
"charge.meta.billingAddress.fao"|
"charge.meta.billingAddress.companyName"|
"charge.meta.deliveryAddress"|
"charge.meta.deliveryAddress.lineOne"|
"charge.meta.deliveryAddress.lineTwo"|
"charge.meta.deliveryAddress.lineThree"|
"charge.meta.deliveryAddress.town"|
"charge.meta.deliveryAddress.county"|
"charge.meta.deliveryAddress.country"|
"charge.meta.deliveryAddress.postalCode"|
"charge.meta.deliveryAddress.fao"|
"charge.meta.deliveryAddress.companyName"|
"charge.meta.person"|
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
"charge.meta.item.taxAmount"|
"charge.meta.item.taxAmount.units"|
"charge.meta.item.taxAmount.currency"|
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
"method.name"|
"method.paymentScheme"|
"method.validFrom"|
"method.expiry"|
"method.type"|
"method.provider"|
"method.info"|


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

