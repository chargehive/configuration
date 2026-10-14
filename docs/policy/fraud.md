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
    "checkTime": "preauth-every",   // Sets when to fraud check
    "merchantCheck": true,          // Checks merchant-initiated payments after their outcome instead
    "checkType": "all"              // Sets how it works through the connectors
  }
}
```

### Spec Definition
FieldName | Required | Definition 
---:|---|:---
connectorIDs|true|List of fraud connector IDs to apply the policy to. These must be **fraud** connectors
[checkTime](#fraudchecktime)|true| Sets when to fraud check
merchantCheck|false| Checks a merchant-initiated payment once, after its outcome (a sale's, or an auth's when it is authed before it is captured), in place of at checkTime. Such a payment has no customer present, so it is reported rather than decided. A sale counts as merchant-initiated whoever initiated it, since it runs asynchronously. Cannot be used with `ondemand`
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
"postauth" | Indicates that a fraud scan should run after an auth's outcome, approved or declined, while an approved auth can still be voided before it is captured
"postcapture" | Indicates that a fraud scan should run after a capture's outcome, collected or declined. A capture is a sale or the capture of an earlier auth
"auth-success" | Deprecated: runs as `postauth`
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