package file

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	us "os/user"
	"path/filepath"
	"qoute/auth"
	"qoute/config"
	"qoute/user"
	"qoute/util"
	"strings"

	"github.com/ashtyn3/zi/api"
	zi "github.com/ashtyn3/zi/pkg"
)

type File struct {
	ID      string
	Name    string
	Content []byte
	Group   string
	Image   bool
}

func Get(id string) File {
	var dirString string
	// godotenv.Load("../.env")
	// url := os.Getenv("url")
	// pd := os.Getenv("pd")
	z, err := zi.Zi(auth.Auth().Url, auth.Auth().Pd)
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
		dirContent := []string{}
		for _, f := range dir {
			var file File
			json.Unmarshal([]byte(f.Value), &file)
			if strings.HasSuffix(strings.Join(group[1:], "/"), ".png") == true || strings.HasSuffix(strings.Join(group[1:], "/"), ".jpg") == true || strings.HasSuffix(strings.Join(group[1:], "/"), ".jpeg") == true || file.Image == true {
				e, _ := base64.StdEncoding.DecodeString(string(file.Content))
				file.Content = e
			}
			// s := strings.Split(file.Name, "/"); s[len(s) - 1]
			if strings.Join(group[1:], "/") == "/" {
				fName := file.Name
				dirContent = append(dirContent, fName)
			} else if strings.HasSuffix(id, "/") == true && strings.Contains(file.Name, strings.Join(group[1:], "/")) && id != "/" {
				fName := strings.Replace(file.Name, strings.Join(group[1:], "/"), "", 1)
				dirContent = append(dirContent, fName)
			} else if p == file.Name || file.Name == strings.Join(group[1:], "/") {
				if file.Image == true {
					pulled := z.Get("^" + strings.Replace(f.Key, "/pointer", "", 1))
					var chunks []api.Pair
					json.Unmarshal([]byte(pulled.Value), &chunks)
					image = true
					for _, chunk := range chunks {
						var filechunk File
						img = filechunk
						json.Unmarshal([]byte(chunk.Value), &filechunk)
						d, _ := base64.StdEncoding.DecodeString(string(filechunk.Content))
						fFull = append(fFull, string(d))
					}
					file.Content = []byte(strings.Join(fFull, ""))
					return file
				}
				return file

			}
		}
		if len(dirContent) != 0 {
			dirString += fmt.Sprintln("Contents of " + strings.Join(group[1:], "/") + ":")
			for _, d := range dirContent {
				dirString += d + "\n"
			}
			// fmt.Printf("\033[F")
			return File{Content: []byte(base64.StdEncoding.EncodeToString([]byte(dirString)))}
		}
		if image == true {
			val := strings.Join(fFull, "")
			img.Content = []byte(val)
			return img
		}
	}
	return File{}
}
