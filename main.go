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
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", "doo-dev", "123456", "localhost", "task")
	pgRepo, err := postgres.New(dsn)
	if err != nil {
		log.Printf("cant open database: %s", err)
	}
	defer pgRepo.Close()

	myLogger := logger.New(log.New(os.Stdout, "", log.LstdFlags))

	app := fiber.New()

	app.Get("/db-health", controller.CheckDbHealth(pgRepo, myLogger))

	if err := app.Listen(":8080"); err != nil {
		fmt.Println("cant listen on port: ", 8080)
	}
}
