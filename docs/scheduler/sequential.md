# Sequential Scheduler

The sequential scheduler allows for a sequence or schedule of connectors to be used when attempting to complete a charge.
They are mapped by attempt number.

## Format

As with all configs, the standard wrapper is used.

```json5
{
  "kind": "SchedulerSequential",
  // Must be set to "SchedulerSequential"
  "metadata": {
    "projectId": "test-project",
    // Must be set to the ChargeHive Project ID you were issued with
    "name": "test-scheduler"
    // Set this to a memorable name for the method lock policy, no spaces, all lowercase
  },
  "specVersion": "v1",
  // Must be set to the correct version
  "selector": {},
  // May be used to apply this to a subset of charges
  "spec": {
    "schedules": {
      // Schedule object with the sequence of operation
      "0": {
        "attemptConfig": {
          // Configuration used when processing this schedule
          "poolType": "single",
          // The order that this attempt should iterate connectors
          "methodSelector": "primary",
          // How payment method should be selected for this attempt
          "connectorLimit": 0,
          // Maximum number of connectors to process within an attempt per method
          "methodLimit": 0,
          // Maximum number of methods to be attempt per method
          "cascadeDelay": null,
          // Delay between connector cascades
          "overridePoolConnectorIDs": [
            // Overrides the pool, and selects these connectors
            "test-connector"
          ]
        },
        "timeDelay": 86400000000000,
        // Delay in nanoseconds before the next Attempt - if set to 0, ChargeHive will still delay 30 minutes to avoid multiple attempts at the same time
        "timeDelayOrigin": "initialisation",
        // Defines when the Delay before the next Attempt is based from - either "initialisation" or "last-failure"
        "timeDelaySync": "Closest",
        // Specifies when the transaction should be performed relative to the schedules TimeSync
        "timeWindowHours": "2",
        // Specifies the available duration for the transaction to be queued within
        "timeSyncHour": 2,
        // An hour designation
        "timeSyncZone": "UTC",
        // UTC or ULT
        "dayOfMonth": 1,
        // Day of the month to run, e.g. 1st
        "dayOfWeek": 2
        // Day of the week to run, e.g. Tuesday (Monday = 1, Sunday = 7)
      },
    }
  }
}
```

### Schedule Definition

FieldName | Required | Definition 
---:|----------|:---
[attemptConfig](#attemptconfig-definition)| true     |Configuration to use when processing this schedule
timeDelay| true     |Amount of time to wait before processing after TimeDelayOrigin in **nanoseconds**
[timeDelayOrigin](#timedelayorigin-values)| true     |Specifies when the time origin is based from
[timeDelaySync](#timedelaysync-values)| true     |Specifies when the transaction should be performed relative to the schedules TimeSync
timeWindowHours| false    |Specifies the available duration for the transaction to be queued within
timeSyncHour| true     |An hour designation (0-23) i.e 2 == 2AM. Ignored if TimeDelaySync is set to None
[timeSyncZone](#timesynczone-values)| true     |Indicates the timezone that the TimeSyncHour is relative to. Ignored if TimeDelaySync is set to None
dayOfMonth| false    |Day of the month to run (1-31)
dayOfWeek| false    |Day of the week to run (1-7) (Monday = 1, Sunday = 7)

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

Value | Definition
-----------:|:--------------------------
"None" |will ignore the TimeSyncHour value
"Earliest" |will run the transaction at the earliest sync hour relative to TimeSync
"Latest" |will run the transaction at the latest sync hour relative to TimeSync
"Closest" |will run the transaction at the closest sync hour relative to TimeSync

### TimeSyncZone Values

 Value | Definition                 
------:|:---------------------------
 "ULT" | Users Local Time           
 "UTC" | Universal Time Coordinated 

### Day Of Month

 Value | Definition              
------:|:------------------------
  1-28 | Day of the month to run 

### Day Of Week

 Value | Definition 
------:|:-----------
1 | Monday     
2 | Tuesday    
3 | Wednesday  
4 | Thursday   
5 | Friday     
6 | Saturday   
7 | Sunday     

## Full Example

```json
{
  "kind": "SchedulerSequential",
  "metadata": {
    "projectId": "test-project",
    "name": "test-scheduler"
  },
  "spec": {
    "schedules": {
      "0": {
        "attemptConfig": {
          "poolType": "cascade",
          "methodSelector": "primary",
          "connectorLimit": 0,
          "methodLimit": 0,
          "cascadeDelay": null,
          "overridePoolConnectorIDs": [
            "sandbox-connector"
          ]
        }
      },
      "1": {
        "attemptConfig": {
          "poolType": "cascade",
          "methodSelector": "primary",
          "connectorLimit": 0,
          "methodLimit": 0,
          "cascadeDelay": null,
          "overridePoolConnectorIDs": [
            "sandbox-connector"
          ]
        }
      },
      "2": {
        "attemptConfig": {
          "poolType": "cascade",
          "methodSelector": "primary",
          "connectorLimit": 0,
          "methodLimit": 0,
          "cascadeDelay": null,
          "overridePoolConnectorIDs": [
            "sandbox-connector"
          ]
        }
      },
      "3": {
        "attemptConfig": {
          "poolType": "cascade",
          "methodSelector": "primary",
          "connectorLimit": 0,
          "methodLimit": 0,
          "cascadeDelay": null,
          "overridePoolConnectorIDs": [
            "sandbox-connector"
          ]
        }
      }
    }
  }
}
```