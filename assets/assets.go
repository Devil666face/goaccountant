package assets

import (
	"embed"
)

const (
	DirMedia  = "media"
	DirStatic = "static"
)

//go:embed templates/*
var ViewFS embed.FS

//go:embed static/*
var StaticFS embed.FS
