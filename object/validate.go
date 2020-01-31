package object

import (
	"github.com/chargehive/configuration/selector"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func (d *Definition) Validate() error {
	validate = validator.New()
	err := validate.RegisterValidation("predicate-key", PredicateKeysValidator)
	err = validate.RegisterValidation("predicate-operator", PredicateOperatorValidator)
	err = validate.RegisterValidation("predicate-operator-conversion", OperatorConversionValidator)
	err = validate.RegisterValidation("kind", KindValidator)
	if err != nil {
		return err
	}

	return validate.Struct(d)
}

func PredicateKeysValidator(fl validator.FieldLevel) bool {
	_, ok := selector.KeyRegister[selector.Key(fl.Field().String())]
	return ok
}

func PredicateOperatorValidator(fl validator.FieldLevel) bool {
	_, ok := selector.PredicateOperatorRegister[selector.PredicateOperator(fl.Field().String())]
	return ok
}

func OperatorConversionValidator(fl validator.FieldLevel) bool {
	_, ok := selector.OperatorConversionRegister[selector.OperatorConversion(fl.Field().String())]
	return ok
}

func KindValidator(fl validator.FieldLevel) bool {
	_, ok := KindRegister[fl.Field().String()]
	return ok
}

var KindRegister = map[string]bool{
	"Connector":           true,
	"ConnectorPool":       true,
	"Integration.Slack":   true,
	"PolicyCascade":       true,
	"PolicyChargeExpiry":  true,
	"PolicyFraud":         true,
	"PolicyMethodLock":    true,
	"PolicyMethodUpgrade": true,
	"PolicySCA":           true,
	"Initiator":           true,
	"SchedulerOnDemand":   true,
	"SchedulerSequential": true,
}
