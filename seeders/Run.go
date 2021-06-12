package seeders

import (
	"fmt"
	"github.com/litondev/gin-gionic-mangament-library/configs"
	"github.com/litondev/gin-gionic-mangament-library/models"
	"strconv"
)

func RunSeed() {
	fmt.Println("Seeder Start =====>")
	fmt.Println("")

	DB := configs.Connection()	
	fmt.Println("")

	for i:=0; i<10; i++ {		
		guest_book := models.GuestBook{}

		guest_book.Description = "Litoninot"
			guest_book.User = models.User{}
			guest_book.User.Username = "Nama" + strconv.Itoa(i)
			guest_book.User.Email = "litoninot" + strconv.Itoa(i) + "@gmail.com"
			guest_book.User.Password = "12345678"

		err := DB.Create(&guest_book).Error
	
		if err != nil {
			break
		}	
	}
  
  	for i:=0; i<10; i++ {
  		book := models.Book{}

  		book.Name = "Book"

  		err := DB.Create(&book).Error

  		if err != nil {  	
  			break
  		}
  	}

 	fmt.Println("")
	fmt.Println("<===== Seeder Compelete")
}