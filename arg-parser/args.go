package args

import (
	"regexp"
)

// Flag is a struct for mapping flags
type Flag struct {
	Flag  string
	Param string
	Index int
}

// ArgParser takes an array of strings and returns parsed value.
func ArgParser(items []string) []Flag {
	var valid []Flag
	for i, arg := range items {
		flagReg := regexp.MustCompile(`^-(\w+)`)
		byteArg := []byte(arg)
		isFlag := flagReg.Match(byteArg)
		if isFlag {
			if i+1+1 > len(items) {
				valid = append(valid, Flag{Param: "", Flag: arg, Index: i})
			} else {
				valid = append(valid, Flag{Param: items[i+1], Flag: arg, Index: i})
			}
		}
	}
	return valid
}
