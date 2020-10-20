package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"nospin/config"
	"nospin/user"
	"nospin/util"
	us "os/user"
	"path/filepath"
	"strings"

	"github.com/vitecoin/zi/api"
	zi "github.com/vitecoin/zi/pkg"
)

func Set(path string, endPath string) {
	z, err := zi.Zi("https://b3b8cd3fd294.ngrok.io/")
	if err != nil {
		log.Fatalln(err)
	}
	if strings.Index(path, "/") != -1 {
		group := strings.Split(path, "/")
		f, _ := ioutil.ReadFile(strings.Join(group[1:], "/"))
		usr, err := us.Current()
		if err != nil {
			log.Fatal(err)
		}
		home := usr.HomeDir
		user := user.Get(group[0])
		if user.ID == "" {
			fmt.Println("The user with the name: " + group[0] + " does not exist.")
		}
		if user.ID == config.Get("name") {
			p, _ := filepath.Abs(strings.Join(group[1:], "/"))
			if endPath != "" {
				data := File{Content: f, Name: endPath, ID: user.ID, Group: user.PrvTok}
				item, _ := json.Marshal(data)
				fID, _ := util.RanString(6)
				z.Set(api.Pair{Key: user.ID + "/" + fID, Value: string(item)})
				fmt.Println(data.Name)
			} else {
				data := File{Content: f, Name: strings.Replace(p, home+"/", "", 1), ID: user.ID, Group: user.PrvTok}
				item, _ := json.Marshal(data)
				fID, _ := util.RanString(6)
				z.Set(api.Pair{Key: user.ID + "/" + fID, Value: string(item)})
				fmt.Println(data.Name)

			}
		} else {
			fmt.Println("Cannot set file for user without write access.")
		}

	}
}
