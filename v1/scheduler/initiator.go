package scheduler

import "github.com/chargehive/configuration/object"

const KindInitiator object.Kind = "Initiator"

type Initiator struct {
	Type             Type
	InitialConnector ConnectorSelector
	// AttemptConfig to be used when using ConnectorSelectorConfig
	AttemptConfig *AttemptConfig
}

func (i Initiator) GetKind() object.Kind { return KindInitiator }
func (i Initiator) GetVersion() string   { return "v1" }
