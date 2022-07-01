package main

import (
	"contacts/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	contact := &models.Contact{ID: 101, Name: "Jiten", Address: "Bangalore", Email: "JitenP@Outlook.Com", ContactNo: "9618558500", Status: "active", LastModified: time.Now().GoString()}
	fmt.Println(contact)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/contact", GetContact(contact))

	r.Run(":50080")
}

// func GetContactfunc(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "pong",
// 	})
// }

func GetContact(contact *models.Contact) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, contact)
	}
}
