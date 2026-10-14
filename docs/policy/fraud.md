# Fraud Policy
The fraud policy defines

## Format
As with all configs, the standard wrapper is used.

```json5
{
  "kind": "PolicyFraud",            // Must be set to "PolicyFraud"
  "metadata": {
    "projectId": "test-project",    // Must be set to the ChargeHive Project ID you were issued with
    "name": "ondemand-fraud"        // Set this to a memorable name for the fraud policy, no spaces, all lowercase
  },
  "specVersion": "v1",              // Must be set to the correct version
  "selector": {},                   // May be used to apply this to a subset of charges
  "spec": {
    "connectorIDs": [               // List of fraud connector IDs to apply the policy to
      "fraud-chargehive",
    ],
    "checkTimes": ["ondemand"],     // Sets when to fraud check
    "checkType": "all"              // Sets how it works through the connectors
  }
}
```

### Spec Definition
FieldName | Required | Definition 
---:|---|:---
connectorIDs|true|List of fraud connector IDs to apply the policy to. These must be **fraud** connectors
[checkTimes](#fraudchecktime)|true| Sets when to fraud check. A policy runs either before the auth (`preauth-first` or `preauth-every`, not both) or after the payment (any of `auth-success`, `auth-decline`, `capture-success`, `capture-decline`), not both; `ondemand` cannot be used with any other time
checkTime|false| Deprecated: a single check time, used in place of checkTimes. Accepts `preauth-first`, `preauth-every`, `auth-success` and `ondemand`
[checkType](#fraudchecktype)|true| Sets how it works through the connectors

### FraudCheckType
value | Definition 
---:|:---
"all"|Perform a fraud check on all provided connectors
"failover"|Will perform a fraud check on one connector ID at a time, stopping at the first success

### FraudCheckTime
value | Definition 
---:|:---
"preauth-first" | Indicates that a fraud scan should check before the first auth
"preauth-every" | Indicates that a fraud scan should check before every auth
"auth-success" | Indicates that a fraud scan should run after a successful auth
"auth-decline" | Indicates that a fraud scan should run after a declined auth
"capture-success" | Indicates that a fraud scan should run after a successful capture, which is a sale or the capture of an earlier auth
"capture-decline" | Indicates that a fraud scan should run after a declined capture
"ondemand" | Indicates that a fraud scan should run on demand only

## Full Example

```json
{
  "kind": "PolicyFraud",
  "metadata": {
    "projectId": "test-project",
    "name": "default-fraud"
  },
  "specVersion": "v1",
  "selector": {},
  "spec": {
    "ConnectorIDs": [
      "fraud-default"
    ],
    "CheckTimes": ["preauth-first"],
    "CheckType": "failover"
  }
}
```