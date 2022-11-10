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
	if IsRoot(structure.MyDirEntry.Pointer) {
		return c.Render("index", fiber.Map{})
	} else if IsEmpty(structure.MyDirEntry.Pointer) {
		structure.MyDirEntry.Pointer, _ = os.Getwd()
		structure.MyDirEntry.Pointer += "/files/"
		structure.MyDirEntry.Cur = "/"
		err := Readdir(c, structure.MyDirEntry.Pointer)
		if err != nil {
			return err
		}
	}
	log.Println("from responder.ListDir [end]")
	return c.Render("index", fiber.Map{})
}

func PreviousDir(c *fiber.Ctx) error {
	var body = c.Body()
	var s interface{}
	err := json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal(err)
		return err
	}
	m := s.(map[string]interface{})
	cur := fmt.Sprintf("%v", m["Cur"])
	if cur == "/" {

	}
	return nil

}

func IsRoot(pointer string) bool {
	temp, _ := os.Getwd()
	if pointer == temp+"/files/" {
		return true
	}
	return false
}

func Readdir(c *fiber.Ctx, pointer string) error {
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
	c.Bind(fiber.Map{
		"dir": structure.MyDirEntry,
	})
	c.Next()
	return nil
}

func IsEmpty(pointer string) bool {
	if pointer == "" {
		return true
	}
	return false
}

func UpdateExhibition(c *fiber.Ctx) error {
	log.Println("from responder.UpdateExhibition")
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
	err = Readdir(c, structure.MyDirEntry.Pointer)
	if err != nil {
		return c.Status(400).JSON(err)
	}
	log.Println("from responder.UpdateExhibition [end]")
	return c.Status(200).JSON(nil)
}
