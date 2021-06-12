package migrations

import (
	"fmt"
	"github.com/litondev/gin-gionic-mangament-library/configs"
	"github.com/litondev/gin-gionic-mangament-library/models"
)

func RunMigration() {
	fmt.Println("Migration Start =====>")
	fmt.Println("")

	DB := configs.Connection()
	
	fmt.Println("")

	DB.Migrator().DropTable(&models.User{})
	DB.Migrator().DropTable(&models.Book{})
	DB.Migrator().DropTable(&models.GuestBook{})

	DB.AutoMigrate(&models.User{})	
	DB.AutoMigrate(&models.Book{})
 	DB.AutoMigrate(&models.GuestBook{})

 	fmt.Println("")
	fmt.Println("<===== Migration Compelete")
}