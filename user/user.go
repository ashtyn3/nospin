package user

import (
	"encoding/json"
	"log"
	"quote/auth"
	"quote/util"
	"strings"

	"github.com/ashtyn3/zi/api"
	zi "github.com/ashtyn3/zi/pkg"
)

// User is the structure for users stored
type User struct {
	Name   string `json:"name"`
	PubTok string `json:"pubTok"`
	PrvTok string `json:"prvTok"`
	ID     string `json:"id"`
}

func Make(name string) User {
	if strings.Index(name, "@") != -1 {
		emailID := strings.Split(name, "@")
		email := name
		name := emailID[0]
		UUID, _ := util.RanString(5)
		UUID = name + "-" + UUID
		pub, _ := util.RanString(32)
		prv, _ := util.RanString(32)
		var user User = User{Name: email, PrvTok: prv, PubTok: pub, ID: UUID}
		data, _ := json.Marshal(user)
		// godotenv.Load("../.env")
		// url := os.Getenv("url")
		// pd := os.Getenv("pd")
		z, err := zi.Zi(auth.Auth().Url, auth.Auth().Pd)
		if err != nil {
			log.Fatalln(err)
		}
		z.Set(api.Pair{Key: user.ID, Value: string(data)})
		return user
	}
	return User{}
}

func Get(name string) User {
	// godotenv.Load("../.env")
	// url := os.Getenv("url")
	// pd := os.Getenv("pd")
	z, err := zi.Zi(auth.Auth().Url, auth.Auth().Pd)
	if err != nil {
		log.Fatalln(err)
	}
	data := z.GetAll()
	for _, u := range data {
		var user User
		json.Unmarshal([]byte(u.Value), &user)
		if user.Name == name {
			return user
		}
	}
	return User{}
}
