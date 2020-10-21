package file

import (
	"encoding/base64"
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

	"github.com/vitecoin/zi/api"
	zi "github.com/vitecoin/zi/pkg"
)

type File struct {
	ID      string
	Name    string
	Content []byte
	Group   string
	Image   bool
}

func Get(id string) File {
	z, err := zi.Zi("https://b3b8cd3fd294.ngrok.io/")
	if err != nil {
		log.Fatalln(err)
	}
	if strings.Index(id, "/") != -1 {
		group := strings.Split(id, "/")
		dir := []api.Pair{}
		g := z.GetAll()
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
		img := File{}
		fFull := []string{}
		image := false
		for _, f := range dir {
			var file File
			json.Unmarshal([]byte(f.Value), &file)
			if strings.HasSuffix(strings.Join(group[1:], "/"), ".png") == true || strings.HasSuffix(strings.Join(group[1:], "/"), ".jpg") == true || strings.HasSuffix(strings.Join(group[1:], "/"), ".jpeg") == true {
				e, _ := base64.StdEncoding.DecodeString(string(file.Content))
				file.Content = e
			}
			if p == file.Name || file.Name == strings.Join(group[1:], "/") {
				if file.Image == true {
					img = file
					fFull = append(fFull, string(file.Content))
					image = true
				} else {
					return file
				}
			}
		}
		if image == true {
			val := strings.Join(fFull, "")
			img.Content = []byte(val)
			return img
		}
	}
	return File{}
}
