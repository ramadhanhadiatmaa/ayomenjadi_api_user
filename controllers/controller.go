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

	user.Prov = "-"
	user.Kab = "-"
	user.Form = "-"
	user.Analy = "-"
	user.Res = 0
	user.To1 = 0
	user.To2 = 0
	user.To3 = 0
	user.To4 = 0
	user.To5 = 0
	user.To6 = 0
	user.To7 = 0
	user.To8 = 0
	user.To9 = 0
	user.To10 = 0

	models.DB.Db.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(user)
}

func Show(c *fiber.Ctx) error {

	email := &models.User{}
	id := c.Params("email")

	if err := models.DB.Db.First(email, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(email)
}

func Check(c *fiber.Ctx) error {

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
