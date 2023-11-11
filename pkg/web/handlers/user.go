package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/Devil666face/goaccountant/pkg/web/models"
	"github.com/Devil666face/goaccountant/pkg/web/validators"

	"github.com/gofiber/fiber/v2"
)

func UserListPage(unit *web.Unit) error {
	// && unit.ViewCtx().IsHtmxCurrentURL()
	if unit.ViewCtx().IsHtmx() {
		return unit.ViewCtx().RenderWithCtx("user_content", fiber.Map{
			"Users": models.GetAllUsers(unit.Database()),
		})
	}
	return unit.ViewCtx().RenderWithCtx("user_list", fiber.Map{
		"Title": "List of users",
		"Users": models.GetAllUsers(unit.Database()),
	}, "base")
}

func UserEditForm(unit *web.Unit) error {
	u := models.User{}
	id, err := strconv.Atoi(unit.ViewCtx().Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	u.ID = uint(id)
	if err := u.Get(unit.Database()); err != nil {
		return fiber.ErrNotFound
	}
	u.Password = ""
	return unit.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
		web.UserKey: u,
	})
}

func UserCreateForm(unit *web.Unit) error {
	return unit.ViewCtx().RenderWithCtx("user_create", fiber.Map{})
}

func UserCreate(unit *web.Unit) error {
	u := models.User{}
	if err := unit.ViewCtx().BodyParser(&u); err != nil {
		return fiber.ErrBadRequest
	}
	if err := u.Validate(unit.Validator()); err != nil {
		return unit.ViewCtx().RenderWithCtx("user_create", fiber.Map{
			"Message": err.Error(),
		})
	}
	if err := u.Create(unit.Database()); err != nil {
		return unit.ViewCtx().RenderWithCtx("user_create", fiber.Map{
			"Message": err.Error(),
		})
	}
	return unit.ViewCtx().RenderWithCtx("user_create", fiber.Map{
		"Success": fmt.Sprintf("User %s - created", u.Email),
	})
}

func UserEdit(unit *web.Unit) error {
	var (
		u  = models.User{}
		in = models.User{}
	)
	if err := unit.ViewCtx().BodyParser(&in); err != nil {
		return err
	}
	id, err := strconv.Atoi(unit.ViewCtx().Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	in.ID = uint(id)
	u.ID = in.ID
	if err := u.Get(unit.Database()); err != nil {
		return fiber.ErrNotFound
	}
	if err := in.Validate(unit.Validator()); err != nil {
		if errors.Is(err, validators.ErrPasswordRequired) {
			in.Password, in.PasswordConfirm = u.Password, u.Password
		} else {
			return unit.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
				web.UserKey: u,
				"Message":   err,
			})
		}
	}
	u.Email, u.Admin, u.Password = in.Email, in.Admin, in.Password
	if err := u.Update(unit.Database()); err != nil {
		return unit.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
			web.UserKey: u,
			"Message":   err,
		})
	}
	if err := unit.DestroySessionByID(u.SessionKey); err != nil {
		return unit.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
			web.UserKey: u,
			"Message":   err,
			// "Message":   ErrInSession,
		})
	}
	return unit.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
		web.UserKey: u,
		"Success":   "Successful update user",
	})
}

func UserDelete(unit *web.Unit) error {
	u := models.User{}
	if err := unit.ViewCtx().BodyParser(&u); err != nil {
		return err
	}
	id, err := strconv.Atoi(unit.ViewCtx().Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	u.ID = uint(id)
	if err := u.Get(unit.Database()); err != nil {
		return fiber.ErrNotFound
	}
	if err := u.Delete(unit.Database()); err != nil {
		return err
	}
	if err := unit.DestroySessionByID(u.SessionKey); err != nil {
		return ErrInSession
	}
	return unit.ViewCtx().RenderWithCtx("user_content", fiber.Map{
		"Users": models.GetAllUsers(unit.Database()),
	})
}
