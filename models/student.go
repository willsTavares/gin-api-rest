package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

// Student struct
type Student struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11"`
	RG   string `json:"rg" validate:"len=9"`
}

var Students []Student

func ValidateStudent(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
