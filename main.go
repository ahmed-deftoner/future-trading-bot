package main

import (
	"log"
	"os"

	"github.com/ahmed-deftoner/future-trading-bot/controllers"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error getting env")
	} else {
		log.Println("Getting Values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	server.Run(":8080")
}

func main() {
	Run()
}
