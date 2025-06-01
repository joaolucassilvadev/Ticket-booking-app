package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaolucassilvadev/Ticket-booking-app/handlers"
	"github.com/joaolucassilvadev/Ticket-booking-app/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "ticketBooking",
		ServerHeader: "Fiber",
	})
	//repositories
	EventRepository := repositories.NewEventRepository(nil)

	//routing
	server := app.Group("/api")

	//handlers
	handlers.NewEventHandler(server.Group("/event"), EventRepository)

	app.Listen(":4090")
}
