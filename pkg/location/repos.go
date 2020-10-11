package location

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/statuser/v2"
)

// Get the root GitHub folder
func GitHubFolder() string {
	var path string
	path, err := os.UserHomeDir()
	if err != nil {
		statuser.Error("Failed to get home directory", err, 1)
	}
	return filepath.Join(path, "github")
}

// Get all repos downloaded
func Repos() (repos []string) {
	ghFolder := GitHubFolder()
	chdir(ghFolder, ghFolder)
	var cwd string
	for _, user := range dirs() {
		cwd = filepath.Join(ghFolder, user)
		chdir(ghFolder, cwd)
		for _, repoType := range dirs() {
			cwd = filepath.Join(ghFolder, user, repoType)
			chdir(ghFolder, cwd)
			for _, language := range dirs() {
				cwd = filepath.Join(ghFolder, user, repoType, language)
				chdir(ghFolder, cwd)
				for _, repoName := range dirs() {
					cwd = filepath.Join(ghFolder, user, repoType, language, repoName)
					chdir(ghFolder, cwd)
					repos = append(repos, cwd)
				}
			}
		}
	}
	return repos
}

func dirs() (folders []string) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		cwd, err1 := os.Getwd()
		if err1 != nil {
			statuser.Error("Failed to get current working directory", err, 1)
		}
		statuser.Error("Failed to list "+cwd, err, 1)
	}

	for _, f := range files {
		if f.IsDir() {
			folders = append(folders, f.Name())
		}
	}
	return folders
}

func chdir(ghFolder string, folder string) {
	if !strings.HasPrefix(folder, ghFolder) {
		folder = filepath.Join(ghFolder, folder)
	}
	err := os.Chdir(folder)
	if err != nil {
		statuser.Error("Failed to change directory to "+folder, err, 1)
	}
}