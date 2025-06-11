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

	// routing
	server := app.Group("/api")

	// handlers
	handlers.NewEventHandler(server.Group("/event"), EventRepository)

	// Porta fixa 4000
	if err := app.Listen(":4000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
