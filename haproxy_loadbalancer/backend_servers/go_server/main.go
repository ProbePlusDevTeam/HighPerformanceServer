package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()
	port := flag.Int("port", 8080, "Port number to listen on")
	flag.Parse()
	// Send a string back for GET calls to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("PONG")
		return err
	})

	// Listen on PORT 3000
	addr := fmt.Sprintf(":%d", *port)
	app.Listen(addr)
}
