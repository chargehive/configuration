# Connectors
  
Connectors in ChargeHive are external services like Payment Providers or Fraud Services.
The connector config defines the account and connection information to communicate with these external payment gateways and fraud checking services.  

[Here](#full-example) is a working example of a Payment Provider Connector config using the Sandbox Connector.
  
## Format  
As with all configs, the standard wrapper is used:  
 
```json5
{
  "kind": "Connector",                        // Must be set to "Connector"
  "metadata": {
    "projectId": "test-project",              // Must be set to the ChargeHive Project ID you were issued with
    "name": "braintree-connector",            // Set this to a memorable name for the connector, no spaces, all lowercase
  },
  "specVersion": "v1",                        // Must be set to the correct version
  "selector": {},                             // May be used to apply this to a subset of charges
  "spec": {
    "Library": "braintree",                   // Set this to the name of the library you wish to use
    "configuration": "eyJQdWJsaWN..."         // Set this to the Base64 Encoded configuration json as featured below
  }
}
```
## Selectors
You can optionally apply Selector rules to Connectors to ensure they are only used in certain circumstances.  Selectors allow you to define Expressions that must be matched in order for a Charge to be attempted on this Connector.
  
  A good example of a Selector for a Connector would be Currency. Using a Selector you can ensure the Currency of the Charge must match the Currency defined in the Expression, so this Connector would only be used in a Charge Attempt if the Currency matched.
  
  Find out more about Selectors and what you can define in the [Selectors](../selectors.md) secton.

## Payment Libraries  
To create a Payment Provider Connector you need to define the `spec` properties in the config of the Connector.
* The `Library` property needs to be set to the Library value defined below for the Payment Provider or Fraud Service you are setting up.
* The `Configuration` property must have the library configuration json as defined below Base64 encoded and inserted as a string.

Here are the configuration options for each of the connectors for the Payment Provider and Fraud Services. 

---

#### Payment Providers
[Authorize.net](#authorizenet)  
[Braintree](#braintree)  
[PayPal Express Checkout](#paypal---express-checkout)  
[PayPal Website Payments Pro](#paypal---website-payments-pro)  
[Paysafe](#paysafe)  
[Qualpay](#qualpay)  
[Sandbox](#sandbox)  
[Stripe](#stripe)  
[Vindicia](#vindicia)
[Worldpay](#worldpay)  

#### Fraud Libraries
[ChargeHive](#chargehive)  
[Kount](#kount)  
[Cybersource](#cybersource)  
[Maxmind](#maxmind)

---

### Authorize.net
Library: `authorize`  
Configuration:
```json  
{  
  "APILoginID": "xxxxxxxxxxxx",  
  "TransactionKey": "xxxxxxxxxxxx",  
  "Environment": "xxxxxxxxxxxx"  
}  
```  
FieldName | Definition   
---:|:---  
APILoginID | The Api Login ID from your Authorize.net account found in Account -> API Credentials & Keys  
TransactionKey | The Transaction Key obtained from your Authorize.net account found in Account -> API Credentials & Keys  
Environment | Must be either "sandbox" or "production"  
  
### Braintree  
Library: `braintree`  
Configuration:
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
PublicKey | The Public encryption key from your Braintree Settings -> API page 
PrivateKey | The Private encryption key from your Braintree Settings -> API page 
MerchantAccountID | The Merchant Account ID defined in Braintree Settings -> Business  
Currency | The Account currency for the defined Merchant Account ID, in standard three character format (e.g. "GBP", "EUR", "USD")  
Environment | Must be either "sandbox" or "production"  
  
### Paypal - Express Checkout  
Library: `paypal-expresscheckout`  
Configuration:
```json  
{  
  "APIUsername": "xxxxxxxxxxxx",  
  "APIPassword": "xxxxxxxxxxxx",  
  "APISignature": "xxxxxxxxxxxx",  
  "SupportedCurrencies": ["USD","GBP"],  
  "Environment": "sandbox"  
}  
```  
FieldName | Definition  
---:|:---  
APIUsername | API Username from your PayPal API settings  
APIPassword | API Password from your PayPal API settings  
APISignature | API Signature from your PayPal API settings  
SupportedCurrencies | The currencies setup to be accepted in your PayPal account in an array (e.g. ["GBP", "USD", "EUR"])  
Environment | Must be either "sandbox" or "live"  
  
### Paypal - Website Payments Pro  
Library: `paypal-websitepaymentspro`  
Configuration:
```json  
{  
  "APIUsername": "xxxxxxxxxxxx",  
  "APIPassword": "xxxxxxxxxxxx",  
  "APISignature": "xxxxxxxxxxxx",  
  "SupportedCurrencies": ["USD","GBP"],  
  "CardinalProcessorID": "xxxxxxxxxxxx",  
  "CardinalMerchantID": "xxxxxxxxxxxx",  
  "CardinalTransactionPw": "xxxxxxxxxxxx",  
  "CardinalTransactionURL": "xxxxxxxxxxxx",   
 "CardinalAPIIdentifier": "xxxxxxxxxxxx",  
  "CardinalAPIKey": "xxxxxxxxxxxx",  
  "CardinalOrgUnitID": "xxxxxxxxxxxx",  
  "Environment":"sandbox"  
}  
```  
*CardinalCommerce* is the PayPal partner for Strong Customer Authentication (SCA) which provides 3-D Secure (3DS) authentication for your PayPal Website Payments Pro account. As such the Cardinal authentication details are required in this config as well.  
  
FieldName | Definition  
---:|:---  
APIUsername | API Username from your PayPal API settings  
APIPassword | API Password from your PayPal API settings  
APISignature | API Signature from your PayPal API settings  
SupportedCurrencies | The currencies setup to be accepted in your PayPal account in an array (e.g. ["GBP", "USD", "EUR"])  
CardinalProcessorID | Your Processor Identification Code assigned by Cardinal when you registered  
CardinalMerchantID | Your Merchant Identification Code assigned by Cardinal when you registered  
CardinalTransactionPw | Your Cardinal Password as you configured it in Cardinal  
CardinalTransactionURL | The Transaction URL obtained from your Cardinal integration
CardinalAPIIdentifier | The Cardinal API Identifier obtained from the CardinalCommerce Integration Payments settings  
CardinalAPIKey | The Cardinal API Key obtained from the CardinalCommerce Integration Payments settings  
CardinalOrgUnitID | The Cardinal API OrgUnitID obtained from the CardinalCommerce Integration Payments settings  
Environment | Must be either "sandbox" or "live"  
  
### Paysafe  
Library: `paysafe`  
Configuration:
```json
{
  "Acquirer":"xxxxxxxxxxxx",  
  "AccountID":"xxxxxxxxxxxx",  
  "APIUsername":"xxxxxxxxxxxx", 
  "APIPassword":"xxxxxxxxxxxxx",
  "Environment": "TEST",  
  "Country": "xxxxxxxxxxxxx",
  "Currency": "USD",  
  "UseVault": "false",  
  "SingleUseTokenPassword": "xxxxxxxxxxxx",  
  "SingleUseTokenUsername": "xxxxxxxxxxxx"
}
```
FieldName | Definition  
---:|:---  
Acquirer | The Acquirer bank setup for this merchant account
AccountID | The AccountID for this merchant account
APIUsername | The API Username from your merchant account in your account settings API page
APIPassword | The API Password from your merchant account in your account settings API page
Environment | Must be "MOCK", "TEST" or "LIVE"  
Country | Optional string field
Currency | The currency setup for this merchant account (e.g. ["GBP", "USD", "EUR"])
UseVault | Boolean field can be set to "true" or "false"  
SingleUseTokenPassword | The Single Use Token Password in your account settings API page
SingleUseTokenUsername | The Single Use Token Username in your account settings API page

### QualPay  
Library: `qualpay`  
Configuration:
```json
{
  "APIKey" : "xxxxxxxxxxxx",  
  "MerchantID" : "int64",  
  "Environment" : "test"
}
```
FieldName | Definition  
---:|:---  
APIKey | Your Qualpay API Key from your Qualpay account administration
MerchantID | This is your Qualpay Merchant ID. It must be a 64 character integer.
Environment | Must be "test" or "live" 

### SandBox  
Library: `sandbox`  
Configuration:
```json  
{  
  "Mode": "dynamic",  
  "TransactionIDPrefix": "xxxxxxxxxxxx"  
}  
```  
FieldName | Definition   
---:|:---  
Mode | Must be "dynamic", "offline", "delayed", "random-timeout" or "chaos"  
TransactionIDPrefix | Prepends transactions with this prefix  
  
### Stripe  
Library: `stripe`  
Configuration:
To be confirmed

### Vindicia  
Library: `vindicia`  
Configuration:
To be confirmed

### WorldPay  
Library: `worldpay`  
Configuration:
```json
{
  "Username": "xxxxxxxxxxxx",
  "Password": "xxxxxxxxxxxx",
  "MerchantID": "xxxxxxxxxxxx",
  "ReportGroup": "xxxxxxxxxxxx",
  "Environment": "prelive",
  "CardinalApiIdentifier": "xxxxxxxxxxxx",
  "CardinalOrgUnitId": "xxxxxxxxxxxx",
  "CardinalApiKey": "xxxxxxxxxxxx"
}
```

*CardinalCommerce* is the Worldpay partner for Strong Customer Authentication (SCA) which provides 3-D Secure (3DS) authentication for your Worldpay account. As such the Cardinal authentication details are required in this config as well.  

FieldName | Definition  
---:|:---   
  Username | The Username for this account in Worldpay
  Password | The Password of this account in Worldpay
  MerchantID | The MerchantID set for this account in Worldpay
  ReportGroup | Optional field to enter a string that can be used in Worldpay reporting to group transactions
  Environment | Must be "sandbox", "postlive", "transactpostlive", "production", "productiontransact", "prelive" or "transactprelive"
  CardinalAPIIdentifier | The Cardinal API Identifier obtained from the CardinalCommerce Integration Payments settings  
  CardinalOrgUnitID | The Cardinal API OrgUnitID obtained from the CardinalCommerce Integration Payments settings 
  CardinalAPIKey | The Cardinal API Key obtained from the CardinalCommerce Integration Payments settings  

## Fraud Libraries  

### ChargeHive  
Library: `chargehive`  
Configuration:
To be confirmed

### CyberSource  
Library: `cybersource`  
Configuration:
```json
{
 "MerchantID": "xxxxxxxxxxxx",
 "TransactionKey": "xxxxxxxxxxxx",
 "Environment": "test" 
}
```
FieldName | Definition  
---:|:---   
MerchantID | The Merchant ID on your Cybersource Account
TransactionKey | The Transaction Key generated in Cybersource on Payment Configuration -> Key Management
Environment | Must be either "test" or "live" 

### Kount
Library: `kount`  
Configuration:
```json
{
"siteID": "xxxxxxxx",
"merchantID": "xxxxxxxxxx",
"configKey": "xxxxxxxx",
"apiKey": "xxxxxxxxxxxx",
"dataCollectorURL": "xxxxxxx",
"riskInquiryServiceURL": "xxxxxxx",
"environment": "xxxxx"
}
```
FieldName | Definition
---:|:---   
siteID | The site ID on your Kount Account
merchantID | The Merchant Id of your Kount Account
configKey | The configuration key defined within your Kount account
apiKey | The API key defined in your Kount account
dataCollectorURL | The Kount data collection URL and will be test or a production URL depending on the environment
riskInquiryServiceURL | The Kount risk inquiry URL and will be test or a production URL depending on the environment
environment | Must be either "test" or "production"



### MaxMind  
Library: `maxmind`  
Configuration:
```json
{
  "AccountID": "xxxxxxxxxxxx",
  "LicenceKey": "xxxxxxxxxxxx"
}
```
FieldName | Definition  
---:|:---   
AccountID | The Account Id found in your MaxMind account Services -> My Licence Key
LicenceKey | The Licence Key from your MaxMind account under Services -> My Licence Key

---

## Full Example
This is a working example using the sandbox connector with the base 64 encoded configuration set to : `{"Mode":"dynamic","TransactionIDPrefix":"1234"}`
```json
{
  "kind": "Connector",
  "metadata": {
    "projectId": "test-project",
    "name": "sandbox-connector"
  },
  "specVersion": "v1",
  "selector": {},
  "spec": {
    "Library": "sandbox",
    "Configuration": "eyJNb2RlIjoiZHluYW1pYyIsIlRyYW5zYWN0aW9uSURQcmVmaXgiOiIxMjM0In0="
  }
}
```