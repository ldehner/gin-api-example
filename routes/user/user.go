package userroutes

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	requestmodels "github.com/ldehner/fiber-rental-api/models/request"
)

func CreateUser(c *gin.Context) {
	return
}

func GetUser(c *gin.Context) {
	id := c.Params("id")
	user, err := conf.Conf{}.GetUserRepository().GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	user.Id = id
	return c.Status(fiber.StatusFound).JSON(CreateResponseUser(user))
}

func UpdateUser(c *gin.Context) {
	var user requestmodels.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	dbuser, err := conf.Conf{}.GetUserRepository().UpdateUser(CreateStoreUser(user))
	if err != nil {
		return c.Status(fiber.StatusConflict).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(CreateResponseUser(dbuser))
}
