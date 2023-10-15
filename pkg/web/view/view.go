package view

import (
	"fmt"
	"html/template"
	"log"

	"github.com/Devil666face/goaccountant/pkg/web/middlewares"

	"github.com/gofiber/fiber/v2"
)

type Ctx struct {
	*fiber.Ctx
}

func New(c *fiber.Ctx) *Ctx {
	return &Ctx{c}
}

func (c Ctx) RenderWithCtx(name string, bind fiber.Map, layouts ...string) error {
	bind["c"] = c
	return c.Render(name, bind, layouts...)
}

func (c Ctx) Csrf() template.HTML {
	html := `<input type="hidden" name="csrf" value="%s">`
	//nolint:gosec //Because not revive data from user
	return template.HTML(fmt.Sprintf(html, c.CsrfToken()))
}

func (c Ctx) CsrfToken() string {
	if token, ok := c.Locals(middlewares.Csrf).(string); ok {
		return token
	}
	return ""
}

func (c Ctx) IsHtmx() bool {
	if htmx, ok := c.Locals(middlewares.Htmx).(bool); ok {
		return htmx
	}
	return false
}

func (c Ctx) URL(name string) string {
	return c.getRouteURL(name, fiber.Map{})
}

func (c Ctx) URLto(name, key string, val any) string {
	return c.getRouteURL(name, fiber.Map{
		key: val,
	})
}

func (c Ctx) getRouteURL(name string, fmap fiber.Map) string {
	url, err := c.GetRouteURL(name, fmap)
	if err != nil {
		log.Printf("Url - %s not found", name)
	}
	return url
}
