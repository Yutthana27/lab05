package model

import (

    "gorm.io/gorm"
)


type Product struct {
	gorm.Model
	Name        string  `valid:"required~Name is required"`
	Price       float64 `valid:"required~Price is required,range(1|1000000)~Price must be greater than 0"`
	Stock       int     `valid:"required~Stock is required,range(0|1000000)~Stock cannot be nagative"`
	Description string  `valid:"required~Description is required, stringlength(10|1000000)~Description must be at least 10 characters"`
}