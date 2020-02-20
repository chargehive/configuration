# ChargeHive Configuration
Charge Hive is controlled though configuration files which are applied using the `chive` tool. 
Configurations can be saved, updated, loaded and deleted using the tool.

## Configuration File Structure
All configurations used in ChargeHive follow the same wrapper pattern:
```json5
{
  "kind": "KindOfConfig",      // [Required] Must be set to the kind of config as detailed in each config section below
  "metaData": {                // Contains general info about this config
    "projectId": "",           // [Required] Use the projectId you have been issued with from ChargeHive
    "name": "",                // [Required] Unique name of this configuration (string, no spaces, all lowercase)
    "displayName": "",         // Display name which will be shown in the front end
    "description": "",         // Long description of the config item
    "annotations": {           // Key value pairs for additional processing
      "key": "value"
    },
    "labels": {                // Key value pairs for front end grouping
      "key": "value"
    }
  },
  "specVersion": "v1",         // [Required] Must specify an API version
  "selector": {                // See the section below on selectors
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
  "spec": {}                   // Configuration details specific to the kind of config
}
```

## Configuration Selectors
Selectors allow the config to only be applied to a subset of charges based on whether they match the set criteria.
For more information see the [Selectors](selectors.md) section

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

### Installing `chive` - the configuration tool

You should have received an access token from ChargeHive, then you'll need to download the tool:
+ [Windows](https://cdn.chargehive.com/tools/chive/dist/windows/chive.exe)
+ [Mac](https://cdn.chargehive.com/tools/chive/dist/mac/chive)

To use the `chive` tool, you can either call it directly with your credentials:
```
chive --project-id="your-project-id" --access-token="your-access-token"
```

Or you can create a credentials file `.chive.yaml` with the following structure:
```yaml
projectID: your-project-id
accessToken: your-access-token
```

You can use the `.chive.yaml` config by any of these methods:
- storing it in the same directory as the `chive` executable
- storing it in the `$HOME` directory
- set where `chive` looks for the file: `chive --config-file="/path/to/your/config/.chive.yaml` and put it there

Once this is done, you should run `chive health` to check that your credentials are correct.
If you receive an error, please get in contact for support.

Options for the `chive` command are as follows:
```
❯ chive --help
chive allows communication with the Charge Hive API via CLI

Usage:
  chive [command]

Available Commands:
  apply       Apply a configuration file
  backup      Backup current project configurations
  delete      Delete a specific configuration
  get         Get a configuration
  health      Check your project and access token details are correct
  help        Help about any command
  list        List stored configurations
  verify      Verify a configuration

Flags:
      --access-token string   Access Token
      --api-host string       API Host
      --config-file string    configuration file (default is $HOME/.chive.yaml)
  -h, --help                  help for chive
      --project-id string     project ID

Use "chive [command] --help" for more information about a command.
```


### Creating a new configuration
Once you're happy with the configuration you have created, run the following command to apply it:
```
chive apply ./path/to/config.json
```

### Retrieving an existing configuration



### Updating an existing configuration
### Deleting an existing configuration


