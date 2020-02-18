# Method Upgrade Policy
Method Upgrade Policy is used to temporarily alter the existing payment method information.

## Format
As with all configs, the standard wrapper is used.

```json5
{
  "Kind": "PolicyMethodUpgrade",    // Must be set to "PolicyMethodUpgrade"
  "metadata": {
    "projectId": "test-project",    // Must be set to the ChargeHive Project ID you were issued with
    "name": "test-method-lock"      // Set this to a memorable name for the method lock policy, no spaces, all lowercase
  },
  "specVersion": "v1",              // Must be set to the correct version
  "selector": {},                   // May be used to apply this to a subset of charges
  "spec": {
    "extendExpiry": true            // Enable or disable the policy
  }
}

```
### Definition
FieldName | Required | Definition 
---:|---|:---
extendExpiry|true| ExtendExpiry date on payment methods to the next likely expiry date


## Full Example

In this example, we use the selector to only apply extendExpiry option to payments which have a response error type of `RESPONSE_ERROR_AVAILABLE_FUNDS`
See the section on selectors for more information

```json
{
  "Kind": "PolicyMethodUpgrade",
  "metadata": {
    "projectId": "test-project",
    "name": "test-method-upgrade"
  },
  "specVersion": "v1",
  "selector": {
    "priority": 50,
    "expressions": [
      {
        "key": "transaction.response.error.type",
        "operator": "Equal",
        "values": [
          "RESPONSE_ERROR_AVAILABLE_FUNDS"
        ]
      }
    ]
  },
  "spec": {
      "extendExpiry": true
  }
}
```