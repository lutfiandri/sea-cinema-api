package middleware

import (
	"errors"

	"sea-cinema-api/internal/contract"
	"sea-cinema-api/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var ErrorHandler = func(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	var validationErrorsResponse *[]contract.ValidationErrorResponse
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		errs := utils.ParseValidationError(c, validationErrs)
		validationErrorsResponse = &errs
	}

	response := contract.NewErrorResponse(err.Error(), validationErrorsResponse)
	return c.Status(code).JSON(response)
}
