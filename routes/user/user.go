package userroutes

import (
	"github.com/gin-gonic/gin"
	requestmodels "github.com/ldehner/fiber-rental-api/models/request"
	storemodels "github.com/ldehner/fiber-rental-api/models/store"
)

func CreateUser(c *gin.Context) {
	var user requestmodels.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}
	storeUser := CreateStoreUser(user)
	// do something with storeUser
	returnUser := CreateResponseUser(storeUser)
	c.JSON(200, gin.H{
		"user": returnUser,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	// get user from db
	storeUser := storemodels.User{}
	storeUser.Id = id
	returnUser := CreateResponseUser(storeUser)
	c.JSON(200, gin.H{
		"user": returnUser,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user requestmodels.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "user " + id,
		})
	}
	// update User in db
	storeUser := CreateStoreUser(user)
	returnUser := CreateResponseUser(storeUser)
	c.JSON(200, gin.H{
		"user": returnUser,
	})
}
