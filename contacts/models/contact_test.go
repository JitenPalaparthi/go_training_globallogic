package models_test

import (
	"contacts/models"
	"testing"
)

// The name of the file must ends with _test.go
// The name of the pakage must be XXXXXXX_test
// Each test method must start with TestXXXXXXXX(t *testing.T)

func TestValidateContactPass(t *testing.T) {
	contact := &models.Contact{Email: "Jitenp@outlook.com", Address: "Bangalore", ContactNo: "9618558500", Name: "Jiten"}
	err := contact.Validate()
	//time.Sleep(time.Second * 31)
	if err != nil {
		t.Error("Validate must not return an error")
		t.Fail()
	}
}

func TestValidateContactFail(t *testing.T) {
	contact := &models.Contact{Email: "Jitenp@outlook.com", Address: "Bangalore", ContactNo: "9618558500"}
	err := contact.Validate()
	if err == nil {
		t.Error("Validate must return an error")
		t.Fail()
	}
}
