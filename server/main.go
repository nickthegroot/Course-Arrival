package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)


func main() {
	log.Println("Starting SD Hacks 21' Server")

	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_CERT"),
	)
	a.Run(":8010")
}
