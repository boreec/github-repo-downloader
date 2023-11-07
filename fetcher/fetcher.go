package fetcher

import (
	"log/slog"

	"github.com/boreec/repo-downloader/model"
	"github.com/boreec/repo-downloader/utils"
	"github.com/gocolly/colly"
)

// Fetch all targets' repositories information (name, url, etc).
//
// Parameters:
//   - targets: slice of github user/organization names with the format
//   `user:user-name` or `org:organization-name`.
//
// Returns:
//   - targetRepos: user associated to their repositories.
//   - errs: potential errors that occurred during fetching.
func FetchAll(targets []string) (
	targetRepos map[string][]model.Repository,
	errs []error,
) {
	targetRepos = make(map[string][]model.Repository)

	for _, target := range targets {
		targetUrl, err := utils.DetermineTargetUrl(target)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		repo, err := FetchRepositoryUrls(targetUrl)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		targetRepos[target] = repo
	}
	return targetRepos, errs
}

// Fetch all public repositories for a target
//
// Parameters:
//   - url: the target repositories GitHub page.
//
// Returns:
//   - repos: slice of repositories found for the target.
//   - err: an error, if any, that occurred in the process.
func FetchRepositoryUrls(url string) (
	repos []model.Repository,
	err error,
) {
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		slog.Debug("Visiting page", slog.String("url", r.URL.String()))
	})

	collector.OnHTML("h3 a", func(e *colly.HTMLElement) {
		repo := model.Repository{
			Name: utils.SanitizeString(e.Text),
			Url:  "https://github.com" + e.Attr("href"),
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

	err = collector.Visit(url)
	if err != nil {
		return nil, err
	}

	return repos, nil
}
