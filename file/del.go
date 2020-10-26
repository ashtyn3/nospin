package file

import (
	"encoding/json"
	"fmt"
	"log"
	"nospin/config"
	"nospin/user"
	"nospin/util"
	"os"
	us "os/user"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/vitecoin/zi/api"
	zi "github.com/vitecoin/zi/pkg"
)

func Del(path string) {
	if strings.Index(path, "/") != -1 {
		group := strings.Split(path, "/")
		godotenv.Load("../.env")
		url := os.Getenv("url")
		pd := os.Getenv("pd")
		z, err := zi.Zi(url, pd)
		if err != nil {
			log.Fatalln(err)
		}
		g := z.GetAll()
		dir := []api.Pair{}
		for _, v := range g {
			if strings.Contains(v.Key, "/") == true {
				data := z.Get(config.Get("name"))
				var z user.User
				json.Unmarshal([]byte(data.Value), &z)
				var file File
				json.Unmarshal([]byte(v.Value), &file)
				tok := z.PrvTok
				group := strings.Split(file.Group, ",")
				if _, ok := util.Find(group, tok); ok == true || file.ID == z.ID {
					dir = append(dir, v)
				}
			}
		}
		if len(dir) == 0 {
			fmt.Println("no file found with the name: " + strings.Join(group[1:], "/"))
			os.Exit(0)
		}
		usr, err := us.Current()
		if err != nil {
			log.Fatal(err)
		}
		home := usr.HomeDir
		p, _ := filepath.Abs(strings.Join(group[1:], "/"))
		p = strings.Replace(p, home+"/", "", 1)
		for _, f := range dir {
			var file File
			json.Unmarshal([]byte(f.Value), &file)
			if p == file.Name || file.Name == strings.Join(group[1:], "/") {
				z.Del(f.Key)
			}
		}
	}
}
