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
    "checkTime": "ondemand",        // Sets when to fraud check
    "checkType": "all"              // Sets how it works through the connectors
  }
}
```

### Spec Definition
FieldName | Required | Definition 
---:|---|:---
connectorIDs|true|List of fraud connector IDs to apply the policy to. These must be **fraud** connectors
[checkTime](#fraudchecktime)|true| Sets when to fraud check
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
    "CheckTime": "preauth-first",
    "CheckType": "failover"
  }
}
```