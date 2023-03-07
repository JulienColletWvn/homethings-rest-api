package handler

import (
	"context"
	"database/sql"
	db "server/db/sqlc"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

type Data struct {
	DataTypeID int32   `json:"data_type_id"`
	Value      float64 `json:"value"`
}

func CreateData(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	data := Data{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	_, creationError := queries.CreateData(ctx, db.CreateDataParams{
		DataTypeID: sql.NullInt32{
			Int32: data.DataTypeID,
			Valid: true,
		},
		Value: data.Value,
	})

	if creationError != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)

}

func GetDatas(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	deviceId := c.AllParams()["deviceId"]
	dataTypeKey := c.AllParams()["dataTypeKey"]

	if deviceId == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	d, err := queries.GetDeviceDatas(ctx, db.GetDeviceDatasParams{
		DeviceID: sql.NullString{
			String: deviceId,
			Valid:  true,
		},
		Key: dataTypeKey,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(d)
}
