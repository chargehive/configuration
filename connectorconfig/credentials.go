package connectorconfig

import (
	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type Credentials interface {
	GetLibrary() Library
	GetSupportedTypes() []LibraryType
	GetSecureFields() []*string
	Validate() error
	ToConnector() connector.Connector
	FromJson(input []byte) error
	SupportsSca() bool
	SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool
	SupportsCountry(country string) bool
	CanPlanModeUse(mode environment.Mode) bool // Determine if this plan mode can use this configuration
	IsRecoveryAgent() bool
}
