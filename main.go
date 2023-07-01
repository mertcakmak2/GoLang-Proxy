package main

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "Proxy App",
		BodyLimit:     10 * 1024 * 1024,
	})

	proxy.WithTlsConfig(&tls.Config{
		InsecureSkipVerify: true,
	})

	// Rate Limiting.
	app.Use(limiter.New(limiter.Config{
		Max:               3,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	// Grouping with Middleware
	api := app.Group("/api", middleware) // /api
	v1 := api.Group("/v1")               // /api/v1

	// http://localhost:3000/api/v1/
	v1.Get("/", helloWorld)

	// Forward to localhost:4000
	// http://localhost:3000/api/v1/file
	v1.Post("/file", proxy.Forward("http://localhost:4000/file"))
	// http://localhost:3000/api/v1/hello
	v1.Get("/hello", proxy.Forward("http://localhost:4000"))

	app.Listen(":3000")
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World from port 3000!")
}

func middleware(c *fiber.Ctx) error {
	log.Println("middleware")
	c.Request().Header.Set("user", "mert_cakmak")
	return c.Next()
}
