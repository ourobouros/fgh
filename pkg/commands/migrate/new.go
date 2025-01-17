package migrate

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
)

// Get the new paths for all the repos
func NewPaths(oldRepos []repos.LocalRepo, config configure.RegularOutline) map[string]string {
	spin := spinner.New(utils.SpinnerCharSet, utils.SpinnerSpeed)
	spin.Suffix = fmt.Sprintf(" Getting latest metadata for %v repos", len(oldRepos))
	spin.Start()

	var (
		newPaths = map[string]string{}
		client   = api.GenerateClient(configuration.GetSecrets().PAT)
	)
	for _, repo := range oldRepos {
		metadata, err := api.RepoData(client, repo.Owner, repo.Name)
		if err != nil {
			statuser.Warning(fmt.Sprintf(
				"%v will not be moved because it has either been deleted from github or you don't have access",
				repo.Path,
			))
		}

		newPaths[repo.Path] = repos.RepoLocation(metadata, config)
	}
	spin.Stop()

	if len(newPaths) == 0 {
		os.Exit(0)
	}

	statuser.Success(fmt.Sprintf("Got latest metadata for %v repos", len(oldRepos)))
	return newPaths
}
