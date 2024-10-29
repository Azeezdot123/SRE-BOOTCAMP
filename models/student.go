package models

import (
	"errors"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         int `json:"age"`    
	Sex 	   string `json:"sex"`
}

func (s *Student) Validate() error {
	if s.FirstName == "" || s.LastName == "" {
		return errors.New("first name and last name is required")
	}
	if s.Age < 0 {
		return errors.New("age must be greater than 0")
	}
	if s.Sex == "" {
		return errors.New("sex is required")
	}
	return nil
}