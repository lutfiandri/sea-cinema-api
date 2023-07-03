package middleware

import (
	"errors"

	"sea-cinema-api/internal/contract"

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
	if valErrsResponse, ok := c.Locals("validation_errors_response").([]contract.ValidationErrorResponse); ok {
		validationErrorsResponse = &valErrsResponse
	}

	response := contract.NewErrorResponse(err.Error(), validationErrorsResponse)
	return c.Status(code).JSON(response)
}
