package share

import (
	"fmt"
	"qoute/file"
	"qoute/user"
	"strings"
)

func Add(name, fname string) {
	u := user.Get(name)
	// fmt.Println(u)
	// d := z.Get(name)
	if u.Name == "" {
		fmt.Println("Could not find account: " + name)
	}
	f := file.Get(fname)
	g := strings.Split(f.Group, ",")

	g = append(g, u.PrvTok)
	f.Group = strings.Join(g, ",")
	// n, _ := json.Marshal(f)
	// file.Set()
	// fID, _ := util.RanString(6)
	// group := strings.Split(fname, "/")
	file.Del(fname)
	file.Set(fname, "", file.Ops{Group: f.Group})
	// z.Set(api.Pair{Key: f.ID + "/" + fID, Value: string(n)})
}
