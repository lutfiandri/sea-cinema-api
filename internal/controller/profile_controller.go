package controller

import (
	"sea-cinema-api/internal/contract"
	"sea-cinema-api/internal/middleware"
	"sea-cinema-api/internal/model"
	"sea-cinema-api/internal/service"
	"sea-cinema-api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type ProfileController interface {
	InitRoute()

	GetProfile(c *fiber.Ctx) error
	TopUpBalance(c *fiber.Ctx) error
}

type profileController struct {
	app            *fiber.App
	profileService service.ProfileService
}

func NewProfileController(app *fiber.App, profileService service.ProfileService) ProfileController {
	return &profileController{
		app:            app,
		profileService: profileService,
	}
}

func (controller *profileController) InitRoute() {
	api := controller.app.Group("/profile")
	api.Get("/", middleware.NewAuthenticator(), controller.GetProfile)
	api.Post("/topup", middleware.NewAuthenticator(), controller.TopUpBalance)
}

func (controller *profileController) GetProfile(c *fiber.Ctx) error {
	claims := c.Locals("claims").(model.JWTClaims)

	response, err := controller.profileService.GetProfile(c.Context(), claims)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK).JSON(contract.NewSuccessResponse(response))

	return nil
}

func (controller *profileController) TopUpBalance(c *fiber.Ctx) error {
	var request contract.TopUpBalanceRequest
	parseOptions := utils.ParseOptions{ParseBody: true}
	if err := utils.ParseAndValidateRequest[contract.TopUpBalanceRequest](c, &request, parseOptions); err != nil {
		return err
	}

	claims := c.Locals("claims").(model.JWTClaims)

	response, err := controller.profileService.TopUpBalance(c.Context(), claims, request)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK).JSON(contract.NewSuccessResponse(response))

	return nil
}
