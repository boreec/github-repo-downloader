package fetcher

import (
	"fmt"
	"log"

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
	for _, repo := range repositories {
		fmt.Printf("url found: %v\n", repo.url)
	}
	fmt.Printf("# repo found: %v\n", len(repositories))
	return nil
}

func fetchRepositories(username string) ([]Repository, error) {
	url := "https://github.com/" + username
	fmt.Printf("fetching from %v\n", url)

	collector := colly.NewCollector()

	collector.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	var repos []Repository
	collector.OnHTML("a", func(e *colly.HTMLElement) {
		repo := Repository{}
		repo.url = e.Text
		repos = append(repos, repo)
	})

	err := collector.Visit(url)
	return repos, err
}
