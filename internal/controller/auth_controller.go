package controller

import (
	"sea-cinema-api/internal/contract"
	"sea-cinema-api/internal/service"
	"sea-cinema-api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	InitRoute()

	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type authController struct {
	app         *fiber.App
	authService service.AuthService
}

func NewAuthController(app *fiber.App, authService service.AuthService) AuthController {
	return &authController{
		app:         app,
		authService: authService,
	}
}

func (controller *authController) InitRoute() {
	api := controller.app.Group("/auth")
	api.Post("/register", controller.Register)
	api.Post("/login", controller.Login)
}

func (controller *authController) Register(c *fiber.Ctx) error {
	var request contract.RegisterRequest
	parseOptions := utils.ParseOptions{ParseBody: true}
	if err := utils.ParseAndValidateRequest[contract.RegisterRequest](c, &request, parseOptions); err != nil {
		return err
	}

	response, err := controller.authService.Register(c.Context(), request)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusCreated).JSON(contract.NewSuccessResponse(response))

	return nil
}

func (controller *authController) Login(c *fiber.Ctx) error {
	var request contract.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	response, err := controller.authService.Login(c.Context(), request)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusCreated).JSON(contract.NewSuccessResponse(response))

	return nil
}
