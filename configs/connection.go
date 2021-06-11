package configs

import (
  "github.com/joho/godotenv"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "fmt"
  "os"
)

func Connection() (*gorm.DB) {
  err := godotenv.Load()
 
  if err != nil {
	panic("failed to load .env file")
  }
 
  db, err := gorm.Open(mysql.Open(os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  fmt.Println("success to connect database")

  return db
}