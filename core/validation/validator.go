package validation

import (
	"time"

	"tampayang-backend/config/constant"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())

	Validate.RegisterValidation("dateformat", func(fl validator.FieldLevel) bool {
		dateStr := fl.Field().String()
		_, err := time.Parse(constant.NOW_TIME_FORMAT, dateStr)
		return err == nil
	})
}
