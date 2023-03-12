package handler

import (
	"context"
	"database/sql"
	db "server/db/sqlc"
	transform "server/transform"
	"server/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

type queryError struct {
	Message string
	Err     error
}

func (q *queryError) Error() string {
	return q.Message
}

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

func getDatas(c *fiber.Ctx) (d []db.GetDatasRow, qError error) {
	ctx := context.Background()
	queries := db.New(utils.Database)
	deviceId := c.AllParams()["deviceId"]
	dataTypeKey := c.AllParams()["dataTypeKey"]
	fromDate := c.Query("fromDate")
	toDate := c.Query("toDate")

	if deviceId == "" || dataTypeKey == "" || fromDate == "" || toDate == "" {
		return nil, &queryError{
			Message: "Missing parameter",
			Err:     nil,
		}
	}

	fromTime, err := time.Parse("2006-01-02", fromDate)

	if err != nil {
		return nil, &queryError{
			Message: "Cannot parse from date",
			Err:     err,
		}
	}

	toTime, err := time.Parse("2006-01-02", toDate)

	if err != nil {
		return nil, &queryError{
			Message: "Cannot parse to date",
			Err:     err,
		}
	}

	d, queryErr := queries.GetDatas(ctx, db.GetDatasParams{
		DeviceID: sql.NullString{
			String: deviceId,
			Valid:  true,
		},
		Key: dataTypeKey,
		CreatedAt: sql.NullTime{
			Time:  fromTime,
			Valid: true,
		},
		CreatedAt_2: sql.NullTime{
			Time:  toTime,
			Valid: true,
		},
	})

	if queryErr != nil {
		return nil, &queryError{
			Message: "Cannot get datas from DB",
			Err:     queryErr,
		}
	}

	return d, nil
}

func GetLastData(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	deviceId := c.AllParams()["deviceId"]
	dataTypeKey := c.AllParams()["dataTypeKey"]

	if deviceId == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	d, err := queries.GetLastData(ctx, db.GetLastDataParams{
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

func GetAllDatas(c *fiber.Ctx) error {
	if datas, err := getDatas(c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else {
		return c.Status(fiber.StatusOK).JSON(datas)
	}
}

func GetHourlyDatas(c *fiber.Ctx) error {
	if datas, err := getDatas(c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else {
		return c.Status(fiber.StatusOK).JSON(transform.GetHourlyAverageData(datas))
	}
}

func GetDailyDatas(c *fiber.Ctx) error {
	if datas, err := getDatas(c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else {
		return c.Status(fiber.StatusOK).JSON(transform.GetDailyAverageData(datas))
	}
}

func GetMonthlyDatas(c *fiber.Ctx) error {
	if datas, err := getDatas(c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else {
		return c.Status(fiber.StatusOK).JSON(transform.GetMonthlyAverageData(datas))
	}
}

func GetYearlyDatas(c *fiber.Ctx) error {
	if datas, err := getDatas(c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	} else {
		return c.Status(fiber.StatusOK).JSON(transform.GetYearlyAverageData(datas))
	}
}
