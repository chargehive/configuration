package initiator

import "github.com/chargehive/configuration/object"

type Initiator struct {
	Type             Type
	InitialConnector ConnectorSelector
}

func (i Initiator) GetKind() object.Kind { return object.KindInitiator }
func (i Initiator) GetVersion() string   { return "v1" }
