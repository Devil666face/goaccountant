package handlers

import (
	"fmt"

	"github.com/Devil666face/goaccountant/pkg/web/models"
	"github.com/gofiber/fiber/v2"
)

func UserList(c *fiber.Ctx) error {
	return c.Render("user_list", fiber.Map{
		"c": ViewCtx{c},
	}, "base")
}

func UserCreateForm(c *fiber.Ctx) error {
	return c.Render("user_create", fiber.Map{
		"c": ViewCtx{c},
	})
}

func UserCreate(c *fiber.Ctx) error {
	u := models.User{}
	if err := c.BodyParser(&u); err != nil {
		return err
	}
	fmt.Println(u)
	return nil
}
