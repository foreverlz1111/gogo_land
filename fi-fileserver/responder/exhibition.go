package responder

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"main.go/data/structure"
	"os"
	"strconv"
)

func bmgt(int642 int64) string {
	switch {
	default:
		return ""
	case int642 < int64(1024):
		return strconv.FormatInt(int64(int642), 10) + "B"
	case int642 < int64(1024*1024):
		return strconv.FormatInt(int64(int642/1024), 10) + "KB"
	case int642 < int64(1024*1024*1024):
		return strconv.FormatInt(int64(int642/1024/1024), 10) + "MB"
	case int642 < int64(1024*1024*1024*1024):
		return strconv.FormatInt(int64(int642/1024/1024/1024), 10) + "GB"
	}
}
func ListDir(c *fiber.Ctx) error {
	log.Println("from responder.ListDir:")
	if structure.MyDirEntry.Pointer == "" {
		structure.MyDirEntry.Pointer, _ = os.Getwd()
		structure.MyDirEntry.Pointer += "/files/"
		structure.MyDirEntry.Cur = "/"
		dir, err := os.ReadDir(structure.MyDirEntry.Pointer)
		if err != nil {
			return err
		}
		structure.MyDirEntry.Files = nil
		for _, v := range dir {
			t, _ := v.Info()
			file := structure.File{v.Name(), v.IsDir(), bmgt(t.Size()), t.ModTime().Format("2006年01月02日 15:04:05")}
			structure.MyDirEntry.Files = append(structure.MyDirEntry.Files, file)
		}
		log.Println(structure.MyDirEntry)
		log.Println("structure.MyDirEntry.Pointer", structure.MyDirEntry.Pointer)
		log.Println("from responder.ListDir [end]")
		c.Bind(fiber.Map{
			"dir": structure.MyDirEntry,
		})
		c.Next()
	}
	return c.Render("index", fiber.Map{})
}
func Update_exhibition(c *fiber.Ctx) error {
	log.Println("from Update_exhibition")
	var body = c.Body()
	var s interface{}
	err := json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal(err)
		return err
	}
	m := s.(map[string]interface{})
	structure.MyDirEntry.Cur = fmt.Sprintf("%v", m["Cur"])
	Next := fmt.Sprintf("%v", m["Name"]) + "/"
	structure.MyDirEntry.Cur += Next
	structure.MyDirEntry.Pointer += Next
	log.Println("structure.MyDirEntry.Cur", structure.MyDirEntry.Cur)
	dir, err := os.ReadDir(structure.MyDirEntry.Pointer)
	if err != nil {
		return c.Status(200).JSON(err)
	}
	structure.MyDirEntry.Files = nil
	for _, v := range dir {
		t, _ := v.Info()
		file := structure.File{v.Name(), v.IsDir(), bmgt(t.Size()), t.ModTime().Format("2006年01月02日 15:04:05")}
		structure.MyDirEntry.Files = append(structure.MyDirEntry.Files, file)
	}
	log.Println(structure.MyDirEntry)
	log.Println("structure.MyDirEntry.Pointer", structure.MyDirEntry.Pointer)
	c.Bind(fiber.Map{
		"dir": structure.MyDirEntry,
	})
	c.Next()
	log.Println("from Update_exhibition [end]")
	return c.Status(200).JSON(nil)
}
