package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	us "qoute/user"
	"strings"
)

func New(u us.User) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := usr.HomeDir
	os.Mkdir(home+"/.qoute", 0777)
	f, _ := os.OpenFile(home+"/.qoute/config", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString("name=" + u.ID)
}

func Set(name string, value string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := usr.HomeDir
	f, _ := os.OpenFile(home+"/.qoute/config", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString("\n" + name + "=" + value)
}
func Get(name string) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := usr.HomeDir
	f, fErr := ioutil.ReadFile(home + "/.qoute/config")
	if fErr != nil {
		if os.IsNotExist(fErr) == true {
			fmt.Println("No config file found.")
			os.Exit(0)
		} else {
			log.Fatalln(fErr)
		}
	}
	file := strings.Split(string(f), "\n")
	for _, l := range file {
		if len(l) != 0 {
			nv := strings.Split(l, "=")
			if nv[0] == name {
				return nv[1]
			}
		}
	}
	return ""
}
