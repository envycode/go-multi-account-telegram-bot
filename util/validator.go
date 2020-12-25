package util

import (
	"github.com/go-playground/validator"
	"reflect"
	"strings"
)

func UseJsonFieldValidation(v *validator.Validate) {
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
