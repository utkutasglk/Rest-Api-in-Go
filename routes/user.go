package routes

import (
	"errors"

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

func GetUsers(c *fiber.Ctx) error{
	users := []models.User{}

	database.Database.Db.Find(&users)
	responseUsers := []UserSerializer{}

	for _, user := range users{
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser) 
		}
		return c.Status(200).JSON(responseUsers)

}

func FindUser(id int, user *models.User) error{
	database.Database.Db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return errors.New("user dont exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error{

	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil{
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindUser(id, & user); err != nil {
		return c.Status(400).JSON(err.Error())

	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)

}

func UpdateUser(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil{
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindUser(id, & user); err != nil {
		return c.Status(400).JSON(err.Error())

	} 

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`

	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil{
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindUser(id, & user); err != nil {
		return c.Status(400).JSON(err.Error())

	} 

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted User")

}
