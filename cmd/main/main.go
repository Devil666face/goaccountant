package main

import (
	"log"

	"github.com/Devil666face/goaccountant/internal/web"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	wa := web.New()
	if err := wa.Listen(); err != nil {
		log.Fatalln(err)
	}
}
