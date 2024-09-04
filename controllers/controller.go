package controllers

import (
	"am_user/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

	email := c.Params("email")
	var user models.User

	result := models.DB.Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": result.Error.Error(),
		})
	}

	return c.JSON(user)
}

func Check(c *fiber.Ctx) error {

	username := c.Params("username")
	var user models.User

	result := models.DB.Db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": result.Error.Error(),
		})
	}

	return c.JSON(user)
}

func Update(c *fiber.Ctx) error {

	username := c.Params("username")
	var updatedData models.User

	// Parse the body to get the updated user data
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	if models.DB.Db.Where("username = ?", username).Updates(&updatedData).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username tidak dapat ditemukan.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate.",
	})

}

func Delete(c *fiber.Ctx) error {

	username := c.Params("username")

	// Delete the user record
	result := models.DB.Db.Where("username = ?", username).Delete(&models.User{})
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).SendString("User deleted successfully")
}
