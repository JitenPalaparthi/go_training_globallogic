package interfaces

import "contacts/models"

type IContact interface {
	Create(*models.Contact) (interface{}, error)
}
