package service

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

type GitStatus struct {
	Path       string   `json:"path"` // git repository path
	Changes    []string `json:"-"`    // list of changed files
	HasChanges bool     `json:"hasChanges"`
	HasCommits bool     `json:"hasCommits"`
}

// CountGitReposWithUnpushedCommits returns the number of Git repositories under the given directory that have unpushed commits.
func CountGitReposWithUnpushedCommits(dirPath string) ([]GitStatus, error) {
	var repos []GitStatus

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	// 使用 Mutex 保护变化项目列表。
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, f := range files {
		if !f.IsDir() || strings.HasPrefix(f.Name(), ".") {
			continue
		}

		repoPath := filepath.Join(dirPath, f.Name())
		gitPath := filepath.Join(repoPath, ".git")
		if _, err := os.Stat(gitPath); err != nil {
			continue // not a git repo
		}
		wg.Add(1)
		go func(projectName string) {
			defer wg.Done()
			gs, err := analysisGitDir(projectName)
			if err != nil {
				return
			}
			mu.Lock()
			repos = append(repos, gs)
			mu.Unlock()
		}(repoPath)
	}
	wg.Wait()

	return repos, nil
}

func analysisGitDir(repoPath string) (GitStatus, error) {
	status := GitStatus{
		Path: repoPath,
	}

	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = repoPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return status, err
	}

	if len(strings.TrimSpace(string(output))) > 0 {
		status.HasChanges = len(strings.TrimSpace(string(output))) > 0
	}

	cmd = exec.Command("git", "--no-pager", "log", "@{u}..", "--pretty=format:%H")
	cmd.Dir = repoPath
	output, err = cmd.Output()
	if err != nil {
		return status, nil
	}

	hasCommits := len(strings.TrimSpace(string(output))) > 0
	status.HasCommits = hasCommits
	return status, nil
}
