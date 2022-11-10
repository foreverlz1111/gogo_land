// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

// main文件里面可以有一定的函数

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"log"
	"main.go/data"
	"main.go/responder"
	"os"
	"time"
)

func _init_log(app *fiber.App) *os.File {
	file, err := os.OpenFile("./request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("文件异常: %v", err)
	}

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} ${latency} ${method} ${host} ${path} ${ua}\n",
		Output: file,
	}))
	return file
}
func _init_dir(app *fiber.App) {
	app.Static("/", "./files", fiber.Static{
		Compress:  true,
		ByteRange: true,
		//Browse:        true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})
	// http://localhost:3000/hello.txt
}

func _init_respond(app *fiber.App) {
	//绑定路由
	app.Get("/", responder.ListDir)
	app.Post("/u", responder.UpdateExhibition)
	app.Post("/p", responder.PreviousDir)
}
func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	logfile := _init_log(app)
	defer logfile.Close()

	_init_dir(app)

	app.Use(data.Index) //绑定数据至图层

	_init_respond(app)

	log.Fatal(app.Listen(":3000"))
}

//header-exi-shooter
