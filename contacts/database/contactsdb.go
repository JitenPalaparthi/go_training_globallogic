package database

import (
	"contacts/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	ERROR_CONTACT_EXISTS = errors.New("Contact already exists with the given email address")
)

type ContactDB struct {
	Client interface{}
}

func (c *ContactDB) IfExists(email string) error {

	filter := map[string]interface{}{}
	filter["email"] = email
	result := c.Client.(*gorm.DB).Model(&models.Contact{}).First(&filter)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected > 0 {
		return ERROR_CONTACT_EXISTS
	}
	return nil
}

func (c *ContactDB) Get(id string) (*models.Contact, error) {
	contact := &models.Contact{}
	result := c.Client.(*gorm.DB).First(contact, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}

func (c *ContactDB) Create(contact *models.Contact) (interface{}, error) {
	c.Client.(*gorm.DB).AutoMigrate(&models.Contact{})
	result := c.Client.(*gorm.DB).Create(contact)
	if result.Error != nil {
		fmt.Println("------------->", result.Error)
		return nil, result.Error
	}
	return contact.ID, nil
}

func (c *ContactDB) Delete(id string) (interface{}, error) {
	result := c.Client.(*gorm.DB).Delete(&models.Contact{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.RowsAffected, nil
}
