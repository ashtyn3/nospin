package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	args "nospin/arg-parser"
	"nospin/auth"
	"nospin/config"
	"nospin/file"
	"nospin/server"
	"nospin/share"
	"nospin/user"
	"nospin/util"
	"os"
	us "os/user"
	"strings"

	zi "github.com/ashtyn3/zi/pkg"
)

func main() {
	auth.Auth()
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
		fmt.Println("    -serve, -S                     Starts web interface server on port 3000.")

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
				file.Set(v.Param, o.Param, file.Ops{})
			} else {
				file.Set(v.Param, "", file.Ops{})
			}
		} else if v.Flag == "-o" || v.Flag == "-add" || v.Flag == "-remove" {
			v.Flag += ""
		} else if v.Flag == "-del" || v.Flag == "-D" {
			file.Del(v.Param)
		} else if v.Flag == "-get" || v.Flag == "-g" {
			r := file.Get(v.Param)
			// fmt.Println(r)
			// if r.Image == true {
			d, _ := base64.StdEncoding.DecodeString(string(r.Content))
			fmt.Println(string(d))

			// string(r.Content)
			// } else {
			// fmt.Println(string(r.Content))
			// }
		} else if v.Flag == "-user" || v.Flag == "-u" {
			z, err := zi.Zi(auth.Auth().Url, auth.Auth().Pd)
			if err != nil {
				fmt.Println(err)
			}
			var u user.User
			raw := z.Get(config.Get("name"))
			json.Unmarshal([]byte(raw.Value), &u)
			if u.Name == "" {
				fmt.Println("Could not find your account.")
				os.Exit(0)
			}
			fmt.Println("Name: " + u.Name + "\nPublic Token: " + u.PubTok)

		} else if v.Flag == "-share" || v.Flag == "-s" {
			place, ok := util.FindParam(Args, "-add")
			rPlace, okRemove := util.FindParam(Args, "-remove")
			if ok == true {
				o := Args[place]
				for _, email := range strings.Split(o.Param, ",") {
					share.Add(email, v.Param)
				}
			} else if okRemove == true {
				o := Args[rPlace]
				for _, email := range strings.Split(o.Param, ",") {
					share.Remove(email, v.Param)
				}
			} else {
				fmt.Println("Needs name of person to share with. Use -p to state that.")
			}
		} else if v.Flag == "-serve" || v.Flag == "-S" {
			server.Run()
		} else {
			fmt.Println("unknown flag " + v.Flag)
			os.Exit(0)
		}

	}

}
