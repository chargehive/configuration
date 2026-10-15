# Fraud Policy
A fraud policy sets which fraud connectors check a charge's payments, at which transaction, and whether each payment is scanned for a result or only sent to the connectors for their records.

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
    "transactionType": "auth",      // The transaction the policy runs at
    "position": "pre",              // When the policy runs, relative to the transaction
    "attempts": "every",            // Which of the charge's attempts at the transaction the policy runs on
    "cit": "scan",                  // What the policy does with a customer-initiated payment
    "mit": "inform",                // What the policy does with a merchant-initiated payment
    "checkType": "all",             // Sets how it works through the connectors
    "declineOn": ["deny"]           // Scan results that decline the transaction
  }
}
```

### Spec Definition
FieldName | Required | Definition 
---:|---|:---
connectorIDs|true|List of fraud connector IDs to apply the policy to. These must be **fraud** connectors
[transactionType](#fraudtransactiontype)|true| The transaction the policy runs at. Leave it out when `position` is `on-demand`
[position](#fraudposition)|true| When the policy runs, relative to the transaction
[attempts](#fraudattempts)|false| Which of the charge's attempts at the transaction the policy runs on. Defaults to `every`
[cit](#fraudmode)|false| What the policy does with a customer-initiated payment. Defaults to `none`
[mit](#fraudmode)|false| What the policy does with a merchant-initiated payment, such as a renewal. Defaults to `none`
[checkType](#fraudchecktype)|true| Sets how it works through the connectors
declineOn|false| Scan results that decline the transaction. Used only when the policy scans before the transaction (`position` is `pre`)
[checkTime](#fraudchecktime)|false| Deprecated. Older policies set this instead of `transactionType`, `position`, `attempts`, `cit` and `mit`, and it cannot be combined with them

At least one of `cit` and `mit` must be `scan` or `inform`.

### FraudTransactionType
value | Definition 
---:|:---
"auth" | An authorisation
"capture-auth" | The capture of an earlier authorisation
"capture" | A payment captured without an earlier authorisation

### FraudPosition
value | Definition 
---:|:---
"pre" | Before the transaction. A scan result listed in `declineOn` declines the transaction
"post" | After the transaction, whether it was approved or declined
"post-success" | After the transaction is approved
"post-decline" | After the transaction is declined
"on-demand" | When a fraud scan is requested for the charge through the API, outside any transaction

The result of a scan after the transaction is sent to the merchant in the fraud scan webhook, so they can act on it, for example by voiding an authorisation or refunding a capture.

### FraudAttempts
value | Definition 
---:|:---
"first" | Only the charge's first attempt at the transaction
"every" | Every attempt at the transaction

### FraudMode
value | Definition 
---:|:---
"scan" | Asks the fraud connectors for a result, and sends it to the merchant in the fraud scan webhook. Some providers, such as Kount, charge for each scan
"inform" | Sends the payment to the fraud connectors for their records, without asking for a result. No fraud scan webhook is sent, and nothing is declined. A connector that has no way to receive a payment without scanning it, such as reCAPTCHA, is skipped
"none" | The policy does nothing with the payment

### FraudCheckType
value | Definition 
---:|:---
"all"|Perform a fraud check on all provided connectors
"failover"|Will perform a fraud check on one connector ID at a time, stopping at the first success

### FraudCheckTime
Deprecated. Each value works as the settings below, with both `cit` and `mit` set to `scan`.

value | transactionType | position | attempts
---:|---|---|---
"preauth-first" | `auth` | `pre` | `first`
"preauth-every" | `auth` | `pre` | `every`
"auth-success" | `auth` | `post-success` | `every`
"ondemand" | | `on-demand` |

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
