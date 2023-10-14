package handlers

import (
	"fmt"
	"html/template"
	"log"

	"github.com/Devil666face/goaccountant/pkg/web/middlewares"
	"github.com/gofiber/fiber/v2"
)

type ViewCtx struct {
	*fiber.Ctx
}

func (c ViewCtx) Csrf() template.HTML {
	html := `<input type="hidden" name="csrf" value="%s">`
	return template.HTML(fmt.Sprintf(html, c.CsrfToken()))
}

func (c ViewCtx) CsrfToken() string {
	if token, ok := c.Locals(middlewares.Csrf).(string); ok {
		return token
	}
	return ""
}

func (c ViewCtx) IsHtmx() bool {
	return c.Locals(middlewares.Htmx).(bool)
}

func (c ViewCtx) Url(name string) string {
	return c.getRouteUrl(name, fiber.Map{})
}

func (c ViewCtx) UrlTo(name, key string, val interface{}) string {
	return c.getRouteUrl(name, fiber.Map{
		key: val,
	})
}

func (c ViewCtx) getRouteUrl(name string, fmap fiber.Map) string {
	url, err := c.GetRouteURL(name, fmap)
	if err != nil {
		log.Fatalf("Url - %s not found", name)
	}
	return url
}

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{"c": ViewCtx{c}})
}

func Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{"c": ViewCtx{c}})
}
