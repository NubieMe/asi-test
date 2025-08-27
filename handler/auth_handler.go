package handler

import (
	"test-asi/model"
	"test-asi/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type AuthHandler struct {
	authService  service.AuthService
	sessionStore *session.Store
}

func NewAuthHandler(authService service.AuthService, sessionStore *session.Store) *AuthHandler {
	return &AuthHandler{
		authService:  authService,
		sessionStore: sessionStore,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var payload model.RegisterPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid data format"})
	}

	if err := h.authService.Register(&payload); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Register success"})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var payload model.LoginPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid data format"})
	}

	_, err := h.authService.Login(&payload)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}

	sess, err := h.sessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}
	sess.Set("username", payload.Username)
	sess.Save()

	return c.JSON(fiber.Map{"message": "Login success", "note": "disini saya tidak implementasi JWT token, hanya sesuai dengan instruksi yang diberikan saja, terima kasih"})
}
