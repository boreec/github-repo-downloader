package fetcher

import (
	"fmt"
	"github.com/gocolly/colly"
	"log/slog"
)

type Repository struct {
	url string
}

func DownloadRepositories(username string) error {
	url := "https://github.com/" + username + "?tab=repositories"
	repositories, err := fetchRepositories(url)
	if err != nil {
		return err
	}
	fmt.Printf("# repo found: %v\n", len(repositories))
	return nil
}

func fetchRepositories(url string) ([]Repository, error) {
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		slog.Debug("Visiting page", slog.String("url", r.URL.String()))
	})

	var repos []Repository
	collector.OnHTML("h3 a", func(e *colly.HTMLElement) {
		repo := Repository{
			url: e.Attr("href"),
		}
		repos = append(repos, repo)
	})

	// Check for the "Next" link and follow it if available
	collector.OnHTML(".next_page", func(e *colly.HTMLElement) {
		nextURL := e.Attr("href")
		if nextURL != "" {
			nextRepos, err := fetchRepositories("https://github.com" + nextURL)
			if err == nil {
				repos = append(repos, nextRepos...)
			}
		}
	})

	err := collector.Visit(url)
	if err != nil {
		return nil, err
	}

	return repos, nil
}
