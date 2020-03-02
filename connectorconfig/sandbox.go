package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type SandboxMode string

const (
	SandboxModeDynamic       SandboxMode = "dynamic"
	SandboxModeOffline       SandboxMode = "offline"
	SandboxModeDelayed       SandboxMode = "delayed"
	SandboxModeRandomTimeout SandboxMode = "random-timeout"
	SandboxModeChaos         SandboxMode = "chaos"
)

type SandboxCredentials struct {
	Mode                SandboxMode `json:"mode" yaml:"mode" validate:"oneof=dynamic offline delayed random-timeout chaos"`
	TransactionIDPrefix string      `json:"yransactionIDPrefix" yaml:"transactionIDPrefix" validate:"-"`
}

func (c SandboxCredentials) GetLibrary() Library {
	return LibrarySandbox
}

func (c *SandboxCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *SandboxCredentials) Validate() error {
	return nil
}

func (c *SandboxCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *SandboxCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *SandboxCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
