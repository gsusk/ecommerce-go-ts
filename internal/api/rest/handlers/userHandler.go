package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(handl *rest.RestHandler) {
	app := handl.App

	svc := service.UserService{
		Repo: repository.NewUserRepository(handl.DB),
	}

	handler := &UserHandler{
		svc,
	}

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
	user := dto.UserSignUp{}
	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid input please try again",
		})
	}

	token, err := h.svc.Signup(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "unexpected server error trying to signup",
			},
		)
	}

	return ctx.Status(http.StatusCreated).JSON(
		&fiber.Map{
			"message": token,
		},
	)

}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	user := dto.UserLogin{}
	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid input please try again",
		})
	}

	token, err := h.svc.Login(user.Email, user.Password)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "unexpected server error trying to login",
			},
		)
	}

	return ctx.Status(http.StatusInternalServerError).JSON(
		&fiber.Map{
			"message": token,
		},
	)
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
