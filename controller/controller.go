package controller

import (
	"github.com/doo-dev/my-task/db/postgres"
	"github.com/doo-dev/my-task/logger"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CheckDbHealth(db *postgres.PgDB, lg *logger.Logger) fiber.Handler {
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
