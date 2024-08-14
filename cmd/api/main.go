package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"test-isi/internal/config"
	"test-isi/internal/delivery/http/route"
	"test-isi/internal/repository/postgresql"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Initialize database
	db, err := postgresql.NewConnector(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	cfg.DB = db

	// Perform database migration
	err = postgresql.Migrate(db)
	if err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}

	// Initialize Fiber app
	app := fiber.New()

	app.Use(logger.New())

	// Setup routes
	route.SetupRoutes(app, cfg)

	// Setup error handler middleware

	// Start server
	log.Fatal(app.Listen(cfg.ServerAddress))
}
