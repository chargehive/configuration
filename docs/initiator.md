# Initiator


## Format
As with all configs, the standard wrapper is used.

```json5
{
  "kind": "Initiator",                // Must be set to "PolicyMethodUpgrade"
  "metadata": {
    "projectId": "test-project",      // Must be set to the ChargeHive Project ID you were issued with
    "name": "test-initiator"          // Set this to a memorable name for the method lock policy, no spaces, all lowercase
  },
  "specVersion": "v1",                // Must be set to the correct version
  "selector": {},                     // May be used to apply this to a subset of charges
  "spec": {
    "type": "auth",                   // Type of this initiator
    "initialConnector": "config",     // Indicates the method used to select a connector
    "attemptConfig": {
      "poolType": "single",           // The order that this attempt should iterate connectors
      "methodSelector": "primary",    // How payment method should be selected for this attempt
      "connectorLimit": 0,            // Maximum number of connectors to process within an attempt per method
      "methodLimit": 0,               // Maximum number of methods to be attempt per method
      "cascadeDelay": null,           // Delay between connector cascades
      "overridePoolConnectorIDs": [   // Overrides the pool, and selects these connectors
        "test-connector"
      ]
    }
  }
}
```

## Spec Definition
FieldName | Required | Definition 
---:|---|:---
type|true| "auth" (how an authorization request is handled), "renewal" (how a renewal request is handled) or "capture" (how a secondary new charge is handled)
[initialConnector](#initialconnector-values)|true|Indicates the method used to select a connector
[attemptConfig](#attemptconfig-definition)|true|Defines additional configuration options for the attempts

### InitialConnector Values
Value | Definition 
---:|:---
"none"|ConnectorSelectorNone indicates to use no connectors (this is the same as setting empty and is the default value)
"sticky-first"|ConnectorSelectorStickyFirst indicates the connector should stick to the first connector that was successful for that payment method
"sticky-last"|Indicates the connector should stick to the most recent connector that was successful for that payment method
"sticky-any"|Indicates to use any connector that has a past success for that payment method
"sticky-verified"|Indicates to use any connector that has a past success and has been verified for that payment method
"config"|Indicates to use only the connectors specified in the configuration

### AttemptConfig Definition
FieldName | Required | Definition 
---:|---|:---
[poolType](#pooltype-values)|true|The order that this attempt should iterate connectors
[methodSelector](#methodselector-values)|true|How payment method should be selected for this attempt
connectorLimit|true|Maximum number of connectors to process within an attempt per method
methodLimit|true|Maximum number of methods to be attempt per method
cascadeDelay|true|Duration to wait between each cascade in **nanoseconds** or null for instant
overridePoolConnectorIDs|false|will use this connectors instead of the ones in the pool

### PoolType Values
Value | Definition 
---:|:---
"single"|Provides a pool of a single connector
"failover"|Processes the pool items in order until retrieving a result
"cascade"|Iterate connectors according to cascade rules

### MethodSelector Values
Value | Definition 
---:|:---
"primary"|Indicates that the first available payment method should be used
"backup"| Indicates that the second available payment method should be used
"all"| Indicates that all methods can be used
"all-backup"| Indicates that anything available other than the 1st (primary) should be used

## Full Example

```json
{
  "kind": "Initiator",
  "metadata": {
    "projectId": "test-project",
    "name": "test-initiator"
  },
  "spec": {
    "type": "auth",
    "initialConnector": "config",
    "attemptConfig": {
      "PoolType": "single",
      "MethodSelector": "primary",
      "ConnectorLimit": 0,
      "MethodLimit": 0,
      "CascadeDelay": null,
      "OverridePoolConnectorIDs": [
        "worldpay-connector"
      ]
    }
  }
}
```