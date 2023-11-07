package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/Devil666face/goaccountant/pkg/web/models"

	"github.com/gofiber/fiber/v2"
)

func UserListPage(uof *web.Uof) error {
	if uof.ViewCtx().IsHtmx() && uof.ViewCtx().IsHtmxCurrentURL() {
		return uof.ViewCtx().RenderWithCtx("user_content", fiber.Map{
			"Users": models.GetAllUsers(uof.Database()),
		})
	}
	return uof.ViewCtx().RenderWithCtx("user_list", fiber.Map{
		"Title": "List of users",
		"Users": models.GetAllUsers(uof.Database()),
	}, "base")
}

func UserEditForm(uof *web.Uof) error {
	u := models.User{}
	id, err := strconv.Atoi(uof.ViewCtx().Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	u.ID = uint(id)
	if err := u.Get(uof.Database()); err != nil {
		return fiber.ErrNotFound
	}
	u.Password = ""
	return uof.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
		web.UserKey: u,
	})
}

func UserCreateForm(uof *web.Uof) error {
	return uof.ViewCtx().RenderWithCtx("user_create", fiber.Map{})
}

func UserCreate(uof *web.Uof) error {
	u := models.User{}
	if err := uof.ViewCtx().BodyParser(&u); err != nil {
		return fiber.ErrBadRequest
	}
	if err := u.Validate(); err != nil {
		return uof.ViewCtx().RenderWithCtx("user_create", fiber.Map{
			"Message": err.Error(),
		})
	}
	if err := u.Create(uof.Database()); err != nil {
		return uof.ViewCtx().RenderWithCtx("user_create", fiber.Map{
			"Message": err.Error(),
		})
	}
	return uof.ViewCtx().RenderWithCtx("user_create", fiber.Map{
		"Success": fmt.Sprintf("User %s - created", u.Username),
	})
}

func UserEdit(uof *web.Uof) error {
	var (
		u  = models.User{}
		in = models.User{}
	)
	if err := uof.ViewCtx().BodyParser(&in); err != nil {
		return err
	}
	id, err := strconv.Atoi(uof.ViewCtx().Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	in.ID = uint(id)
	u.ID = in.ID
	if err := u.Get(uof.Database()); err != nil {
		return fiber.ErrNotFound
	}
	if err := in.Validate(); err != nil {
		if errors.Is(err, models.ErrPasswordRequired) {
			in.Password, in.PasswordConfirm = u.Password, u.Password
		} else {
			return uof.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
				web.UserKey: u,
			})
		}
	}
	u.Username, u.Admin, u.Password = in.Username, in.Admin, in.Password
	if err := u.Update(uof.Database()); err != nil {
		return uof.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
			web.UserKey: u,
			"Message":   err,
		})
	}
	if err := uof.DestroySessionByID(u.SessionKey); err != nil {
		return uof.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
			web.UserKey: u,
			"Message":   err,
			// "Message":   ErrInSession,
		})
	}
	return uof.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
		web.UserKey: u,
		"Success":   "Successful update user",
	})
}

func UserDelete(uof *web.Uof) error {
	u := models.User{}
	if err := uof.ViewCtx().BodyParser(&u); err != nil {
		return err
	}
	id, err := strconv.Atoi(uof.ViewCtx().Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	u.ID = uint(id)
	if err := u.Get(uof.Database()); err != nil {
		return fiber.ErrNotFound
	}
	if err := u.Delete(uof.Database()); err != nil {
		return err
	}
	return uof.ViewCtx().RenderWithCtx("user_content", fiber.Map{
		"Users": models.GetAllUsers(uof.Database()),
	})
}
