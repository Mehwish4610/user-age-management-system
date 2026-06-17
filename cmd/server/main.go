package main

import (
	"log"
	"os"

	"go-user-age-api/config"
	"go-user-age-api/internal/handler"
	"go-user-age-api/internal/repository"
	"go-user-age-api/internal/routes"
	"go-user-age-api/internal/service"
	"go-user-age-api/internal/logger"
	"go-user-age-api/internal/middleware"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using environment variables")
	}

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer db.Close()

	app := fiber.New()

	zapLogger := logger.NewLogger()
	defer zapLogger.Sync()

	zapLogger.Info("Server started successfully")

	app.Use(middleware.RequestMiddleware(zapLogger))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Go Fiber server is running",
		})
	})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, zapLogger)

	routes.UserRoutes(app, userHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = os.Getenv("PORT")
	}
	if port == "" {
		port = "8081"
	}

	log.Println("Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}