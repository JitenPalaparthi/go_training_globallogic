package models

import (
	"encoding/json"
	"fmt"
)

type Contact struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name" gorm:"index"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	ContactNo    string `json:"contactNo"`
	Status       string `json:"status"`
	LastModified string `json:"lastModifed"`
}

func (c *Contact) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("invalid data for the filed:Name")
	}
	if c.Email == "" {
		return fmt.Errorf("invalid data for the filed:Email")
	}
	if c.ContactNo == "" {
		return fmt.Errorf("invalid data for the filed:ContactNo")
	}
	return nil
}

func (c *Contact) ToJsonByte() ([]byte, error) {
	return json.Marshal(c)
}
