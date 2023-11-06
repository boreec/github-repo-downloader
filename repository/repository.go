package repository

import (
	"os"

	"github.com/boreec/repo-downloader/model"
	"github.com/go-git/go-git/v5"
)

func SaveRepositoriesLocally(
	targetRepos map[string][]model.Repository,
	dir string,
) (errs []error) {
	for target, repos := range targetRepos {
		targetDir := dir + "/" + target
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			errs = append(errs, err)
			continue
		}

		for _, repo := range repos {
			targetRepoDir := targetDir + "/" + repo.Name
			_, err := git.PlainClone(targetRepoDir, false, &git.CloneOptions{
				URL:      repo.Url,
				Progress: os.Stdout,
			})
			if err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errs
}
