package main

import (
	"log"
	"log/slog"
	"time"

	"github.com/dyxgou/notas/cmd/api/routes"
	"github.com/dyxgou/notas/pkg/config"
	"github.com/dyxgou/notas/pkg/repositories/sqlite"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const idleTimeout = 5 * time.Second

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	// middlewares
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4321",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(logger.New())
	app.Use(healthcheck.New())

	// connecting to database
	db := sqlite.ConnectClient(config.GetEnv("DB_PATH"))
	defer db.Close()

	slog.Info("database connected successfully")

	api := app.Group("/api")

	val := validator.New(validator.WithRequiredStructEnabled())

	// registering routes
	r := routes.NewRouter(db, val)

	r.RegisterUserGroup(api.Group("/student"))
	r.RegisterSubjectGroup(api.Group("/subject"))
	r.RegisterGradeGroup(api.Group("/grade"))

	// initializing server
	log.Fatal(app.Listen(config.GetEnv("PORT")))
}
