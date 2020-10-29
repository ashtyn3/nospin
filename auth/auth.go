package auth

import (
	"os"

	"github.com/joho/godotenv"
)

func Auth() {
	godotenv.Load("../.env")
}

var _ string = os.Getenv("url")

// @zap var: url
var Url string = os.Getenv("url")

// @zap var: pd
var Pd string = os.Getenv("pd")
