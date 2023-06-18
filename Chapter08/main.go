package main

import (
	"log"

	"github.com/ws/chapter08/module"
	"github.com/ws/chapter08/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	go module.RunHub()
	site := fiber.New()
	url.Web(site)
	log.Fatal(site.Listen(":3000"))
}
