package fetcher

import (
	"log/slog"

	"github.com/boreec/repo-downloader/model"
	"github.com/boreec/repo-downloader/utils"
	"github.com/gocolly/colly"
)

func FetchAll(targets []string) (map[string][]model.Repository, []error) {
	var errs []error
	var targetRepos map[string][]model.Repository = make(map[string][]model.Repository)

	for _, target := range targets {
		targetUrl := "https://github.com/" + target + "?tab=repositories"
		repo, err := FetchRepositoryUrls(targetUrl)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		targetRepos[target] = repo
	}
	return targetRepos, errs
}

func FetchRepositoryUrls(url string) ([]model.Repository, error) {
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		slog.Debug("Visiting page", slog.String("url", r.URL.String()))
	})

	var repos []model.Repository
	collector.OnHTML("h3 a", func(e *colly.HTMLElement) {
		repo := model.Repository{
			Name: utils.SanitizeString(e.Text),
			Url:  e.Attr("href"),
		}
		repos = append(repos, repo)
	})

	// Check for the "Next" link and follow it if available
	collector.OnHTML(".next_page", func(e *colly.HTMLElement) {
		nextURL := e.Attr("href")
		if nextURL != "" {
			nextRepos, err := FetchRepositoryUrls("https://github.com" + nextURL)
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
