package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/pharmacity-xyz/server/database"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(struct {
		Age  int
		Name string
	})

	if err := c.BodyParser(&user); err != nil {
		c.Status(400).JSON("Error Parsing Input")
		return err
	}

	createdUser, err := database.EntClient.User.Create().SetAge(user.Age).SetName(user.Name).Save(context.Background())

	if err != nil {
		c.Status(500).JSON("Unable to save user")
		return err
	}

	c.Status(200).JSON(createdUser)
	return nil
}
