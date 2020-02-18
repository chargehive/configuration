#Slack Integration
The Slack integration config defines the slack account and settings in order to get updates from ChargeHive to be sent to slack.

##Format
As with all configs, the standard wrapper is used.
```json5
{
  "kind": "Integration.Slack",            // Must be set to "Integration.Slack"
  "metadata": {
    "projectId": "test-project",          // Must be set to the ChargeHive Project ID you were issued with
    "name": "test-slack-integration"      // Set this to a memorable name for the slack integration, no spaces, all lowercase
  },
  "spec": {
    "accessToken": "xxxxxxxxxxx",         // AccessToken slack access token
    "scopes": [
      "xxxxxxxxxxx"                       // Scopes for OAuth authentication
    ],
    "teamName": "xxxxxxxxxxx",            // TeamName for posting
    "teamID": "xxxxxxxxxxx",              // TeamID for posting
    "webhook": {                          // Webhook endpoint
      "url": "",                          // Url is the slack webhook URL
      "channel": "",                      // Channel is the slack channel to post in
      "configurationUrl": ""              // ConfigurationUrl is the slack endpoint for configuration
    },
    "transactionNotifications": true,     // TransactionNotifications indicates the action to perform
  }
}
```   
