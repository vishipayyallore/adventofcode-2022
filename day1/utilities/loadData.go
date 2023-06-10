package utilities

import (
	"io/ioutil"
	"os"
	"strings"
)

func GetInput(filename string) []string {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		print("Error occurred at newDeckFromFile() - ", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), "\r\n")
}
