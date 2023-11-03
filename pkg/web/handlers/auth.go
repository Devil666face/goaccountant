package handlers

import (
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/Devil666face/goaccountant/pkg/web/models"

	"github.com/gofiber/fiber/v2"
)

func LoginPage(uof *web.Uof) error {
	return uof.ViewCtx().RenderWithCtx("login", fiber.Map{
		"Title": "Login",
	}, "base")
}

func Login(uof *web.Uof) error {
	var (
		u  = &models.User{}
		in = &models.User{}
	)
	if err := uof.ViewCtx().BodyParser(in); err != nil {
		return err
	}
	u.Username = in.Username
	if !u.IsFound(uof.Database()) {
		return uof.ViewCtx().RenderWithCtx("login", fiber.Map{
			"Title":   "Login",
			"Message": "Username or password is wrong",
		}, "base")
	}
	if !u.ComparePassword(in.Password) {
		return uof.ViewCtx().RenderWithCtx("login", fiber.Map{
			"Title":   "Login",
			"Message": "Username or password is wrong",
		}, "base")
	}
	if uof.GetSession() != nil {
		return fiber.ErrInternalServerError
	}
	uof.SetInSession(web.AuthKey, true)
	uof.SetInSession(web.UserID, u.ID)
	if err := uof.SaveSession(); err != nil {
		return fiber.ErrInternalServerError
	}
	return uof.ViewCtx().RedirectToRoute("user_list", fiber.Map{})
}
