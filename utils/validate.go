package utils

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/chargehive/configuration/connectorconfig"
	"github.com/chargehive/configuration/object"
	"github.com/chargehive/configuration/selector"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

// Validate will check all structs for structural and invalid parameters, will return an empty map if valid
func Validate(rawJson []byte, version string) map[string]string {
	result := map[string]string{}
	if version != "v1" {
		result["config version"] = "version mismatch"
		return result
	}

	o, err := object.FromJsonStrict(rawJson)
	if err != nil {
		result["json"] = err.Error()
		return result
	}

	errs := validatePredicates(o.Selector.Expressions)
	if errs != "" {
		result["validation expressions"] = errs
		return result
	}

	validate = validator.New()
	// register custom validators
	err = validate.RegisterValidation("predicate-key", PredicateKeysValidator)
	if err == nil {
		err = validate.RegisterValidation("connector-library", ConnectorLibraryValidator)
	}
	if err != nil {
		result["validation registration"] = err.Error()
		return result
	}

	enLocale := en.New()
	uni = ut.New(enLocale, enLocale)
	trans, _ := uni.GetTranslator("en")

	err = enTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		result["validation translation"] = err.Error()
		return result
	}

	// validate serialized connector credentials
	if o.Kind == "Connector" {
		c, ok := o.Spec.(*connector.Connector)
		if !ok {
			result["spec error"] = "spec is not a connector"
			return result
		}
		conn, err := connectorconfig.GetCredentialsStrict(c)
		if err != nil {
			result["connector error"] = err.Error()
			return result
		}

		if conn.GetLibrary() == connectorconfig.LibraryNone {
			return result
		}

		if c.ConfigurationID != "" {
			return result
		}

		err = validate.Struct(conn)
		if err != nil {
			errs := err.(validator.ValidationErrors)
			result = errs.Translate(trans)
		}
	}

	// validate main struct
	err = validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		transErrs := errs.Translate(trans)
		// merge translated errors
		for k, v := range transErrs {
			result[k] = v
		}
	}
	return result
}

func PredicateKeysValidator(fl validator.FieldLevel) bool {
	_, ok := selector.KeyRegister[selector.Key(fl.Field().String())]
	return ok || strings.HasPrefix(fl.Field().String(), "charge.label.")
}

func ConnectorLibraryValidator(fl validator.FieldLevel) bool {
	_, ok := connectorconfig.LibraryRegister[connectorconfig.Library(fl.Field().String())]
	return ok
}

func validatePredicates(predicates []selector.Predicate) string {

	buff := bytes.NewBufferString("")

	for _, predicate := range predicates {

		switch predicate.Operator {
		case selector.PredicateOperatorEqual, selector.PredicateOperatorNotEqual:
			if len(predicate.Values) != 1 {
				buff.WriteString(fmt.Sprintf("An %s operator must have exactly one value\n", predicate.Operator))
			}
		case selector.PredicateOperatorIn, selector.PredicateOperatorNotIn, selector.PredicateOperatorInLike, selector.PredicateOperatorNotInLike:
			if len(predicate.Values) <= 1 {
				buff.WriteString(fmt.Sprintf("An %s operator must have at least two values\n", predicate.Operator))
			}
		}
	}

	return strings.TrimSpace(buff.String())
}
