package validation

import (
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

func Validate(rawJson []byte) validator.ValidationErrorsTranslations {
	result := map[string]string{}

	d, err := object.FromJson(rawJson)
	if err != nil {
		result["json"] = err.Error()
		return result
	}

	validate = validator.New()
	// register custom validators
	err = validate.RegisterValidation("predicate-key", PredicateKeysValidator)
	err = validate.RegisterValidation("connector-library", ConnectorLibraryValidator)

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
	if d.Kind == "Connector" {
		c, ok := d.Spec.(*connector.Connector)
		if !ok {
			result["connector error"] = "spec is not a connector"
			return result
		}
		conn, err := connectorconfig.GetCredentials(c)
		if err != nil {
			result["connector error"] = "invalid connector in config"
			return result
		}
		err = validate.Struct(conn)
		if err != nil {
			errs := err.(validator.ValidationErrors)
			result = errs.Translate(trans)
		}
	}

	// validate main struct
	err = validate.Struct(d)
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
	return ok
}

func ConnectorLibraryValidator(fl validator.FieldLevel) bool {
	_, ok := connectorconfig.LibraryRegister[connectorconfig.Library(fl.Field().String())]
	return ok
}
