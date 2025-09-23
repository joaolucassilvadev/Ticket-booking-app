package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joaolucassilvadev/Ticket-booking-app/models"
)

type TicketHandler struct {
	repository models.TicketRepository
}

func (h *TicketHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	tickets, err := h.repository.GetMany(context)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"Status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status":  "success",
		"message": "",
		"data":    tickets,
	})
}

func (h *TicketHandler) GetOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	ticketId, _ := strconv.Atoi(ctx.Params("ticketId"))
	ticket, err := h.repository.GetOne(context, uint(ticketId))
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"Status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status":  "success",
		"message": "",
		"data":    ticket,
	})
}

func (h *TicketHandler) CreateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()
	ticket := &models.Ticket{}
	if err := ctx.BodyParser(ticket); err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"Status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}
	ticket, err := h.repository.CreateOne(context, ticket)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"Status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Status":  "success",
		"message": "",
		"data":    ticket,
	})
}

func (h *TicketHandler) ValidateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	validatebody := &models.ValidateTicket{}

	if err := ctx.BodyParser(&validatebody); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"Status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	validateData := make(map[string]interface{})
	validateData["entered"] = true
	ticket, err := h.repository.UpdateOne(context, validatebody.TicketId, validateData)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"Status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status":  "success",
		"message": "Welcome to the show",
		"data":    ticket,
	})
}

func NewTicketHandler(router fiber.Router, repository models.TicketRepository) {
	handler := &TicketHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:ticketId", handler.GetOne)
	router.Post("/validate", handler.ValidateOne)
}
