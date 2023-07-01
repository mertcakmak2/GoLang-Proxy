package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "File Handler App",
		BodyLimit:     10 * 1024 * 1024,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from port 4000!")
	})

	app.Post("/file", func(c *fiber.Ctx) error {
		if form, err := c.MultipartForm(); err == nil {

			files := form.File["file"]

			for _, file := range files {
				fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

				if err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename)); err != nil {
					return err
				}
				return c.Status(fiber.StatusCreated).SendString("File saved.")
			}
			return err
		}

		return c.Status(fiber.StatusInternalServerError).SendString("Error while file saving.")
	})

	app.Listen(":4000")
}
