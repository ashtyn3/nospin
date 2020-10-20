package main

import (
	"fmt"
	"io/ioutil"
	"log"
	args "nospin/arg-parser"
	"nospin/config"
	"nospin/user"
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
		}

	}

}
