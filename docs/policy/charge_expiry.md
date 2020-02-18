# Charge Expiry Policy
This policy defines how long a charge lasts until it expires

##Format
As with all configs, the standard wrapper is used.

```json5
{
  "kind": "PolicyChargeExpiry",               // Must be set to "PolicyChargeExpiry"
  "metadata": {
    "projectId": "test-project",              // Must be set to the ChargeHive Project ID you were issued with
    "name": "test-cascade-policy",            // Set this to a memorable name for the cascade policy, no spaces, all lowercase
  },
  "spec": {
    "timeout": 604800000000000,               // Length of time before charge expires 
    "attempts": 10                            // Number of attempts before the charge expires
  }
}
```
###Spec Definition
FieldName | Required | Definition 
---:|---|:---
timeout | false | Int64 Number of nanoseconds before the charge expires. The representation limits the largest representable duration to approximately 290 years
attempts | false | Int64 value representing the number of times to attempt a charge before expiring

##Full Example

Example which limits charges to a maximum of 10 attempts or 1 week, whichever comes first.
```json
{
  "kind": "PolicyChargeExpiry",
  "metadata": {
    "projectId": "test-project",
    "name": "test-cascade-policy",
  },
  "spec": {
    "timeout": 604800000000000, 
    "attempts": 10
  }
}
```