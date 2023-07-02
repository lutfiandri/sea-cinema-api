package utils

import (
	"reflect"

	"sea-cinema-api/internal/contract"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ParseValidationError(c *fiber.Ctx, errs validator.ValidationErrors) []contract.ValidationErrorResponse {
	var validationErrorsResponse []contract.ValidationErrorResponse

	// get json tags
	jsonTags := make(map[string]string)
	request := c.Locals("request")
	t := reflect.TypeOf(request)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTags[field.Name] = field.Tag.Get("json")
	}

	// create error list
	for _, err := range errs {
		var element contract.ValidationErrorResponse

		element.FailedField = jsonTags[err.Field()]
		element.Tag = err.Tag()
		element.Value = err.Param()

		validationErrorsResponse = append(validationErrorsResponse, element)
	}

	return validationErrorsResponse
}
