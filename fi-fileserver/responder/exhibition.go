package responder

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"main.go/data/structure"
	"os"
	"path/filepath"
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
	type Parser struct {
		Poi  string `json:"Poi"`
		Cur  string `json:"Cur"`
		Name string `json:"Name"`
	}
	log.Println(c.Body())
	parser := Parser{}
	err := c.BodyParser(&parser)
	log.Println("parser err", err)
	if err != nil {
		poi, _ := os.Getwd()
		poi += "/files/"
		cur := "/"
		err = Readdir(c, cur, poi)
		if err != nil {
			return err
		}
		log.Println("from responder.ListDir [end]")
		return nil
	}
	parser.Poi += parser.Name + "/"
	parser.Cur += parser.Name + "/"
	log.Println("ListDir.parser.Poi", parser.Poi)
	err = Readdir(c, parser.Cur, parser.Poi)
	if err != nil {
		return c.Status(400).JSON(err)
	}
	log.Println("ListDir [end]")
	return c.Status(200).JSON(nil)
}

func PreviousDir(c *fiber.Ctx) error {
	if filepath.Dir(filepath.Clean(structure.MyDirEntry.Cur)) != "/" {
		structure.MyDirEntry.Cur = filepath.Dir(filepath.Clean(structure.MyDirEntry.Cur)) + "/"
		structure.MyDirEntry.Pointer = filepath.Dir(filepath.Clean(structure.MyDirEntry.Pointer)) + "/"
	} else {
		structure.MyDirEntry.Cur = "/"
		structure.MyDirEntry.Pointer, _ = os.Getwd()
		structure.MyDirEntry.Pointer += "/files/"
	}
	err := Readdir(c, "", structure.MyDirEntry.Pointer)
	if err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(200).JSON("OK")
}

func IsRoot(pointer string) bool {
	temp, _ := os.Getwd()
	if pointer == temp+"/files/" {
		return true
	}
	return false
}

func Readdir(c *fiber.Ctx, cur string, poi string) error {
	dir, err := os.ReadDir(poi)
	if err != nil {
		return err
	}
	direntry := structure.DirEntry{}
	direntry.Cur = cur
	direntry.Pointer = poi
	for _, v := range dir {
		t, _ := v.Info()
		file := structure.File{v.Name(), v.IsDir(), bmgt(t.Size()), t.ModTime().Format("2006年01月02日 15:04:05")}
		direntry.Files = append(direntry.Files, file)
	}
	log.Println("readdir:", direntry)
	return c.Render("index", fiber.Map{
		"dir": direntry,
	})
}

func IsEmpty(pointer string) bool {
	if pointer == "" {
		return true
	}
	return false
}

func UpdateExhibition(c *fiber.Ctx) error {
	log.Println("from responder.UpdateExhibition")
	type Parser struct {
		Poi  string `json:"Poi"`
		Cur  string `json:"Cur"`
		Name string `json:"Name"`
	}
	parser := Parser{}
	err := c.BodyParser(&parser)
	if err != nil {
		return c.Status(400).JSON(err)
	}
	parser.Poi += parser.Name + "/"
	parser.Cur += parser.Name + "/"
	log.Println("UpdateExhibition parser.Poi", parser.Poi)
	err = Readdir(c, parser.Cur, parser.Poi)
	if err != nil {
		return c.Status(400).JSON(err)
	}
	log.Println("from responder.UpdateExhibition [end]")
	return c.Status(200).JSON(nil)
}
func getbody() {
	//var body = c.Body()
	//var s interface{}
	//err := json.Unmarshal(body, &s)
	//if err != nil {
	//	log.Fatal(err)
	//	return err
	//}
	//m := s.(map[string]interface{})
	//structure.MyDirEntry.Cur = fmt.Sprintf("%v", m["Cur"])
}
