package data

import (
	"github.com/gofiber/fiber/v2"
	"main.go/data/structure"
	"time"
)

func Index(app *fiber.App) {
	loc := time.Local
	header := structure.Header{
		Title:     "标题",
		HintTitle: "文件预览",
		Edit:      "编辑",
		Upload:    "上传",
		Date:      time.Date(2022, 11, 7, 0, 0, 0, 0, loc),
	}
	footer := structure.Footer{
		Title: "footer",
	}
	footer.Todo = append(footer.Todo,
		"日志记录格式加入日期",
		"标题显示当前目录",
		"编辑文件目录",
		"下载文件按钮",
	)
	dir := structure.DirEntry{}
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"header": header,
			"dir":    dir,
			"footer": footer,
		})
	})
}
