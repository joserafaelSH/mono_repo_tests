package main

import (
	"fmt"

	"github.com/joserafaelSH/test_go/pkg/db"
	"github.com/joserafaelSH/test_go/pkg/models"
)

func main() {
	DB := db.Init()

	var user []models.User

	if result := DB.Find(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println(user)
}
