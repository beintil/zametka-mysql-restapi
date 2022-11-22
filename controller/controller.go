package controllers

import (
	"github.com/beintil/zametka-mysql-restapi/database"
	"github.com/beintil/zametka-mysql-restapi/zametka"
	"github.com/gofiber/fiber/v2"
)

func AllZametka(c *fiber.Ctx) error {
	var zametka []zametka.Zametka

	database.DB.Find(&zametka)
	return c.JSON(zametka)
}

func Create(c *fiber.Ctx) error {
	var zametka zametka.Zametka

	if err := c.BodyParser(&zametka); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			`error`: err.Error(), 
		})
	}

	if err := database.DB.Create(&zametka).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			`error`: err.Error(), 
		})
	}

	return c.JSON(zametka)
}


func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var zametka zametka.Zametka

	if database.DB.Delete(&zametka, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"err": "cant't delete data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "ok delete",
	})
}
