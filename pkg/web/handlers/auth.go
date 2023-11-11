package handlers

import (
	"github.com/Devil666face/goaccountant/pkg/web"
	"github.com/Devil666face/goaccountant/pkg/web/models"

	"github.com/gofiber/fiber/v2"
)

var ErrInSession = fiber.ErrInternalServerError

func LoginPage(unit *web.Unit) error {
	return unit.ViewCtx().RenderWithCtx("login", fiber.Map{
		"Title": "Login",
	}, "base")
}

func Login(unit *web.Unit) error {
	var (
		u   = &models.User{}
		in  = &models.User{}
		err error
	)
	if err := unit.ViewCtx().BodyParser(in); err != nil {
		return err
	}
	u.Email = in.Email
	if err := u.LoginValidate(unit.Database(), unit.Validator(), in.Password); err != nil {
		return unit.ViewCtx().RenderWithCtx("login", fiber.Map{
			"Title":   "Login",
			"Message": err.Error(),
		}, "base")
	}
	if err := unit.SetInSession(web.AuthKey, true); err != nil {
		return ErrInSession
	}
	if err := unit.SetInSession(web.UserID, u.ID); err != nil {
		return ErrInSession
	}
	if u.SessionKey, err = unit.SessionID(); err != nil {
		return ErrInSession
	}
	if err := u.Update(unit.Database()); err != nil {
		return ErrInSession
	}
	return unit.ViewCtx().ClientRedirect(unit.ViewCtx().URL("index"))
}

func Logout(unit *web.Unit) error {
	if err := unit.DestroySession(); err != nil {
		return ErrInSession
	}
	return unit.ViewCtx().RedirectToRoute("login", nil)
}
