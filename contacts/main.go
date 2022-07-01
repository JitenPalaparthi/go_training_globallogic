package main

import (
	"contacts/database"
	"contacts/models"
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (
	contact      *models.Contact = &models.Contact{ID: 101, Name: "Jiten", Address: "Bangalore", Email: "JitenP@Outlook.Com", ContactNo: "9618558500", Status: "active", LastModified: time.Now().GoString()}
	DBCONNECTION string          = "host=localhost user=jiten password=admin dbname=sample port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	PORT         string          = ":50080"
)

func usage() {
	//fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARNING|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}

func main() {
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}

	if os.Getenv("DBCONNECTION") != "" {
		DBCONNECTION = os.Getenv("DBCONNECTION")
	}

	_, err := database.GetConnection(DBCONNECTION)

	if err != nil {
		//	panic(err)
		glog.Fatal("Database Error:", err)
	}
	//fmt.Println(db)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/contact", GetContact(contact))

	r.Run(PORT)
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
