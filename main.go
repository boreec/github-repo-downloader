package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("You must provide at least one string argument!")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println("hello word!")
}
