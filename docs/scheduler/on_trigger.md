# On-Trigger Scheduler

## Format
As with all configs, the standard wrapper is used.

```json5
{
  "kind": "SchedulerOnTrigger",            // Must be set to "SchedulerOnTrigger"
  "metadata": {
    "projectId": "test-project",            // Must be set to the ChargeHive Project ID you were issued with
    "name": "test-on-trigger"                // Set this to a memorable name for the on trigger scheduler policy, no spaces, all lowercase
  },
  "specVersion": "v1",                      // Must be set to the correct version
  "selector": {},                           // May be used to apply this to a subset of charges
  "spec": {
    "schedule": {                           // Schedule object
      "attemptConfig": {                    // Configuration used when processing this schedule
            "poolType": "single",           // The order that this attempt should iterate connectors
            "methodSelector": "primary",    // How payment methods should be selected for this attempt
            "connectorLimit": 0,            // Maximum number of connectors to process within an attempt per method
            "methodLimit": 0,               // Maximum number of methods to be attempt per method
            "cascadeDelay": null,           // Delay between connector cascades
            "overridePoolConnectorIDs": [   // Overrides the pool, and selects these connectors
              "test-connector"
            ]
      },
      "timeDelay": 86400000000000,          // Delay in nanoseconds
      "timeDelayOrigin": "initialisation",  // Defines when a given time is based from
      "timeDelaySync": "Closest",           // Specifies when the transaction should be performed relative to the schedules TimeSync
      "timeSyncHour": 2,                    // An hour designation
      "timeSyncZone": "UTC"                 // UTC or ULT
    }
  }
}
```

### Spec Definition
FieldName | Required | Definition
---:|---|:---
[attemptConfig](#attemptconfig-definition)|true|Configuration to use when processing this schedule
timeDelay|true|Amount of time to wait before processing after TimeDelayOrigin in **nanoseconds**
[timeDelayOrigin](#timedelayorigin-values)|true|Specifies when the time origin is based from
[timeDelaySync](#timedelaysync-values)|true|Specifies when the transaction should be performed relative to the schedules TimeSync
timeSyncHour|true|An hour designation (0-23) i.e 2 == 2AM. Ignored if TimeDelaySync is set to None
[timeSyncZone](#timesynczone-values)|true|Indicates the timezone that the TimeSyncHour is relative to. Ignored if TimeDelaySync is set to None

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

### TimeDelayOrigin Values
Value| Definition
---:|:---
"initialisation"|Indicates that the time is based from the initialisation of the charge
"last-failure"|Indicates that the time is based from the last transaction failure

### TimeDelaySync Values
Value| Definition
---:|:---
"None"|will ignore the TimeSyncHour value
"Earliest"|will run the transaction at the earliest sync hour relative to TimeSync
"Latest"|will run the transaction at the latest sync hour relative to TimeSync
"Closest"|will run the transaction at the closest sync hour relative to TimeSync

### TimeSyncZone Values
Value| Definition
---:|:---
"ULT"|Users Local Time
"UTC"|Universal Time Coordinated

## Full Example
```json
{
  "kind": "SchedulerOnTrigger",
  "metadata": {
    "projectId": "test-project",
    "name": "test-on-trigger"
  },
  "spec": {
    "schedule": {
      "attemptConfig": {
            "poolType": "single",
            "methodSelector": "primary",
            "connectorLimit": 0,
            "methodLimit": 0,
            "cascadeDelay": null,
            "overridePoolConnectorIDs": [
              "test-connector"
            ]
      },
      "timeDelay": 86400000000000,
      "timeDelayOrigin": "initialisation",
      "timeDelaySync": "Closest",
      "timeSyncHour": 2,
      "timeSyncZone": "UTC"
    }
  }
}
```
