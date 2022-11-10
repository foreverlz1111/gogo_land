package data

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"main.go/data/structure"
)

func Index(c *fiber.Ctx) error {
	log.Println("from data.Index")
	c.Bind(fiber.Map{
		"header": structure.Myheader,
		"footer": structure.Myfooter,
	})
	return c.Next()
}
