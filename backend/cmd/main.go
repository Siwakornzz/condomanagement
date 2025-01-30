package main

import (
	"condomanagement/internal/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {

	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("cfg : ", cfg)

	// start app service
	app := fiber.New()

	app.Listen(":8080")

	log.Info("Server started on port 8080")
}
