# On-Demand Scheduler

## Format
As with all configs, the standard wrapper is used.

```json5
{
  "kind": "SchedulerOnDemand",              // Must be set to "SchedulerOnDemand"
  "metadata": {
    "projectId": "test-project",            // Must be set to the ChargeHive Project ID you were issued with
    "name": "test-on-demand"                // Set this to a memorable name for the on demand scheduler policy, no spaces, all lowercase
  },
  "spec": {
    "schedule": {                           // Schedule object
      "attemptConfig": {},                  // Configuration used when processing this schedule
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
attemptConfig|true|Configuration to use when processing this schedule
timeDelay|true|Amount of time to wait before processing after TimeDelayOrigin in **nanoseconds**
[timeDelayOrigin](#timedelayorigin-values)|true|Specifies when the time origin is based from
[timeDelaySync](#timedelaysync-values)|true|Specifies when the transaction should be performed relative to the schedules TimeSync
timeSyncHour|true|An hour designation (1-24) i.e 2 == 2AM or where less than 1 indicates that this value is not set
[timeSyncZone](#timesynczone-values)|true|Indicates the timezone that the TimeSyncHour is relative to

### TimeDelayOrigin Values
Value| Definition
---:|:---
"initialisation"|Indicates that the time is based from the initialisation of the charge
"last-failure"|Indicates that the time is based from the last transaction failure

### TimeDelaySync Values
Value| Definition
---:|:---
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
  "kind": "SchedulerOnDemand",
  "metadata": {
    "projectId": "test-project",
    "name": "test-on-demand"
  },
  "spec": {
    "schedule": {
      "attemptConfig": {},
      "timeDelay": 86400000000000,
      "timeDelayOrigin": "initialisation",
      "timeDelaySync": "Closest",
      "timeSyncHour": 2,
      "timeSyncZone": "UTC"
    }
  }
}
```