package main

import (
	"math/rand"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}


type Account struct {
	gorm.Model
	ID int           
	FirstName string`json:"firstName"`
	LastName string`json:"lastName"`
	Balance int64`json:"balance"`
	Number int64`json:"number"`
}

func NewAccount(firstName , lastName string) *Account{
	return &Account{
		FirstName:  firstName,
		LastName: lastName,
		Number: rand.Int63n(1000000),
	}
}



