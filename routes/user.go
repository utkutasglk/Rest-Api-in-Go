package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/utkutasglk/Rest-Api-in-Go/database"
	"github.com/utkutasglk/Rest-Api-in-Go/models"
)

type UserSerializer struct{
	//
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}
func CreateResponseUser(userModel models.User) UserSerializer{
	return UserSerializer{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error{
	var user models.User

	if err := c.BodyParser(&user); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}