# SCA (Secure Customer Authentication) Policy
The cascade policy defines the behaviour of how failed attempts are retried through different connectors

## Format
As with all configs, the standard wrapper is used.

```json5
{
  "kind": "PolicySCA",                  // Must be set to "PolicyMethodUpgrade"
  "metadata": {
    "projectId": "test-project",        // Must be set to the ChargeHive Project ID you were issued with
    "name": "sca-policy"                // Set this to a memorable name for the method lock policy, no spaces, all lowercase
  },
  "specVersion": "v1",                  // Must be set to the correct version
  "selector": {},                       // May be used to apply this to a subset of charges
  "spec": {
    "shouldIdentify": true,             // Enable or disable identification step
    "shouldChallengeOptional": true,    // If challenge is optional from issuer, use this value to decide
    "shouldByPassChallenge": "",        // If challenge is required, override with this value (see below for options)
    "shouldAuthOnError": true,          // If the gateway authentication response is an error, authorize payment anyway
    "shouldAuthOnN": true,              // If the gateway authentication response is an N, authorize payment anyway
    "shouldAuthOnR": true               // If the gateway authentication response is an R, authorize payment anyway
  }
}
```
## Definition
FieldName | Required | Default | Definition 
---:|---|---|:---
shouldIdentify |false| false | Indicates if the identification stages should take place
shouldChallengeOptional|false| false | Challenge based on an optional response from the connector (setting this to false will not display the challenge)
[shouldByPassChallenge](#shouldbypasschallenge-definition)|false| false | If the challenge is required, bypassing this will attempt an auth without displaying the challenge
shouldChallenge3dSecureV1|false| false | Determines if the connector can fallback to 3DS v1 when 3DS v2 is not available
shouldAuthOnError|false| false | If true and an error response is returned from the connector; proceed to auth anyway


### ShouldByPassChallenge Definition
Value | Description 
---:|:---
"" | Indicates to not bypass a required challenge
"cascade" | Indicates to auth on the next connector
"current" | Indicate to bypass, but stay on the current connector (attempt auth anyway)

## Full Example

```json
{
  "kind": "PolicySCA",
  "metadata": {
    "projectId": "test-project",
    "name": "sca-policy"
  },
  "specVersion": "v1",
  "selector": {},
  "spec": {
    "ShouldIdentify": true,
    "ShouldChallengeOptional": true,
    "ShouldByPassChallenge": "",
    "ShouldAuthOnError": true
  }
}
```