package auth

import (
	"os"

	"github.com/joho/godotenv"
	"qoute/config"
)

type Credentials struct {
	Url string
	Pd  string
}
func Auth() Credentials {
	godotenv.Load("../.env")
	url := config.Get("url")	
	if url == "" {
		// @zap var: url
		var Url string = os.Getenv("url")

		// @zap var: pd
		var Pd string = os.Getenv("pd")
		return Credentials{Url: Url, Pd: Pd}
	} else if url != "" {
		url := config.Get("url")
		pd := config.Get("pd")
		return Credentials{Url: url, Pd: pd}
	}
}

var _ string = os.Getenv("url")
