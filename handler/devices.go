package handler

import (
	"context"
	db "server/db/sqlc"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

type Device struct {
	DeviceId string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"Location"`
}

func CreateDevice(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	device := Device{}

	if err := c.BodyParser(&device); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	_, creationError := queries.CreateDevice(ctx, db.CreateDeviceParams{
		ID:       device.DeviceId,
		Name:     device.Name,
		Location: device.Location,
	})

	if creationError != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusCreated)

}

func GetDevices(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	d, err := queries.GetDevices(ctx)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(d)
}

func GetDevice(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	deviceId := c.AllParams()["deviceId"]

	if deviceId == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	d, err := queries.GetDevice(ctx, deviceId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(d)
}
