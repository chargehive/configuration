package object

type Kind string

const (
	kindUnknown       Kind = "_unknown_"
	KindInitiator     Kind = "Initiator"
	KindScheduler     Kind = "Scheduler"
	KindConnectorPool Kind = "ConnectorPool"
	KindConnector     Kind = "Connector"
	KindPolicy        Kind = "Policy"
	KindProject       Kind = "Project"
	KindPlacement     Kind = "Placement"
	KindCharge        Kind = "Charge"
)
