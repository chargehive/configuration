# Cascade Policy
The cascade policy defines whether or not a failed attempt will cascade to a different connector based on specific error response codes from that connector

##Format
As with all configs, the standard wrapper is used.
```json5
{
  "kind": "PolicyCascade",                // Must be set to "PolicyCascade"
  "metadata": {
    "projectId": "test-project",          // Must be set to the ChargeHive Project ID you were issued with
    "name": "test-cascade-policy",        // Set this to a memorable name for the cascade policy, no spaces, all lowercase
  },
  "specVersion": "v1",                    // Must be set to the correct version
  "selector": {},                         // May be used to apply this to a subset of charges
  "spec": {
    "rules": [                            // list of the cascade policy rules
      {
        "library": "sandbox",             // Name of the library that this cascade rule is applied to (see options below)
        "originalResponseCode": "100",    // Raw error code returned by the library to be matched
        "cascade": false                  // Determines if this rule results in a cascade or not
      }
    ]
  }
}
```

###Spec Definition
FieldName | Required | Definition 
---:|---|:---
rules | true | Contains a non-empty list of cascade rules.

###Rules Definition
FieldName | Required | Definition 
---:|---|:---
library | true | string value representing the library (see below)
originalResponseCode | true | string representation of the response code from the connector
cascade | true | boolean flag, true will cascade to another connector

###Available Libraries:
Library Name | Value to use 
---:|:---
Sandbox | "sandbox"
Authorize | "authorize"
Braintree | "braintree"
QualPay | "qualpay"
Stripe | "stripe"
PaySafe | "paysafe"
PaySafeApplePay | "paysafe-applepay"
PaySafeGooglePay | "paysafe-googlepay"
Worldpay | "worldpay"
PayPalWebsitePaymentsPro | "paypal-websitepaymentspro"
PayPalExpressCheckout | "paypal-expresscheckout"
Vindicia | "vindicia"
ChargeHive | "chargehive"
MaxMind | "maxmind"
CyberSource | "cybersource"


##Full Example
```json5
{
  "kind": "PolicyCascade",
  "metadata": {
    "projectId": "test-project",
    "name": "test-cascade-policy",
  },
  "specVersion": "v1",
  "selector": {},
  "spec": {
    "rules": [
      {
        "library": "sandbox",
        "originalResponseCode": "100",
        "cascade": false
      }
    ]
  }
}
```