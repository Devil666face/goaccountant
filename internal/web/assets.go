package web

import (
	"log"
	"net/http"

	"github.com/Devil666face/goaccountant/assets"
	"github.com/Devil666face/goaccountant/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
)

const (
	pathPrefix   = "static"
	templateSuff = ".html"
	isBrowse     = false
)

type Media struct {
	path    string
	handler fiber.Static
}

func NewStatic() func(*fiber.Ctx) error {
	return filesystem.New(filesystem.Config{
		Root:       http.FS(assets.StaticFS),
		PathPrefix: pathPrefix,
		MaxAge:     86400,
	})
}

func NewViews() *html.Engine {
	return html.NewFileSystem(http.FS(assets.ViewFS), templateSuff)
}

func NewMedia() *Media {
	path, err := utils.SetPath(assets.DirMedia)
	if err != nil {
		path = assets.DirMedia
		log.Fatalln(err)
	}
	return &Media{
		path: path,
		handler: fiber.Static{
			Compress:  true,
			ByteRange: true,
			Browse:    isBrowse,
		},
	}
}
