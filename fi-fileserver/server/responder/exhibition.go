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
	poi, _ := os.Getwd()
	poi += "/files/"
	cur := "/"
	direntry := Readdir(cur, poi)
	log.Println(direntry)
	//dirs := structure.DirEntry{}
	//file := structure.File{"发表的小论文", true, "4KB", "2022年11月10日 18:59:44"}
	//dirs.Files = append(dirs.Files, file)
	//dirs.Pointer = "/home/sakurinn/project/go/fi-fileserver/files/"
	//dirs.Cur = "/"
	log.Println("from responder.ListDir [end]")
	return c.Status(200).JSON(direntry)
}

func PreviousDir(c *fiber.Ctx) error {
	type Parser struct {
		Poi  string `json:"Poi"`
		Cur  string `json:"Cur"`
		Name string `json:"Name"`
	}
	parser := Parser{}
	err := c.BodyParser(&parser)
	if err != nil {
		log.Fatal(err)
	}
	if filepath.Dir(filepath.Clean(parser.Cur)) != "/" {
		parser.Cur = filepath.Dir(filepath.Clean(parser.Cur)) + "/"
		parser.Poi = filepath.Dir(filepath.Clean(parser.Poi)) + "/"
	} else {
		parser.Cur = "/"
		parser.Poi, _ = os.Getwd()
		parser.Poi += "/files/"
	}
	direntry := Readdir(parser.Cur, parser.Poi)
	return c.Status(200).JSON(direntry)
}

func IsRoot(pointer string) bool {
	temp, _ := os.Getwd()
	if pointer == temp+"/files/" {
		return true
	}
	return false
}

func Readdir(cur string, poi string) structure.DirEntry {
	dir, err := os.ReadDir(poi)
	if err != nil {
		log.Fatal(err)
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
	return direntry
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
	direntry := Readdir(parser.Cur, parser.Poi)
	c.Render("index", fiber.Map{"dir": direntry})

	log.Println("from responder.UpdateExhibition [end]")
	return c.Status(200).JSON(direntry)
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
