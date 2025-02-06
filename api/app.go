package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Handler untuk menjalankan aplikasi Fiber
func Handler(w http.ResponseWriter, r *http.Request) {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Vercel!")
    })

    // Menjalankan Fiber aplikasi pada HTTP handler
    app.Handler()
}

// Main function harus tetap ada untuk Vercel
func main() {
    http.HandleFunc("/", Handler)
    http.ListenAndServe(":3000", nil)
}
