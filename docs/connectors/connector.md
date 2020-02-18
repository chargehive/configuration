#Connector Config

The connector config defines the account and connection information to communicate with external payment gateways and fraud checking services

##Format
As with all configs, the standard wrapper is used.
```json5
{
  "kind": "Connector", //Must be set to "Connector"
  "metadata": {
    "projectId": "test-project", //Must be set to the ChargeHive Project ID you were issued with
    "name": "braintree-connector" //Set this to a memorable name for the connector, no spaces, all lowercase
  },
  "spec": {
    "library": "braintree", //Set this to the name of the library you wish to use
    "configuration": "eyJQdWJsaWN..." //Set this to the Base64 Encoded configuration json as featured below
  }
}
```

##Payment Libraries
Below is a list of the configuration options for each of the connectors for the payments gateways and fraud services.
The library config json needs to be Base64 encoded and inserted into the `Configuration` property of the library as a string
###Authorize.net 
Library: `authorize`
```json
{
  "APILoginID": "xxxxxxxxxxxx",
  "TransactionKey": "xxxxxxxxxxxx",
  "Environment": "xxxxxxxxxxxx"
}
```
FieldName | Definition 
---:|:---
APILoginID | Api Login ID
TransactionKey | Transaction Key
Environment | Either "sandbox" or "production"

###Braintree
Library: `braintree`
```json
{
  "PublicKey": "xxxxxxxxxxxx",
  "PrivateKey": "xxxxxxxxxxxx",
  "MerchantAccountID": "xxxxxxxxxxxx",
  "Currency": "EUR",
  "Environment": "sandbox"
}
```
FieldName | Definition 
---:|:---
PublicKey | Public encryption key
PrivateKey | Private encryption key
MerchantAccountID | Merchant Account ID
Currency | Account currency (e.g. "GBP", "EUR", "USD")
Environment | Either "sandbox" or "production"

###Paypal - Express Checkout
Library: `paypal-expresscheckout`
###Paypal - Website Payments Pro
Library: `paypal-websitepaymentspro`
###Paysafe
Library: `paysafe`
###Paysafe - ApplePay
Library: `paysafe-applepay`
###Paysafe - GooglePay
Library: `paysafe-googlepay`
###QualPay
Library: `qualpay`
###SandBox
Library: `sandbox`
```json
{
  "Mode": "dynamic",
  "TransactionIDPrefix": "xxxxxxxxxxxx"
}
```
FieldName | Definition 
---:|:---
Mode | "dynamic", "offline", "delayed", "random-timeout" or "chaos"
TransactionIDPrefix | Prepends transactions with this prefix

###Stripe
Library: `stripe`
###Vindicia
Library: `vindicia`
###WorldPay
Library: `worldpay`
##Fraud Libraries
###ChargeHive
Library: `chargehive`
###CyberSource
Library: `cybersource`
###MaxMind
Library: `maxmind`

##Full Example
Here's an working example using the sandbox connector with the configuration set to : `{"Mode":"dynamic","TransactionIDPrefix":"1234"}`
```json
{
  "kind": "Connector",
  "metadata": {
    "projectId": "test-project",
    "name": "sandbox-connector"
  },
  "spec": {
    "Library": "sandbox",
    "Configuration": "eyJNb2RlIjoiZHluYW1pYyIsIlRyYW5zYWN0aW9uSURQcmVmaXgiOiIxMjM0In0="
  }
}
```