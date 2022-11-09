package data

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"main.go/data/structure"
)

func Index(c *fiber.Ctx) error {
	log.Println("from data.Index")
	err := c.Bind(fiber.Map{
		"header": structure.Myheader,
		"dir":    structure.MyDirEntry,
		"footer": structure.Myfooter,
	})
	if err != nil {
		return err
	}
	return c.Next()
}
