# Method Verify Policy
Method Verify policy contains the settings used to perform a payment method verification. 
This might be done using the payment gateways verify method, or by authorizing a small payment and voiding that authorisation once completed

## Format
As with all configs, the standard wrapper is used.

```json5
{
  "kind": "PolicyMethodUpgrade",         // Must be set to "PolicyMethodUpgrade"
  "metadata": {
    "projectId": "test-project",         // Must be set to the ChargeHive Project ID you were issued with
    "name": "test-method-lock"           // Set this to a memorable name for the method lock policy, no spaces, all lowercase
  },
  "specVersion": "v1",                   // Must be set to the correct version
  "selector": {},                        // May be used to apply this to a subset of charges
  "spec": {
    "verifyMethodOnTokenization": true,  // Verify at the same time as verify
    "amount": 1999,                      // Amount to auth and void if verify not available
    "amountCurrency": "USD",             // Currency to auth with
    "connectorID": "conn1"               // ID of the connector to validate against
  }
}

```
## Spec Definition
FieldName | Required | Definition 
---:|---|:---
verifyMethodOnTokenization|false|If true the payment method will be verified at the same time it is tokenized
amount|true|Amount is a monetary value integer that will be authorized on a card to verify its ability to make payments this should be an amount in the currencies smallest denomination i.e a value of 44 would equate to 0.44 GBP
amountCurrency|true|This is the currency code for the specified amount i.e GBP
connectorId|true|This is the ID of the connector that is used to verify payment methods

## Full Example

```json
{
  "kind": "PolicyMethodUpgrade",
  "metadata": {
    "projectId": "test-project",
    "name": "test-method-lock"
  },
  "specVersion": "v1",
  "selector": {},
  "spec": {
    "verifyMethodOnTokenization": true,
    "amount": 1999,
    "amountCurrency": "USD",
    "connectorId": "conn1"
  }
}
```