package clean

import (
	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/jedib0t/go-pretty/v6/progress"
)

// Get all the repos locally that have been deleted on GitHub
func GetDeleted(pw progress.Writer, clonedRepos []repos.LocalRepo) (deleted []repos.LocalRepo) {
	if !utils.HasInternetConnection() {
		statuser.Warning("Failed to establish an internet connection")
	}

	var (
		client  = api.GenerateClient(configuration.GetSecrets().PAT)
		tracker = progress.Tracker{
			Message: "Checking if any repos have been deleted from GitHub",
			Total:   int64(len(clonedRepos)),
		}
	)
	tracker.SetValue(1)
	pw.AppendTracker(&tracker)

	for _, localRepo := range clonedRepos {
		_, err := api.RepoData(client, localRepo.Owner, localRepo.Name)
		if err != nil {
			deleted = append(deleted, localRepo)
		}
		tracker.Increment(1)
	}

	return deleted
}
