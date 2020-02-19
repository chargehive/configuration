# Intro
Charge Hive is controlled though configuration files

## Configuration Files
All configs used in ChargeHive follow the same wrapper pattern:

```json
{
  "kind": "KindOfConfig",
  "metaData": {
    "projectId": "",
    "name": "",
    "uuid": "",
    "displayName": "",
    "description": "",
    "annotations": {
      "key": "value"
    },
    "labels": {
      "key": "value"
    }
  },
  "specVersion": "",
  "selector": {
    "priority": 0,
    "expressions": [
      {
        "key": "NameOfTheKey",
        "operator": "Equal",
        "conversion": "",
        "values": [
          "value1",
          "value2"
        ]
      }
    ]
  },
  "spec": {}
}
```


## Configuration Types
#### Connectors
+ [Connector](connectors/connector.md) is an external api services like payment gateways or fraud services
+ [ConnectorPool](connectors/pool.md) is a pool of connectors ... 

#### Integration
+ [Slack](integration/slack.md) is an integration with the slack messaging service for service/event notifications

#### Policies
+ [Cascade Policy](policy/cascade.md) defines how a failed charge retries with a different connector
+ [Charge Expiry Policy](policy/charge_expiry.md) defines how long a charge should last before expiring
+ [Fraud Policy](policy/fraud.md) defines how and when a charge is check for fraudulent characteristics 
+ [Method Lock Policy](policy/method_lock.md) defines how long to block a payment method in the event of a decline or failure
+ [Method Upgrade Policy](policy/method_upgrade.md) defines what modifications can be made to a payment method in order to complete a transaction 
+ [Method Verify Policy](policy/method_verify.md) defines how and when a card should be verified
+ [SCA (Secure Customer Authentication) Policy](policy/sca.md) defines PSD2 SCA policy for transactions

#### Schedulers
+ [Initiator Scheduler](scheduler/initiator.md) defines the first scheduler on a charge
+ [On Demand Scheduler](scheduler/on_demand.md) defines schedules for arbitrary billing requests like golden days
+ [Refund Scheduler](scheduler/refund.md) defines refund retries at max attempts
+ [Sequential Scheduler](scheduler/sequential.md) defines which connectors to attempt, in what order when processing a charge


## Handling Configurations

#### Using Chive - the configuration tool
#### Creating a new configuration
#### Retrieving an existing configuration
#### Updating an existing configuration
#### Deleting an existing configuration


