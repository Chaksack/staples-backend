package controllers

import (
	"strconv"

	"github.com/Chaksack/staples-backend/database"
	"github.com/Chaksack/staples-backend/models"
	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.Database.Db.Preload("Role").Find(&users)

	return c.JSON(users)

}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.Database.Db.Create(&user)

	return c.JSON(user)

}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}
	database.Database.Db.Preload("Role").Find(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	database.Database.Db.Model(&user).Updates(user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.Database.Db.Delete(&user)
	return nil
}
