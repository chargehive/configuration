package object

import (
	"github.com/chargehive/configuration/selector"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func (d *Definition) Validate() validator.ValidationErrorsTranslations {
	validate = validator.New()
	err := validate.RegisterValidation("predicate-key", PredicateKeysValidator)
	err = validate.RegisterValidation("predicate-operator", PredicateOperatorValidator)
	err = validate.RegisterValidation("predicate-operator-conversion", OperatorConversionValidator)
	if err != nil {
		return map[string]string{}
	}

	enLocale := en.New()
	uni = ut.New(enLocale, enLocale)
	trans, _ := uni.GetTranslator("en")

	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	err = validate.Struct(d)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		return errs.Translate(trans)
	}
	return map[string]string{}
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
