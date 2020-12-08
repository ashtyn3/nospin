package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"qoute/file"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

type putFile struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Content  string `json:"content"`
}

func Run() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	var ConfigDefault = cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}
	app.Use(cors.New(ConfigDefault))
	godotenv.Load("../env")
	// var path string
	// if os.Getenv("mode") == "dev" {
	// 	path = "../web"
	// } else {
	// 	path = "./Qoute/web"
	// }
	path, _ := os.UserHomeDir()

	app.Static("/", path+"/.qoute/share/web")
	app.Post("/put", func(c *fiber.Ctx) error {
		f := new(putFile)
		c.BodyParser(f)
		F, _ := os.Create(f.Name)
		F.WriteString(f.Content)
		defer F.Close()
		defer os.Remove(f.Name)
		file.Set(f.Username+"/"+f.Name, f.Name, file.Ops{})
		return c.SendStatus(202)
	})
	app.Post("/del", func(c *fiber.Ctx) error {
		f := new(putFile)
		c.BodyParser(f)
		file.Del(f.Username + "/" + f.Name)
		return c.SendStatus(202)
	})
	app.Post("/get", func(c *fiber.Ctx) error {
		f := new(putFile)
		json.Unmarshal(c.Body(), f)
		file := file.Get(f.Username + "/" + f.Name)
		if file.Image == true {
			// fmt.Println(file)
			content, _ := base64.StdEncoding.DecodeString(string(file.Content))

			// file.Content = content
			// d, _ := json.Marshal(file)
			return c.SendString(string(content))
		}
		content, _ := base64.StdEncoding.DecodeString(string(file.Content))
		return c.SendString(string(content))

	})
	fmt.Println("listening on 3000")
	app.Listen(":3000")
}
