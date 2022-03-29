package pos

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	repository *ProductRepository
}

func (handler *ProductHandler) GetAllProduct(c *fiber.Ctx) error {

	var products []Products = handler.repository.FindAllProduct(c)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data": fiber.Map{
			"products": products,
			"meta": fiber.Map{
				"total": len(products),
				"limit": c.Query("limit"),
				"skip":  c.Query("skip"),
			},
		},
	})
}

func (handler *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	product, err := handler.repository.FindProduct(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Product Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    product,
	})
}

func (handler *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	data := new(Products)

	if len(c.Body()) == 0 {
		return c.Status(400).JSON(fiber.Map{"success": false, "message": "body ValidationError: \"name\" is required"})
	}

	var p struct {
		Name string `json:"name"`
	}
	err := json.Unmarshal(c.Body(), &p)
	data.Name = p.Name

	item, err := handler.repository.CreateProduct(*data)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Failed creating item",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    item,
	})
}

func (handler *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "ID not found",
			"error":   err,
		})
	}

	product, err := handler.repository.FindProduct(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Product Not Found",
		})
	}

	var p struct {
		Name string `json:"name"`
	}
	err = json.Unmarshal(c.Body(), &p)
	product.Name = p.Name

	item, err := handler.repository.SaveProduct(product)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Product Not Found",
		})
	}

	return c.JSON(item)
}

func (handler *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Failed deleting product",
			"err":     err,
		})
	}
	RowsAffected := handler.repository.DeleteProduct(id)
	if RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Product Not Found",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
	})
}

func NewProductHandler(repository *ProductRepository) *ProductHandler {
	return &ProductHandler{
		repository: repository,
	}
}
