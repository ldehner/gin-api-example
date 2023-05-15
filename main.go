package main

import (
	"github.com/gin-gonic/gin"
	userroutes "github.com/ldehner/fiber-rental-api/routes/user"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080

}

func setupRoutes(r *gin.Engine) {
	// user endpoints
	r.GET("/user/:id", userroutes.GetUser)
}
