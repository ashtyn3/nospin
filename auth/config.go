package auth 

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

func getConfigItem(name string) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := usr.HomeDir
	f, fErr := ioutil.ReadFile(home + "/.quote/config")
	if fErr != nil {
		if os.IsNotExist(fErr) == true {
			return ""
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
