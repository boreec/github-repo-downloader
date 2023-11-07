package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/boreec/repo-downloader/fetcher"
	"github.com/boreec/repo-downloader/logger"
	"github.com/boreec/repo-downloader/repository"
)

const (
	outputDirDefault = "cloned-repos"
)

func main() {
	debug := flag.Bool("debug", false, "Show debugging messages")
	dryRun := flag.Bool("dry-run", false, "Show fetched repository information")
	outputDir := flag.String("output-directory", outputDirDefault, "Output directory for cloned repositories")

	flag.Parse()

	if flag.NArg() < 1 {
		slog.Error("At least one string argument is expected!")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *debug {
		logger.SetLoggerLevelToDebug()
	}

	fetchedRepos, errs := fetcher.FetchAll(flag.Args())
	if len(errs) > 0 {
		for _, err := range errs {
			slog.Warn(err.Error())
		}
	}

	for target, targetRepos := range fetchedRepos {
		slog.Info("list of repositories found", slog.String("target", target))
		for _, repo := range targetRepos {
			slog.Info("", slog.String("url", repo.Url), slog.String("name", repo.Name))
		}
	}

	if !*dryRun {
		if len(fetchedRepos) == 0 {
			slog.Info("no repositories found!")
			os.Exit(0)
		}
		slog.Info("cloning repositories")
		errs = repository.CloneAll(fetchedRepos, *outputDir)
		if len(errs) > 0 {
			for _, err := range errs {
				slog.Warn(err.Error())
			}
		}

	} else {
		slog.Info("dry run, no cloning")
	}

	if len(errs) > 0 {
		os.Exit(1)
	}
	os.Exit(0)
}
