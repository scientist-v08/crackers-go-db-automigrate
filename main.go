package main

import (
	"log"

	"github.com/scientist-v08/crackers-db-automigrate/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	log.Println("Database migration completed successfully")
}
