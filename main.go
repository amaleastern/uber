package main

import (
	"log"
	"os"
	"uber/src/db"
	"uber/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbHandler := db.Init()
	db.Migrate(dbHandler)

	app := fiber.New()
	routes.Setup(app, dbHandler)

	address := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	log.Println("listening on " + address + " ok. ")
	log.Fatal(app.Listen(address))
}
