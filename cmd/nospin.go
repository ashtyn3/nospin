package main

import (
	"fmt"
	"nospin/user"
)

func main() {
	// user.UserMake("ashtyn@gmail.com")
	fmt.Println(user.Get("ashtyn@gmail.com"))
}
