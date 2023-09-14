package main

import (
	"fmt"
	"github.com/doo-dev/my-task/controller"
	"github.com/doo-dev/my-task/db/postgres"
	"github.com/doo-dev/my-task/logger"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	dsn := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", "hossein", "123456", "task")
	pgRepo, err := postgres.New(dsn)
	if err != nil {
		log.Printf("cant open database: %s", err)
	}
	defer pgRepo.Close()

	myLogger := logger.New(log.New(os.Stdout, "", log.LstdFlags))

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		SetLocal[*postgres.PgDB](c, "db", pgRepo)
		SetLocal[*logger.Logger](c, "logger", myLogger)
		return c.Next()
	})
	
	app.Get("/db-health", controller.CheckDbHealth)

	app.Listen(":8080")
}

func SetLocal[T any](c *fiber.Ctx, key string, value T) {
	c.Locals(key, value)
}
