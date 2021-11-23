package controllers

import (
	"project/config"
	"project/models"

	"github.com/gofiber/fiber/v2"
)

// for getting all products
func GetAllbooks(c *fiber.Ctx) error {
	db := config.DB
	var products []models.Book
	db.Find(c.Context(), &products)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All products",
		"data":    products,
	})
}

// for getting single products by ID
func Getbook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.DB
	var product models.Book
	db.Find(c.Context(), &product, id)
	if product.Title == "" {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "No product found with ID",
			"data":    nil,
		})

	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product found",
		"data":    product,
	})
}

// creating a new products in database
func Createbook(c *fiber.Ctx) error {
	db := config.DB
	product := new(models.Book)
	if err := c.BodyParser(product); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create product",
			"data":    err,
		})
	}
	db.Create(&product)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created product",
		"data":    product,
	})
}

// DeleteProduct by the products id
func Deletebook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.DB

	var product models.Book
	db.First(&product, id)
	if product.Title == "" {
		return c.JSON(fiber.Map{"status": "error",
			"message": "No product found with ID",
			"data":    nil,
		})

	}
	db.Delete(&product)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product successfully deleted",
		"data":    nil,
	})
}
