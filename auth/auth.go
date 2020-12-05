package auth

import (
	"os"

	"github.com/joho/godotenv"
)

type Credentials struct {
	Url string
	Pd  string
}

func Auth() Credentials {
	godotenv.Load("../.env")
	// @zap var: url
	var Url string = os.Getenv("url")

	// @zap var: pd
	var Pd string = os.Getenv("pd")

	return Credentials{Url: Url, Pd: Pd}
}

var _ string = os.Getenv("url")
