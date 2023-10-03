package web

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Devil666face/goaccountant/assets"

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
	path, err := setPath(assets.DirMedia)
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
func setPath(dir string) (string, error) {
	baseDir, err := os.Getwd()
	if err != nil {
		return dir, err
	}
	pathToDir, err := filepath.Abs(filepath.Join(baseDir, dir))
	if err != nil {
		return dir, err
	}
	if _, err := os.Stat(pathToDir); os.IsNotExist(err) {
		os.MkdirAll(pathToDir, os.ModePerm)
	}
	return pathToDir, nil
}
