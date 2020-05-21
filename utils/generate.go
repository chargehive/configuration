package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chargehive/configuration/connectorconfig"
	"github.com/chargehive/configuration/object"
	"github.com/chargehive/configuration/selector"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/configuration/v1/integration"
	"github.com/chargehive/configuration/v1/policy"
	"github.com/chargehive/configuration/v1/scheduler"
	"time"
)

type Template string

const (
	// connectors
	confConnAuthorize        Template = "con_authorize"
	confConnBrainTree        Template = "con_brainTree"
	confConnChargeHive       Template = "con_chargeHive"
	confConnCyberSource      Template = "con_cyberSource"
	confConnMaxMind          Template = "con_maxMind"
	confConnPayPalExpress    Template = "con_payPalExpressCheckout"
	confConnPayPalWPP        Template = "con_payPalWPP"
	confConnPaysafe          Template = "con_paysafe"
	confConnPaysafeApplePay  Template = "con_paysafeApplePay"
	confConnPaysafeGooglePay Template = "con_paysafeGooglePay"
	confConnQualPay          Template = "con_qualPay"
	confConnSandbox          Template = "con_sandbox"
	confConnStripe           Template = "con_stripe"
	confConnVindicia         Template = "con_vindicia"
	confConnWorldPay         Template = "con_worldPay"

	// connector Pool
	confConnectorPool Template = "con_pool"

	// integration
	confSlack Template = "int_slack"

	// policy
	confPolCascade       Template = "pol_cascade"
	confPolChargeExpiry  Template = "pol_chargeExpiry"
	confPolFraud         Template = "pol_fraud"
	confPolMethodLock    Template = "pol_methodLock"
	confPolMethodUpgrade Template = "pol_methodUpgrade"
	confPolMethodVerify  Template = "pol_methodVerify"
	confPolSCA           Template = "pol_sca"

	// scheduler
	confSchInitiator  Template = "sch_initiator"
	confSchOnDemand   Template = "sch_onDemand"
	confSchRefund     Template = "sch_refund"
	confSchSequential Template = "sch_sequential"
)

var Templates = map[Template]string{
	confConnAuthorize:        "Connector: Authorize.net",
	confConnBrainTree:        "Connector: Braintree",
	confConnChargeHive:       "Connector: ChargeHive (fraud)",
	confConnCyberSource:      "Connector: Cybersource (fraud)",
	confConnMaxMind:          "Connector: MaxMind (fraud)",
	confConnPayPalExpress:    "Connector: Paypal Express Checkout",
	confConnPayPalWPP:        "Connector: Paypal Website Payments Pro",
	confConnPaysafe:          "Connector: Paysafe",
	confConnPaysafeApplePay:  "Connector: Paysafe Apple Pay",
	confConnPaysafeGooglePay: "Connector: Paysafe Google Pay",
	confConnQualPay:          "Connector: QualPay",
	confConnSandbox:          "Connector: ChargeHive Sandbox",
	confConnStripe:           "Connector: Stripe",
	confConnVindicia:         "Connector: Vindicia",
	confConnWorldPay:         "Connector: Worldpay",
	confConnectorPool:        "Connector Pool",
	confSlack:                "Integration: Slack",
	confPolCascade:           "Policy: Cascade",
	confPolChargeExpiry:      "Policy: Charge Expiry",
	confPolFraud:             "Policy: Fraud",
	confPolMethodLock:        "Policy: Method Lock",
	confPolMethodUpgrade:     "Policy: Method Upgrade",
	confPolMethodVerify:      "Policy: Method Verify",
	confPolSCA:               "Policy: SCA",
	confSchInitiator:         "Schedule: Initiator",
	confSchOnDemand:          "Schedule: OnDemand",
	confSchRefund:            "Schedule: Refund",
	confSchSequential:        "Schedule: Sequential",
}

var chg = "change-me"

// Generate can be used to create a basic but valid config of any type
func Generate(conf Template, version string, pretty bool) ([]byte, error) {

	if version != "v1" {
		return nil, errors.New("version mismatch")
	}

	spec, err := buildSpec(conf)
	if err != nil {
		return nil, err
	}
	def := object.DefinitionFromSpec(spec)
	def.MetaData.Name = chg
	def.MetaData.ProjectID = chg
	def.MetaData.Disabled = false
	def.Selector = buildSelector()

	var data []byte
	if pretty {
		data, err = json.MarshalIndent(def, "", "    ")
	} else {
		data, err = json.Marshal(def)
	}

	// run validation against generated configs to ensure compliance
	errs := Validate(data, version)
	if len(errs) > 0 {
		fmt.Println(errs)
		err = errors.New("generated config does not pass validation")
	}

	return data, err
}

