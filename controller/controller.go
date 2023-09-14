package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Logger interface {
	Warn(msg string)
	Error(msg string)
	Info(msg string)
}

type Repository interface {
	CheckConnection() error
}

type Controller struct {
	repo   Repository
	Logger Logger
}

func New(repo Repository, logger Logger) Controller {
	return Controller{
		repo:   repo,
		Logger: logger,
	}
}

func (cl Controller) CheckDbHealth(c *fiber.Ctx) error {
	if err := cl.repo.CheckConnection(); err != nil {
		cl.Logger.Error(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	cl.Logger.Info("db is up and running")
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "everything is good and healthy",
	})
}
