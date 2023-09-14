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

func CheckDbHealth(db Repository, lg Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := db.CheckConnection(); err != nil {
			lg.Error(err.Error())
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		lg.Info("db is up and running")
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "everything is good and healthy",
		})
	}
}
