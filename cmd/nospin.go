package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	args "nospin/arg-parser"
	"nospin/config"
	"nospin/file"
	"nospin/user"
	"nospin/util"
	"os"
	us "os/user"

	"github.com/joho/godotenv"
	zi "github.com/vitecoin/zi/pkg"
)

func main() {
	usr, err := us.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := usr.HomeDir
	a := os.Args[1:]
	Args := args.ArgParser(a)
	if len(Args) == 0 {
		fmt.Println("Usage:")
		fmt.Println("nospin COMMAND")
		fmt.Println("commands:")
		fmt.Println("    -auth string                   Generates user with passed in email.")
		fmt.Println("    -put, -p string [-o string]    Put file with path of -o option or with passed in path.")
		fmt.Println("    -get, -g string                Get file with path.")
		fmt.Println("    -del, -D string                Delete file with path.")
		fmt.Println("    -user, -u                      Prints information about current user.")

	}
	for _, v := range Args {
		if v.Flag == "-auth" {
			_, err := ioutil.ReadDir(home + "/nospin")
			if os.IsNotExist(err) == false {
				fmt.Println("You already have an active nospin user on this device.")
			} else {
				u := user.Make(v.Param)
				config.New(u)
				fmt.Println("Created user:\n" + "Name: " + u.Name + "\nPublic Token: " + u.PubTok)
			}
		} else if v.Flag == "-put" || v.Flag == "-p" {
			place, ok := util.FindParam(Args, "-o")
			if ok == true {
				o := Args[place]
				file.Set(v.Param, o.Param)
			} else {
				file.Set(v.Param, "")
			}
		} else if v.Flag == "-o" {
			v.Flag = ""
		} else if v.Flag == "-del" || v.Flag == "-D" {
			file.Del(v.Param)
		} else if v.Flag == "-get" || v.Flag == "-g" {
			r := file.Get(v.Param)
			// fmt.Println(r)
			// if r.Image == true {
			// d, _ := base64.StdEncoding.DecodeString(string(r.Content))
			fmt.Println(string(r.Content))
			// } else {
			// fmt.Println(string(r.Content))
			// }
		} else if v.Flag == "-user" || v.Flag == "-u" {
			godotenv.Load("../.env")
			url := os.Getenv("url")
			pd := os.Getenv("pd")

			z, _ := zi.Zi(url, pd)
			var u user.User
			raw := z.Get(config.Get("name"))
			json.Unmarshal([]byte(raw.Value), &u)
			fmt.Println("Name: " + u.Name + "\nPublic Token: " + u.PubTok)

		} else {
			fmt.Println("unknown flag " + v.Flag)
			os.Exit(0)
		}

	}

}
