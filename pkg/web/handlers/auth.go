package handlers

import (
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/Devil666face/goaccountant/pkg/web/models"

	"github.com/gofiber/fiber/v2"
)

var ErrInSession = fiber.ErrInternalServerError

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
	if err := u.LoginValidate(uof.Database(), in.Password); err != nil {
		return uof.ViewCtx().RenderWithCtx("login", fiber.Map{
			"Title":   "Login",
			"Message": err.Error(),
		}, "base")
	}
	// if uof.GetSession() != nil {
	// 	return fiber.ErrInternalServerError
	// }
	if err := uof.SetInSession(web.AuthKey, true); err != nil {
		return ErrInSession
	}
	if err := uof.SetInSession(web.UserID, u.ID); err != nil {
		return ErrInSession
	}
	// if err := uof.SaveSession(); err != nil {
	// 	return fiber.ErrInternalServerError
	// }

	// Add Hx-Redirect to index page
	return uof.ViewCtx().RedirectToRoute("index", fiber.Map{})
}
