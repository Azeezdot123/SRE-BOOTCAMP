package models

import (
	"errors"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	FirstName   string 
	LastName    string 
	Age         int    
	Sex 	   string
}

func (s *Student) Validate() error {
	if s.FirstName == "" && s.LastName == "" {
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