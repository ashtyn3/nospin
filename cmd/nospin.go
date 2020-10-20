package main

import (
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
)

func main() {
	usr, err := us.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := usr.HomeDir
	a := os.Args[1:]
	Args := args.ArgParser(a)
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
			fmt.Println(string(r.Content))
		} else {
			fmt.Println("unknown flag " + v.Flag)
			os.Exit(0)
		}

	}

}
