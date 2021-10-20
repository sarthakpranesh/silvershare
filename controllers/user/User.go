package user

import (
	"Github/sarthakpranesh/silvershare/connections"
	keyControllers "Github/sarthakpranesh/silvershare/controllers/key"
	"fmt"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Keys  []keyControllers.Key
}

func CreateUser(user *User) error {
	db, _ := connections.PostgresConnector()
	db.AutoMigrate(&User{})
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(user)
	return nil
}
