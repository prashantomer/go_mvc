package initializers

import (
	"fmt"
	"go_mvc/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
  var err error

  dsn := os.Getenv("DB_URL")
  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    fmt.Println("Failed to connect to Database!")
  }
}

// Migrate DB
func SyncDB()  {
  // Migrate the schema
  DB.AutoMigrate(&models.Post{})
}