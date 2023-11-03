package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/boreec/repo-downloader/fetcher"
)

func main() {
	debug := flag.Bool("debug", false, "Show debugging messages")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("You must provide at least one string argument!")
		flag.PrintDefaults()
		return
	}

	if *debug {
		setLoggerLevelToDebug()
	}

	err := fetcher.DownloadRepositories(flag.Arg(0))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func setLoggerLevelToDebug() {
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(textHandler)
	slog.SetDefault(logger)
}
