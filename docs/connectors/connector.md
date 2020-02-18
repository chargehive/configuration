# Connector Config  
  
The connector config defines the account and connection information to communicate with external payment gateways and fraud checking services  
  
## Format  
As with all configs, the standard wrapper is used.  
```json5  
{  
  "kind": "Connector", //Must be set to "Connector"  
  "metadata": {  
  "projectId": "test-project", //Must be set to the ChargeHive Project ID you were issued with  
  "name": "braintree-connector" //Set this to a memorable name for the connector, no spaces, all lowercase  
 }, 
  "spec": {  
  "Library": "braintree", //Set this to the name of the library you wish to use  
  "Configuration": "eyJQdWJsaWN..." //Set this to the Base64 Encoded configuration json as featured below  
 }}  
```  
  
## Payment Libraries  
Below is a list of the configuration options for each of the connectors for the payments gateways and fraud services.  
The library configuration json needs to be Base64 encoded and inserted into the `Configuration` property of the library as a string  
### Authorize.net
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
APILoginID | The Api Login ID from your Authorize.net account found in Account -> API Credentials & Keys  
TransactionKey | The Transaction Key obtained from your Authorize.net account found in Account -> API Credentials & Keys  
Environment | Must be either "sandbox" or "production"  
  
### Braintree  
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
PublicKey | The Public encryption key from your Braintree Settings -> API page 
PrivateKey | The Private encryption key from your Braintree Settins -> API page 
MerchantAccountID | The Merchant Account ID defined in Braintree Settings -> Business  
Currency | The Account currency for the defined Merchant Account ID, in standard three character format (e.g. "GBP", "EUR", "USD")  
Environment | Must be either "sandbox" or "production"  
  
### Paypal - Express Checkout  
Library: `paypal-expresscheckout`  
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
APIUsername | API Username from PayPal API settings  
APIPassword | API Password from PayPal API settings  
APISignature | API Signature from PayPal API settings  
SupportedCurrencies | The currencies setup to be accepted in your PayPal account in an array (e.g. ["GBP", "USD", "EUR"])  
Environment | Must be either "sandbox" or "live"  
  
### Paypal - Website Payments Pro  
Library: `paypal-websitepaymentspro`  
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
*CardinalCommerce* is the PayPal partner for Strong Customer Authentication (SCA) which provides 3-D Secure (3DS) authentication for your PayPal Website Payments Pro account. As such the Cardinal authentication details are required as well.  
  
FieldName | Definition  
---:|:---  
APIUsername | API Username from PayPal API settings  
APIPassword | API Password from PayPal API settings  
APISignature | API Signature from PayPal API settings  
SupportedCurrencies | The currencies setup to be accepted in your PayPal account in an array (e.g. ["GBP", "USD", "EUR"])  
CardinalProcessorID | Your Processor Identification Code assigned by Cardinal when you registered  
CardinalMerchantID | Your Merchant Identification Code assigned by Cardinal when you registered  
CardinalTransactionPw | Your Cardinal Password as you configured it in Cardinal  
CardinalTransactionURL | The Transaction URL obtained from your Cardinal integration
CardinalAPIIdentifier | The Cardinal API Identifier obtained from the CardinalCommerce Integration Payments settings  
CardinalAPIKey | The Cardinal API Key obtained from the CardinalCommerce Integration Payments settings  
CardinalOrgUnitID | The Cardinal API OrgUnitID obtained from the CardinalCommerce Integration Payments settings  
Environment | Either "sandbox" or "live"  
  
### Paysafe  
Library: `paysafe`  
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

To be confirmed


### Vindicia  
Library: `vindicia`  
```json
{
  
}
```
FieldName | Definition  
---:|:---  

### WorldPay  
Library: `worldpay` 
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
```json
{
  
}
```
FieldName | Definition  
---:|:---    

### CyberSource  
Library: `cybersource` 
```json
{
  
}
```
FieldName | Definition  
---:|:---   

### MaxMind  
Library: `maxmind` 
```json
{
  
}
```
FieldName | Definition  
---:|:---   

##Full Example
Here's an working example using the sandbox connector with the configuration set to : `{"Mode":"dynamic","TransactionIDPrefix":"1234"}`
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