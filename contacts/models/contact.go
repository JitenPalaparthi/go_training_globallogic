package models

import "fmt"

type Contact struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	ContactNo    string `json:"contactNo"`
	Status       string `json:"status"`
	LastModified string `json:"lastModifed"`
}

func (c *Contact) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("Invalid data for the filed:Name")
	}
	if c.Email == "" {
		return fmt.Errorf("Invalid data for the filed:Email")
	}
	if c.ContactNo == "" {
		return fmt.Errorf("Invalid data for the filed:ContactNo")
	}
	return nil
}
