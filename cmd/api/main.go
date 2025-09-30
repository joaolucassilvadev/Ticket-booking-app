package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joaolucassilvadev/Ticket-booking-app/db"
	"github.com/joaolucassilvadev/Ticket-booking-app/handlers"
	"github.com/joaolucassilvadev/Ticket-booking-app/repositories"
)

func main() {
	// Inicializa o banco de dados com o migrator
	db := db.Init(db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "ticketBooking",
		ServerHeader: "Fiber",
	})

	// repositories
	EventRepository := repositories.NewEventRepository(db)
	TicketRepository := repositories.NewTicketRepository(db)
	AuthRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(AuthRepository)
	// routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthMiddleware(db))
	// handlers
	handlers.NewEventHandler(server.Group("/event"), EventRepository)
	handlers.NewTicketHandler(server.Group("/ticket"), TicketRepository)

	// Porta fixa 4000
	if err := app.Listen(":4000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
