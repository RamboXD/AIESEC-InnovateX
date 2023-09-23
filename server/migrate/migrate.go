package main

import (
	"fmt"
	"log"

	"github.com/RamboXD/SRS/initializers"

	"github.com/RamboXD/SRS/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.City{})
	initializers.DB.AutoMigrate(&models.Board{})
	initializers.DB.AutoMigrate(&models.Company{})
	initializers.DB.AutoMigrate(&models.Game{})
	fmt.Println("? Migration complete")
}