func buildSpec(conf Template) (object.Specification, error) {
	switch conf {
	case confConnAuthorize:
		j, _ := json.Marshal(connectorconfig.AuthorizeCredentials{APILoginID: &chg, TransactionKey: &chg, Environment: "sandbox"})
		return connector.Connector{Library: string(connectorconfig.LibraryAuthorize), Configuration: j}, nil
	case confConnBrainTree:
		j, _ := json.Marshal(connectorconfig.BraintreeCredentials{PublicKey: &chg, PrivateKey: &chg, MerchantID: chg, MerchantAccountID: chg, Currency: "USD", Environment: "sandbox"})
		return connector.Connector{Library: string(connectorconfig.LibraryBraintree), Configuration: j}, nil
	case confConnChargeHive:
		j, _ := json.Marshal(connectorconfig.ChargeHiveCredentials{})
		return connector.Connector{Library: string(connectorconfig.LibraryChargeHive), Configuration: j}, nil
	case confConnCyberSource:
		j, _ := json.Marshal(connectorconfig.CyberSourceCredentials{MerchantID: chg, TransactionKey: &chg, Environment: "test"})
		return connector.Connector{Library: string(connectorconfig.LibraryCyberSource), Configuration: j}, nil
	case confConnMaxMind:
		j, _ := json.Marshal(connectorconfig.MaxMindCredentials{AccountID: chg, LicenceKey: &chg, ServiceType: 0})
		return connector.Connector{Library: string(connectorconfig.LibraryMaxMind), Configuration: j}, nil
	case confConnPayPalExpress:
		j, _ := json.Marshal(connectorconfig.PayPalExpressCheckoutCredentials{APIUsername: &chg, APIPassword: &chg, APISignature: &chg, SupportedCurrencies: []string{"USD"}, Environment: "sandbox"})
		return connector.Connector{Library: string(connectorconfig.LibraryPayPalExpressCheckout), Configuration: j}, nil
	case confConnPayPalWPP:
		j, _ := json.Marshal(connectorconfig.PayPalWebsitePaymentsProCredentials{
			APIUsername:            &chg,
			APIPassword:            &chg,
			APISignature:           &chg,
			SupportedCurrencies:    []string{"USD"},
			CardinalProcessorID:    &chg,
			CardinalMerchantID:     &chg,
			CardinalTransactionPw:  &chg,
			CardinalTransactionURL: &chg,
			CardinalAPIIdentifier:  &chg,
			CardinalAPIKey:         &chg,
			CardinalOrgUnitID:      &chg,
			Environment:            "sandbox",
		})
		return connector.Connector{Library: string(connectorconfig.LibraryPayPalWebsitePaymentsPro), Configuration: j}, nil
	case confConnPaysafe:
		j, _ := json.Marshal(connectorconfig.PaySafeCredentials{
			Acquirer:               chg,
			AccountID:              chg,
			APIUsername:            &chg,
			APIPassword:            &chg,
			Environment:            "MOCK",
			Currency:               "USD",
			UseVault:               new(bool),
			SingleUseTokenPassword: new(string),
			SingleUseTokenUsername: "",
		})
		return connector.Connector{Library: string(connectorconfig.LibraryPaySafe), Configuration: j}, nil
	case confConnQualPay:
		j, _ := json.Marshal(connectorconfig.QualpayCredentials{APIKey: &chg, MerchantID: 1, Environment: "test"})
		return connector.Connector{Library: string(connectorconfig.LibraryQualPay), Configuration: j}, nil
	case confConnSandbox:
		j, _ := json.Marshal(connectorconfig.SandboxCredentials{Mode: "dynamic"})
		return connector.Connector{Library: string(connectorconfig.LibrarySandbox), Configuration: j}, nil
	case confConnStripe:
		j, _ := json.Marshal(connectorconfig.StripeCredentials{APIKey: &chg})
		return connector.Connector{Library: string(connectorconfig.LibraryStripe), Configuration: j}, nil
	case confConnVindicia:
		j, _ := json.Marshal(connectorconfig.VindiciaCredentials{Login: chg, Password: &chg, HMACKey: &chg, PGPPrivateKey: &chg, Environment: "development"})
		return connector.Connector{Library: string(connectorconfig.LibraryVindicia), Configuration: j}, nil
	case confConnWorldPay:
		j, _ := json.Marshal(connectorconfig.WorldpayCredentials{
			Username:                 &chg,
			Password:                 &chg,
			MerchantID:               chg,
			ReportGroup:              chg,
			Environment:              "sandbox",
			CardinalApiIdentifier:    &chg,
			CardinalApiKey:           &chg,
			CardinalOrgUnitId:        &chg,
			AppleMerchantIdentifier:  chg,
			AppleMerchantDisplayName: chg,
			AppleInitiative:          "web",
			AppleInitiativeContext:   chg,
			AppleMerchantCertificate: &chg,
			AppleMerchantPrivateKey:  &chg,
		})
		return connector.Connector{Library: string(connectorconfig.LibraryWorldpay), Configuration: j}, nil
	case confConnectorPool:
		return connector.Pool{Restriction: "unrestricted", Connectors: []connector.PoolItem{{ConnectorID: chg}}}, nil
	case confSlack:
		return integration.Slack{AccessToken: chg, TeamName: chg, TeamID: chg, TransactionNotifications: new(bool), Webhook: &integration.SlackWebhook{Url: chg, Channel: chg, ConfigurationUrl: chg}}, nil
	case confPolCascade:
		return policy.CascadePolicy{Rules: []policy.CascadeRule{{Library: connectorconfig.Library(chg), OriginalResponseCode: chg}}}, nil
	case confPolChargeExpiry:
		return policy.ChargeExpiryPolicy{}, nil
	case confPolFraud:
		return policy.FraudPolicy{ConnectorIDs: []string{chg}, CheckTime: "preauth-first", CheckType: "all"}, nil
	case confPolMethodLock:
		return policy.MethodLockPolicy{Duration: 1, Reason: chg}, nil
	case confPolMethodUpgrade:
		return policy.MethodUpgradePolicy{ExtendExpiry: new(bool)}, nil
	case confPolMethodVerify:
		return policy.MethodVerifyPolicy{Amount: 100, AmountCurrency: "GBP", ConnectorID: chg, VerifyMethodOnTokenization: new(bool)}, nil
	case confPolSCA:
		return policy.ScaPolicy{ShouldIdentify: new(bool), ShouldChallengeOptional: new(bool), ShouldByPassChallenge: "cascade", ShouldAuthOnError: new(bool), RequireSca: new(bool)}, nil
	case confSchInitiator:
		return scheduler.Initiator{Type: scheduler.InitiatorTypeAuth, InitialConnector: scheduler.ConnectorSelectorConfig, AttemptConfig: &scheduler.AttemptConfig{PoolType: scheduler.PoolTypeCascade, MethodSelector: scheduler.MethodSelectorPrimaryMethod, OverridePoolConnectorIDs: []string{}, CascadeDelay: new(time.Duration)}}, nil
	case confSchOnDemand:
		return scheduler.OnDemand{Schedule: scheduler.Schedule{AttemptConfig: scheduler.AttemptConfig{PoolType: "single", MethodSelector: "primary", CascadeDelay: new(time.Duration)}, TimeDelayOrigin: "initialisation", TimeDelaySync: "Earliest", TimeSyncZone: "UTC"}}, nil
	case confSchRefund:
		return scheduler.Refund{Schedules: map[int]scheduler.ScheduleRefund{0: {0}}}, nil
	case confSchSequential:
		return scheduler.Sequential{Schedules: map[int]scheduler.Schedule{0: {AttemptConfig: scheduler.AttemptConfig{PoolType: "single", MethodSelector: "primary", CascadeDelay: new(time.Duration)}, TimeDelayOrigin: "initialisation", TimeDelaySync: "Earliest", TimeSyncZone: "UTC"}}}, nil
	}
	return nil, errors.New("invalid config to generate")
}

func buildSelector() selector.Selector {
	return selector.Selector{Priority: 50, Expressions: []selector.Predicate{{Key: selector.KeyChargeAmountCurrency, Operator: selector.PredicateOperatorEqual, Values: []string{"GBP"}}}}
}
