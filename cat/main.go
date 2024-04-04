package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		for _, arg := range os.Args[1:] {
			data, err := ioutil.ReadFile(arg)
			if err != nil {
				printStr("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
			printStr(string(data))
		}
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
