package file

import (
	"encoding/base64"
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
	//
	z, err := zi.Zi("https://62c4ecd63d32.ngrok.io/")
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
		image := false
		if strings.HasSuffix(strings.Join(group[1:], "/"), ".png") == true || strings.HasSuffix(strings.Join(group[1:], "/"), ".jpg") == true || strings.HasSuffix(strings.Join(group[1:], "/"), ".jpeg") == true {
			f = []byte(base64.StdEncoding.EncodeToString(f))
			image = true
		}
		if user.ID == config.Get("name") {
			p, _ := filepath.Abs(strings.Join(group[1:], "/"))
			if endPath != "" {
				data := File{Content: f, Name: endPath, ID: user.ID, Group: user.PrvTok, Image: image}
				item, _ := json.Marshal(data)
				fID, _ := util.RanString(6)
				if data.Image == true {
					d := util.ChunkString(string(data.Content), 2000)
					for i, c := range d {
						fmt.Printf("\r\033[K%d/%d", i+1, len(d))
						data = File{Content: []byte(c), Name: endPath, ID: user.ID, Group: user.PrvTok, Image: image}
						n, _ := json.Marshal(data)
						// z.Set(api.Pair{Key: user.ID + "/" + fID, Value: string(n)})
						z.Del(user.ID + "/" + fID)
						z.Dump(api.Pair{Key: user.ID + "/" + fID, Value: string(n)}, fID+".zi")
						// time.Sleep(1 * time.Second)
					}
					bFile, _ := json.Marshal(File{Name: endPath, Group: user.PrvTok, Image: image})
					z.Set(api.Pair{Key: user.ID + "/" + fID + "/pointer", Value: string(bFile)})
				} else {
					z.Set(api.Pair{Key: user.ID + "/" + fID, Value: string(item)})
				}
				fmt.Printf("\r\033[K")
				fmt.Printf("\033[F")
				fmt.Println("\n" + data.Name)
			} else {
				data := File{Content: f, Name: strings.Replace(p, home+"/", "", 1), ID: user.ID, Group: user.PrvTok, Image: image}
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
