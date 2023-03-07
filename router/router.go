package router

import (
	"server/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	apiKey := api.Group("/api-keys")
	apiKey.Get("/:apiKey", handler.GetApiKey)

	devices := api.Group("/devices")
	devices.Post("/", handler.CreateDevice)
	devices.Get("/", handler.GetDevices)
	devices.Get("/:deviceId", handler.GetDevice)

	dataTypes := devices.Group("/:deviceId/data-types")
	dataTypes.Post("/", handler.CreateDataType)
	dataTypes.Get("/", handler.GetDataTypes)

	datas := api.Group("/datas")
	datas.Post("/", handler.CreateData)
	datas.Get("/", handler.GetDatas)

}
