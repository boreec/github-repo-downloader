package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/boreec/repo-downloader/fetcher"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("You must provide at least one string argument!")
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := fetcher.DownloadRepositories(flag.Arg(0))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
