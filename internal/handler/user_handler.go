package handler

import (
	"strconv"

	"go-user-age-api/internal/models"
	"go-user-age-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type UserHandler struct {
	service  *service.UserService
	validate *validator.Validate
	logger   *zap.Logger
}

func NewUserHandler(service *service.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		service:  service,
		validate: validator.New(),
		logger:   logger,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("Failed to parse create user request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validate.Struct(req); err != nil {
		h.logger.Error("Create user validation failed", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.service.CreateUser(c.Context(), req)
	if err != nil {
		h.logger.Error("Failed to create user", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Info("User created successfully", zap.Int32("user_id", user.ID))

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    user,
	})
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		h.logger.Error("Invalid user ID in get user request", zap.String("id", c.Params("id")), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := h.service.GetUserByID(c.Context(), int32(id))
	if err != nil {
		h.logger.Error("Failed to get user", zap.Int32("user_id", int32(id)), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Info("User fetched successfully", zap.Int32("user_id", user.ID))

	return c.JSON(fiber.Map{"data": user})
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page value"})
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit value"})
	}

	users, err := h.service.ListUsersPaginated(c.Context(), int32(page), int32(limit))
	if err != nil {
		h.logger.Error("Failed to list users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Info("Users listed successfully",
		zap.Int("count", len(users)),
		zap.Int("page", page),
		zap.Int("limit", limit),
	)

	return c.JSON(fiber.Map{
		"page":  page,
		"limit": limit,
		"data":  users,
	})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		h.logger.Error("Invalid user ID in update user request", zap.String("id", c.Params("id")), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var req models.UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("Failed to parse update user request body", zap.Int32("user_id", int32(id)), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validate.Struct(req); err != nil {
		h.logger.Error("Update user validation failed", zap.Int32("user_id", int32(id)), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.service.UpdateUser(c.Context(), int32(id), req)
	if err != nil {
		h.logger.Error("Failed to update user", zap.Int32("user_id", int32(id)), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Info("User updated successfully", zap.Int32("user_id", user.ID))

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"data":    user,
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		h.logger.Error("Invalid user ID in delete user request", zap.String("id", c.Params("id")), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	if err := h.service.DeleteUser(c.Context(), int32(id)); err != nil {
		h.logger.Error("Failed to delete user", zap.Int32("user_id", int32(id)), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Info("User deleted successfully", zap.Int32("user_id", int32(id)))

	return c.SendStatus(fiber.StatusNoContent)
}