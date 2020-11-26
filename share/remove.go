package share

import (
	"fmt"
	"os"
	"qoute/file"
	"qoute/user"
	"qoute/util"
	"strings"
)

func Remove(name, fname string) {
	u := user.Get(name)
	// fmt.Println(u)
	// d := z.Get(name)
	if u.Name == "" {
		fmt.Println("Could not find account: " + name)
		os.Exit(0)
	}
	fmt.Println(u.PrvTok)

	f := file.Get(fname)
	g := strings.Split(f.Group, ",")
	tokIndex, ok := util.Find(g, u.PrvTok)

	if ok == false {
		fmt.Println("The file " + fname + " is not shared with: " + name)
		os.Exit(0)
	}
	g[tokIndex] = ""
	f.Group = strings.Join(g, ",")
	// n, _ := json.Marshal(f)
	// file.Set()
	// fID, _ := util.RanString(6)
	// group := strings.Split(fname, "/")
	file.Del(fname)
	file.Set(fname, "", file.Ops{Group: f.Group})
	// z.Set(api.Pair{Key: f.ID + "/" + fID, Value: string(n)})
}
