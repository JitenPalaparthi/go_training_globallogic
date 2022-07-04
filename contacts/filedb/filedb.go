package filedb

import (
	"contacts/models"
	"fmt"
	"os"
)

//IfExists(email string) error
//Get(string) (*models.Contact, error)
//Create(*models.Contact) (interface{}, error)
//Delete(string) (interface{}, error)

type FileDB struct{}

func (c *FileDB) IfExists(email string) error {
	return nil
}

func (fd *FileDB) Create(contact *models.Contact) (interface{}, error) {
	var f *os.File

	f, err := os.Open("contacts.fdb")
	if err != nil {
		f, err = os.Create("contacts.fdb")
		if err != nil {
			return nil, err
		}
	} else {
		defer f.Close()
	}
	return f.WriteString(fmt.Sprintln(contact))
}

func (c *FileDB) Get(id string) (*models.Contact, error) {
	return nil, nil
}

func (c *FileDB) Delete(id string) (interface{}, error) {
	return nil, nil
}
