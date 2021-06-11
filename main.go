package main

import (
	"fmt"
	"github.com/litondev/gin-gionic-mangament-library/configs"
	"github.com/litondev/gin-gionic-mangament-library/models"
)

func main(){
	DB := configs.Connection()

	fmt.Println(DB.Migrator().CurrentDatabase())

	fmt.Println("Hello World")

	DB.Migrator().DropTable(&models.User{})

	DB.AutoMigrate(&models.User{})
}