# Refund Scheduler
The refund scheduler defines the time delay between refund attempts

## Format
As with all configs, the standard wrapper is used.

```json5
{
  "kind": "SchedulerRefund",          // Must be set to "SchedulerRefund"
  "metadata": {
    "projectId": "test-project",      // Must be set to the ChargeHive Project ID you were issued with
    "name": "refund-schedule"         // Set this to a memorable name for the method lock policy, no spaces, all lowercase
  },
  "specVersion": "v1",                // Must be set to the correct version
  "selector": {},                     // May be used to apply this to a subset of charges
  "spec": {
    "Schedules": {                    // Refund schedules indexed by the attempt
      "0": {                          // First schedule in the sequence
        "TimeDelay": 60000000000      // Time delay in nanoseconds (1min)
      },
      "1": {                          // Second schedule in the sequence
        "TimeDelay": 60000000000      // Time delay in nanoseconds (1min)
      },
      "2": {                          // Third schedule in the sequence
        "TimeDelay": 120000000000      // Time delay in nanoseconds (2min)
      }
    }
  }
}
```

## Full Example

```json
{
  "kind": "SchedulerRefund",
  "metadata": {
    "projectId": "test-project",
    "name": "refund-schedule"
  },
  "specVersion": "v1",
  "selector": {},
  "spec": {
    "Schedules": {
      "0": {
        "TimeDelay": 60000000000
      },
      "1": {
        "TimeDelay": 60000000000
      },
      "2": {
        "TimeDelay": 120000000000
      }
    }
  }
}
```