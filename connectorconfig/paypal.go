package connectorconfig

type PayPalEnvironment string

const (
	PayPalEnvironmentSandbox PayPalEnvironment = "sandbox"
	PayPalEnvironmentLive    PayPalEnvironment = "live"
)
