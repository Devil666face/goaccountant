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

func (c ViewCtx) RenderWithCtx(name string, bind fiber.Map, layouts ...string) error {
	bind["c"] = c
	return c.Render(name, bind, layouts...)
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

func (c ViewCtx) URL(name string) string {
	return c.getRouteURL(name, fiber.Map{})
}

func (c ViewCtx) URLto(name, key string, val any) string {
	return c.getRouteURL(name, fiber.Map{
		key: val,
	})
}

func (c ViewCtx) getRouteURL(name string, fmap fiber.Map) string {
	url, err := c.GetRouteURL(name, fmap)
	if err != nil {
		log.Printf("Url - %s not found", name)
	}
	return url
}
