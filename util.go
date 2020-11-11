package util

import (
	"os"
)

//numbers
const (
	Thousand = 1000
	Million  = Thousand * Thousand
	Billion  = Thousand * Million
)

//Global strings
var (
	HomeDir string
)

func init() {
	var err error
	HomeDir, err = os.UserHomeDir()

	if err != nil {
		panic(err)
	}

}
