package controllers

import (
	"Github/sarthakpranesh/silvershare/connections"
	"fmt"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(user *User) error {
	db, _ := connections.PostgresConnector()
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(user)
	return nil
}
