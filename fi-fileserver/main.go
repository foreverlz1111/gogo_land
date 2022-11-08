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
	"os"
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
		//Browse: true,
		//Download: true,
	})
	// http://localhost:3000/hello.txt
}
func _init_view(app *fiber.App) {
	//绑定数据至图层
	data.Index(app)

}
func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	logfile := _init_log(app)
	defer logfile.Close()

	_init_dir(app)
	_init_view(app)
	log.Fatal(app.Listen(":3000"))
}

//header-exi-shooter
