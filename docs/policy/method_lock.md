# Method Lock Policy
Method Lock Policy is used to lock a payment method preventing it being used for payment for a period of time

## Format
As with all configs, the standard wrapper is used.

```json5
{
  "kind": "PolicyMethodLock",    // Must be set to "PolicyMethodLock"
  "metadata": {
    "projectId": "test-project", // Must be set to the ChargeHive Project ID you were issued with
    "name": "test-method-lock"   // Set this to a memorable name for the method lock policy, no spaces, all lowercase
  },
  "specVersion": "v1",           // Must be set to the correct version
  "selector": {},                // May be used to apply this to a subset of charges
  "spec": {
    "duration": 300,             // Duration to lock the payment method
    "reason": "Description"      // Definition of why the payment was locked
  }
}

```
## Definition
FieldName | Required | Definition 
---:|---|:---
duration|true|Duration is the duration of time (in seconds) that a payment method should be locked for on application of this policy
reason|false|Description of why the payment method was locked at this time.

## Full Example

In this example, we use the selector to only apply the 5min lock to payments which have a response error type of `RESPONSE_ERROR_AVAILABLE_FUNDS`
See the section on selectors for more information

```json
{
  "kind": "PolicyMethodLock",
  "metadata": {
    "projectId": "test-project",
    "name": "test-method-lock"
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
    "duration": 300,
    "reason": "No available funds on last attempt"
  }
}
```