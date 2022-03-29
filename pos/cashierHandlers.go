package pos

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CashierHandler struct {
	repository *CashierRepository
}

func (handler *CashierHandler) GetAll(c *fiber.Ctx) error {

	var cashiers []Cashiers = handler.repository.FindAll(c)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data": fiber.Map{
			"cashiers": cashiers,
			"meta": fiber.Map{
				"total": len(cashiers),
				"limit": c.Query("limit"),
				"skip":  c.Query("skip"),
			},
		},
	})
}

func (handler *CashierHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	cashier, err := handler.repository.Find(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashier,
	})
}

func (handler *CashierHandler) Passcode(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	passcode, err := handler.repository.Passcode(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data": fiber.Map{
			"passcode": passcode,
		},
	})
}

func (handler *CashierHandler) Create(c *fiber.Ctx) error {
	data := new(Cashiers)

	// if err := c.Body("name"); err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "error": err})
	// }

	fmt.Println("------------ Test ------------")

	payload := struct {
		Name string `json:"name"`
	}{}
	c.BodyParser(payload)
	fmt.Println(payload)

	item, err := handler.repository.Create(*data)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Failed creating item",
			"error":   err,
		})
	}

	return c.JSON(item)
}

func (handler *CashierHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "ID not found",
			"error":   err,
		})
	}

	cashier, err := handler.repository.Find(id)

	if err != nil {
		return c.Status(400).SendString("Failed to update, cashier ID does not exist")
	}

	cashierData := new(Cashiers)

	if err := c.BodyParser(cashierData); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	cashier.Name = cashierData.Name
	cashier.Passcode = cashierData.Passcode

	item, err := handler.repository.Save(cashier)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error updating cashier",
			"error":   err,
		})
	}

	return c.JSON(item)
}

func (handler *CashierHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Failed deleting cashier",
			"err":     err,
		})
	}
	RowsAffected := handler.repository.Delete(id)
	if RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
	})
}

func NewCashierHandler(repository *CashierRepository) *CashierHandler {
	return &CashierHandler{
		repository: repository,
	}
}
