package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"jirani-api/api/controllers"
	//"jirani-api/api/seed"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	//seed.Load(server.DB)

	server.Run(":8080")

}