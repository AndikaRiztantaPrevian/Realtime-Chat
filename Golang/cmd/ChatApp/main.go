package main

import (
	"github.com/AndikaRiztantaPrevian/ChatApp/internal/models"
	"github.com/AndikaRiztantaPrevian/ChatApp/internal/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost port=5432 user=postgres dbname=chat_app sslmode=disable password=Andika7123"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
        panic("failed to connect to database")
    }

	models.Migrate(db)

	router := routes.NewRouter(db)
	router.Run(":8080")
}