package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/RohitKuwar/go_fiber/routes"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	// app.Use(logger.New())
	routes.Setup(app)
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	app.Listen(":" + port)
	fmt.Println("Server is runnig on port: %v", port)

}
