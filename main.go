package main

import (
	"fmt"
	"github.com/game-app/entity"
	"github.com/game-app/repository/mysql"
)

func main() {
	mysqlRepo := mysql.New()
	createdUser, err := mysqlRepo.Register(entity.User{
		ID:          0,
		PhoneNumber: "0914",
		Name:        "cavid",
	})
	if err != nil {
		fmt.Errorf("Error: %w", err)
	} else {
		fmt.Println("Created user is: %v", createdUser)
	}
	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber)
	if err != nil {
		fmt.Println("unique err", err)
	}
	fmt.Println("isUnique", isUnique)
}
