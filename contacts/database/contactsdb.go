package database

import (
	"contacts/models"

	"gorm.io/gorm"
)

type ContactDB struct {
	Client interface{}
}

func (c *ContactDB) Create(contact *models.Contact) (interface{}, error) {
	c.Client.(*gorm.DB).AutoMigrate(&models.Contact{})

	result := c.Client.(*gorm.DB).Create(contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return contact.ID, nil
}

func (c *ContactDB) Delete(id uint) (interface{}, error) {
	result := c.Client.(*gorm.DB).Delete(&models.Contact{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.RowsAffected, nil
}
