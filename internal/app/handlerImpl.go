package app

import (
	"embed"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

var (
	viewfs embed.FS
)

func handleContext(argsPtr *App) {
	var app = fiber.New()

	app.Get("/*", static.New("./public"))

	err := app.Listen(argsPtr.address + ":" + argsPtr.port)
	if err != nil {
		println(err.Error())
		return
	}
}

func (appPtr *App) Instance() interface{} {
	var args = App{
		address: "localhost",
		port:    "8080",
	}

	*appPtr = args
	handleContext(appPtr)
	return nil
}
