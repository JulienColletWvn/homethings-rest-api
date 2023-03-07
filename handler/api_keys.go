package handler

import (
	"context"
	db "server/db/sqlc"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

func GetApiKey(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	apiKey := c.AllParams()["apiKey"]

	if apiKey == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	k, err := queries.GetApiKey(ctx, apiKey)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(k)

}
