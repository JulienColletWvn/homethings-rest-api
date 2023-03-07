package handler

import (
	"context"
	"database/sql"
	db "server/db/sqlc"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

type DataType struct {
	Key  string `json:"key"`
	Unit string `json:"unit"`
}

func CreateDataType(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	dataType := DataType{}
	deviceId := c.AllParams()["deviceId"]

	if err := c.BodyParser(&dataType); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	_, creationError := queries.CreateDataType(ctx, db.CreateDataTypeParams{
		Key:  dataType.Key,
		Unit: dataType.Unit,
		DeviceID: sql.NullString{
			String: deviceId,
			Valid:  true,
		},
	})

	if creationError != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusCreated)

}

func GetDataTypes(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	deviceId := c.AllParams()["deviceId"]

	dts, err := queries.GetDeviceDataTypes(ctx, sql.NullString{
		String: deviceId,
		Valid:  true,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(dts)
}
