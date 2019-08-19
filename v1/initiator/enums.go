package initiator

type ConnectorSelector string
type Type string

const (
	ConnectorSelectorStickyFirst    ConnectorSelector = "sticky-first"
	ConnectorSelectorStickyLast     ConnectorSelector = "sticky-last"
	ConnectorSelectorStickyAny      ConnectorSelector = "sticky-any"
	ConnectorSelectorStickyVerified ConnectorSelector = "sticky-verified"
	ConnectorSelectorConfig         ConnectorSelector = "config"
	ConnectorSelectorNone           ConnectorSelector = "none"
)

const (
	TypeAuth    Type = "auth"
	TypeRenewal Type = "renewal"
)
