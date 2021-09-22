package validate

import (
	"github.com/go-playground/locales/zh"
	zt "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pot-code/gobit/pkg/validate"
	"golang.org/x/text/language"
)

func NewValidator() *validate.ValidatorV10 {
	return validate.NewValidator(validate.AddLocale(language.Chinese, zh.New(), zt.RegisterDefaultTranslations))
}
