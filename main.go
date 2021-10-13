package main

import (
	"fmt"
	"log"
	// "os"

	"github.com/RohitKuwar/go_fiber/config"
	"github.com/RohitKuwar/go_fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	routes.Setup(app)

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// port := os.Getenv("PORT")

	fmt.Println("Server is runnig on port from config:", config.ServerPort)
	// fmt.Println("Server is runnig on port from dotenv:", port)

	app.Listen(":" + config.ServerPort)

}