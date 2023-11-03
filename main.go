package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/boreec/repo-downloader/fetcher"
)

func main() {
	debug := flag.Bool("debug", false, "Show debugging messages")

	flag.Parse()

	if flag.NArg() < 1 {
		slog.Error("At least one string argument is expected!")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *debug {
		setLoggerLevelToDebug()
	}

	err := fetcher.DownloadRepositories(flag.Arg(0))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func setLoggerLevelToDebug() {
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(textHandler)
	slog.SetDefault(logger)
}
