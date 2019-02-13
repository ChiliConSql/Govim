package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "Hello,World!"
	strings.Map(func(r rune) rune {
		if r == 'W' {
			fmt.Println("Exit...")
			os.Exit(1)
		}
		fmt.Println(r)
		return r
	}, s)
}
