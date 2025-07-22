package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	res := map[string]string{}
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			// res[v.StructField()] = v.Error()
			res[v.StructField()] = TranslateTag(v)

		}
	}
	return res
}

func TranslateTag(fd validator.FieldError) string {
	switch fd.ActualTag() {
	case "required":
		return fmt.Sprintf("Field %s is required", fd.StructField())
	case "unique":
		return fmt.Sprintf("Field %s must be unique", fd.StructField())
	case "min":
		return fmt.Sprintf("Field %s size minimal %s", fd.StructField(), fd.Param())
	}
	return "Validation failed"
}
