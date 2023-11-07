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
		u   = &models.User{}
		in  = &models.User{}
		err error
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
	if err := uof.SetInSession(web.AuthKey, true); err != nil {
		return ErrInSession
	}
	if err := uof.SetInSession(web.UserID, u.ID); err != nil {
		return ErrInSession
	}
	if u.SessionKey, err = uof.SessionID(); err != nil {
		return ErrInSession
	}
	if err := u.Update(uof.Database()); err != nil {
		return ErrInSession
	}
	return uof.ViewCtx().ClientRedirect(uof.ViewCtx().URL("index"))
}

func Logout(uof *web.Uof) error {
	if err := uof.DestroySession(); err != nil {
		return ErrInSession
	}
	return uof.ViewCtx().RedirectToRoute("login", nil)
}
