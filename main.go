package main

import (
	"fmt"
	"github.com/go-playground/validator"
)

type Student struct {
	Name      string     `validate:"required"`
}

func main() {
	bob := &Student{
	}

	appValidator := validator.New()

	err := appValidator.Struct(bob)
	if err != nil{

		validationErrors := err.(validator.ValidationErrors)
		fmt.Println("validationErrors",validationErrors)
	}
	print("done")
}
