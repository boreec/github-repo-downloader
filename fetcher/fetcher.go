package fetcher

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Repository struct {
	url string
}

func DownloadRepositories(username string) error {
	repositories, err := fetchRepositories(username)
	if err != nil {
		return err
	}
	fmt.Printf("# repo found: %v\n", len(repositories))
	return nil
}

func fetchRepositories(username string) ([]Repository, error) {
	url := "https://github.com/" + username + "?tab=repositories"
	fmt.Printf("fetching from %v\n", url)

	collector := colly.NewCollector()

	var repos []Repository
	collector.OnHTML("h3 a", func(e *colly.HTMLElement) {
		repo := Repository{
			url: e.Attr("href"),
		}
		repos = append(repos, repo)
	})

	err := collector.Visit(url)
	if err != nil {
		return nil, err
	}
	return repos, nil
}
