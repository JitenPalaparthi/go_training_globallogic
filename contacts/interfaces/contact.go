package interfaces

import "contacts/models"

type IContact interface {
	IfExists(email string) error
	Get(string) (*models.Contact, error)
	Create(*models.Contact) (interface{}, error)
	Delete(string) (interface{}, error)
}
