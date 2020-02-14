#Initiator Configuration



```json5
{
  "kind": "Initiator",
  "metadata": {
    "projectId": "test-project",
    "name": "test-initiator"
  },
  "spec": {
    "type": "auth", // can either be "auth" or "renewal"
    "InitialConnector": "config", //
    "AttemptConfig": {
      "PoolType": "single",
      "MethodSelector": "primary",
      "ConnectorLimit": 0,
      "MethodLimit": 0,
      "CascadeDelay": null,
      "OverridePoolConnectorIDs": [
        "braintree-connector"
      ]
    }
  }
}
```

	// ConnectorSelectorNone indicates to use no connectors
	// (this is the same as setting empty and is the default value)
	ConnectorSelectorNone ConnectorSelector = "none"

	// ConnectorSelectorStickyFirst indicates the connector should stick to the first connector
	// that was sucessful for that payment method
	ConnectorSelectorStickyFirst ConnectorSelector = "sticky-first"

	// ConnectorSelectorStickyLast indicates the connector should stick to the most recent connector
	// that was sucessful for that payment method
	ConnectorSelectorStickyLast ConnectorSelector = "sticky-last"

	// ConnectorSelectorStickyAny indicates to use any connector that has a past success for that payment method
	ConnectorSelectorStickyAny ConnectorSelector = "sticky-any"

	// ConnectorSelectorStickyVerified indicates to use any connector that has a past success and has been verified for that payment method
	ConnectorSelectorStickyVerified ConnectorSelector = "sticky-verified"

	// ConnectorSelectorConfig indicates to use only the connectors specified in the configuration
	ConnectorSelectorConfig ConnectorSelector = "config"