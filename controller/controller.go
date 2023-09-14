package controller

import (
	"github.com/doo-dev/my-task/db/postgres"
	"github.com/doo-dev/my-task/logger"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetLocal[T any](c *fiber.Ctx, key string) T {
	return c.Locals(key).(T)
}

func CheckDbHealth(c *fiber.Ctx) error {
	db := GetLocal[*postgres.PgDB](c, "db")
	lg := GetLocal[*logger.Logger](c, "logger")

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
