package handlers

import (
	"contacts/interfaces"
	"contacts/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type ContactHandler struct {
	IContact interfaces.IContact
}

func (ch *ContactHandler) CreateContact() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.IContact == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ContactHandler or IContact is nil")
			c.Abort()
			return
		}

		buf, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the body",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}

		contact := &models.Contact{}
		err = json.Unmarshal(buf, contact)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in body json format",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		err = contact.Validate()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			glog.Errorln(err)
			c.Abort()
			return
		}

		contact.Status = "active"
		contact.LastModified = time.Now().UTC().String()
		id, err := ch.IContact.Create(contact)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error to store in database",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": id,
		})
		glog.Info("Contact successfully created:", id)
		c.Abort()
	}
}
