package validator

import (
	"errors"
	enLocale "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"net/http"
)

type JsonValidator struct {
	validator  v10.Validate
	translator ut.Translator
}

func New() *JsonValidator {
	validator := v10.New()

	universalTranslator := ut.New(enLocale.New())
	englishTranslator, _ := universalTranslator.GetTranslator("en")

	err := en.RegisterDefaultTranslations(validator, englishTranslator)
	if err != nil {
		panic(err)
	}

	return &JsonValidator{
		validator:  *validator,
		translator: englishTranslator,
	}
}

func (v *JsonValidator) Validate(i any) error {
	err := v.validator.Struct(i)
	if err != nil {
		return v.handleErrors(err)
	}

	return nil
}

func (v *JsonValidator) handleErrors(err error) error {
	var vr v10.ValidationErrors
	if errors.As(err, &vr) {
		var details []map[string]string

		for i := 0; i < len(vr); i++ {
			details = append(details, map[string]string{
				(vr)[i].Field(): (vr)[i].Translate(v.translator),
			})
		}

		return &ValidationError{
			Message:        validationErrorMessage,
			Details:        details,
			HttpStatusCode: http.StatusBadRequest,
		}
	}

	return err
}
