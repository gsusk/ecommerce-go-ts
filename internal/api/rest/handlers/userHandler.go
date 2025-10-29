package handlers

import (
	"go-ecommerce-app/internal/api/rest"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	//svc UserSvc
}

func SetupUserRoutes(handl *rest.RestHandler) {
	app := handl.App

	handler := UserHandler{}

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	//protected routes
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.VerifyCode)
	app.Get("/profile", handler.GetProfile)
	app.Post("/profile", handler.CreateProfile)
	app.Get("/cart", handler.GetCart)
	app.Post("/cart", handler.AddToCart)
	app.Get("/order", handler.GetOrders)
	app.Get("/order/:id", handler.GetOrder)
	app.Post("/be-seller", handler.CreateSeller)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return nil
}
func (h *UserHandler) VerifyCode(ctx *fiber.Ctx) error {
	return nil
}
func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return nil
}
func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return nil
}
func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return nil
}
func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return nil
}
func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return nil
}
func (h *UserHandler) CreateSeller(ctx *fiber.Ctx) error {
	return nil
}
