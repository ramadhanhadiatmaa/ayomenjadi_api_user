package controllers

import (
	"am_user/models"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	var user []models.User

	models.DB.Db.Find(&user)

	return c.Status(fiber.StatusOK).JSON(user)

}

func Create(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	models.DB.Db.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(user)
}

func Show(c *fiber.Ctx) error {

	user := &models.User{}
	id := c.Params("username")

	if err := models.DB.Db.First(user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func Update(c *fiber.Ctx) error {

	user := &models.User{}
	id := c.Params("username")

	if err := models.DB.Db.First(user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	models.DB.Db.Save(user)

	return c.Status(fiber.StatusOK).JSON(user)

}

func Reset(c *fiber.Ctx) error {

	user := &models.User{}
	id := c.Params("username")

	if err := models.DB.Db.First(user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	models.DB.Db.Save(user)

	return c.Status(fiber.StatusOK).JSON(user)
}

func Delete(c *fiber.Ctx) error {

	user := &models.User{}
	id := c.Params("username")

	if err := models.DB.Db.First(user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	models.DB.Db.Delete(user, id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Delete Data"})
}