# Fraud Policy
The fraud policy defines

## Format
As with all configs, the standard wrapper is used.

```json5
{
  "kind": "PolicyFraud",            // Must be set to "PolicyFraud"
  "metadata": {
    "projectId": "test-project",    // Must be set to the ChargeHive Project ID you were issued with
    "name": "preauth-fraud"         // Set this to a memorable name for the fraud policy, no spaces, all lowercase
  },
  "specVersion": "v1",              // Must be set to the correct version
  "selector": {},                   // May be used to apply this to a subset of charges
  "spec": {
    "connectorIDs": [               // List of fraud connector IDs to apply the policy to
      "fraud-chargehive",
    ],
    "transactionType": "auth",      // The payment step the policy runs at
    "position": "pre",              // When the policy runs relative to that step
    "attempts": "every",            // Which attempts at that step the policy runs on
    "cit": "scan",                  // What to do with a customer-initiated payment
    "mit": "inform",                // What to do with a merchant-initiated payment
    "checkType": "all",             // Sets how it works through the connectors
    "declineOn": ["deny"]           // Scan results that decline the payment, before it
  }
}
```

### Spec Definition
FieldName | Required | Definition 
---:|---|:---
connectorIDs|true|List of fraud connector IDs to apply the policy to. These must be **fraud** connectors
[transactionType](#fraudtransactiontype)|true, except for `ondemand`| The payment step the policy runs at
[position](#fraudposition)|true| When the policy runs relative to its payment step
[attempts](#fraudattempts)|false| Which of the charge's attempts at the payment step the policy runs on. Defaults to `every`
[cit](#fraudmode)|false| What the policy does with a customer-initiated payment. Defaults to `none`
[mit](#fraudmode)|false| What the policy does with a merchant-initiated payment, such as a renewal. Defaults to `none`
[checkType](#fraudchecktype)|true| Sets how it works through the connectors
declineOn|false| Scan results that decline the payment. Applies to a policy that scans before its payment step
[checkTime](#fraudchecktime)|false| Deprecated: sets when to fraud check, in place of the fields above

### FraudTransactionType
value | Definition 
---:|:---
"auth" | An authorisation
"capture-auth" | The capture of an earlier authorisation
"capture" | A payment captured without an earlier authorisation

### FraudPosition
value | Definition 
---:|:---
"pre" | Before the payment step. A scan's result can decline the payment through `declineOn`
"post" | After the payment step, approved or declined
"post-success" | After the payment step is approved
"post-decline" | After the payment step is declined
"ondemand" | When a fraud scan is requested for the charge. Has no `transactionType`

A scan after the payment step still returns its result to the merchant, who can act on it, for example by voiding an auth or refunding a capture.

### FraudAttempts
value | Definition 
---:|:---
"first" | The charge's first attempt only
"every" | Every attempt

### FraudMode
value | Definition 
---:|:---
"scan" | Asks the fraud connectors for a result. Some providers, such as Kount, charge for this
"inform" | Sends the payment to the fraud connectors for their records, without asking for a result. A connector that can only scan is skipped
"none" | Leaves the payment out

### FraudCheckType
value | Definition 
---:|:---
"all"|Perform a fraud check on all provided connectors
"failover"|Will perform a fraud check on one connector ID at a time, stopping at the first success

### FraudCheckTime
Deprecated. Each value runs as the fields above, scanning both customer- and merchant-initiated payments.

value | Runs as
---:|:---
"preauth-first" | `auth`, `pre`, `first`
"preauth-every" | `auth`, `pre`, `every`
"auth-success" | `auth`, `post-success`, `every`
"ondemand" | `ondemand`

## Full Example

```json
{
  "kind": "PolicyFraud",
  "metadata": {
    "projectId": "test-project",
    "name": "renewal-fraud"
  },
  "specVersion": "v1",
  "selector": {},
  "spec": {
    "connectorIDs": [
      "fraud-default"
    ],
    "transactionType": "capture",
    "position": "post",
    "mit": "inform",
    "checkType": "all"
  }
}
```
