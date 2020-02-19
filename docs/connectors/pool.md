# Connector Pool
The Slack integration config defines the slack account and settings in order to get updates from ChargeHive to be sent to slack.

## Format
As with all configs, the standard wrapper is used.

```json5
{
  "kind": "ConnectorPool",                    // Must be set to "ConnectorPool"
  "metadata": {
    "projectId": "test-project",              // Must be set to the ChargeHive Project ID you were issued with
    "name": "test-pool-x",                    // Set this to a memorable name for the slack integration, no spaces, all lowercase
  },
  "specVersion": "v1",                        // Must be set to the correct version
  "selector": {},                             // May be used to apply this to a subset of charges
  "spec": {
    "restriction": "unrestricted",            // Can be set to: "unrestricted", "noRepeat" or "lowestUsage"
    "connectors": [                           // Connectors contains the list of connectors in the pool
      {
        "connectorId": "xxxxxxxxxxxx",        // Identifier for a connector
        "priority": 10,                       // Highest priority item has the lowest integer value (0 is the highest priority)
        "weighting": 0,                       // Weighting is used to weigh items of the same priority, secondary to priority (0-1000)
        "uses": 10,                           // Uses is the maximum times a connector can be used in a single charge
      }
    ]
  }
}
```
## Spec Definition
FieldName | Required | Definition 
---:|---|:---
restriction | false | "unrestricted" (Default) , "noRepeat" or "lowestUsage"
[connectors](#connector-definition) | false | Non-empty list of the connectors in the pool

### Connector definition
FieldName | Required | Definition 
---:|---|:---
connectorId | true | Identifier for a connector, must match a previously defined connector id
priority | false | Highest priority item has the lowest integer value (0 is the highest priority)
weighting | false | Weighting is used to weigh items of the same priority, secondary to priority (0-1000)
uses | false | Uses is the maximum times a connector can be used in a single charge

## Full Example
Here is a working example to create a connector pool called `test-pool-x` which contains two connectors.

```json
{
  "kind": "ConnectorPool",                    
  "metadata": {
    "projectId": "test-project",              
    "name": "test-pool-x"                 
  },
  "specVersion": "v1",
  "selector": {},
  "spec": {
    "restriction": "unrestricted",            
    "connectors": [                           
      {
        "connectorId": "connector1",        
        "priority": 0,                       
        "weighting": 0,                       
        "uses": 10                           
      },
      {
        "connectorId": "connector2",        
        "priority": 10,                       
        "weighting": 0,                       
        "uses": 10                           
      }
    ]
  }
}
```