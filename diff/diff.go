// diff/diff.go
package diff

import (
	"encoding/json"
	"fmt"
	"mygitapp/fetch"  // Import the fetch package
	"mygitapp/logger" // Import the logger package
	"os"
	"path/filepath"
	"strings"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// CommitDetails holds commit metadata.
type CommitDetails struct {
	Hash      string    `json:"hash"`
	Timestamp time.Time `json:"timestamp"`
}

// ComparisonResultGrouped holds the JSON output structure with files grouped under parent folders.
type ComparisonResultGrouped struct {
	CommitDetails   [2]CommitDetails    `json:"commit_details"`
	ModifiedFiles   map[string][]string `json:"modified_files"`
	CreatedFiles    map[string][]string `json:"created_files"`
	DeletedFiles    map[string][]string `json:"deleted_files"`
	UnmodifiedFiles map[string][]string `json:"unmodified_files"`
	ChangedFolders  []string            `json:"changed_folders"`
}

// RunDiff runs the diff comparison.
func RunDiff(args []string) {
	if len(args) < 2 || len(args) > 3 {
		logger.Log.Error("Invalid arguments for diff")
		logger.Log.Error("Usage: diff <repository path or remote URI> <first_commit SHA-1> [second_commit SHA-1]")
		os.Exit(1)
	}

	repoPathOrURI := args[0]
	firstCommitSHA := args[1]
	secondCommitSHA := ""
	if len(args) == 3 {
		secondCommitSHA = args[2]
	}

	var repo *git.Repository
	var err error

	if isRemoteURI(repoPathOrURI) {
		// Use fetch package to clone the repository to a temporary directory
		tempDir, err := os.MkdirTemp("", "git-repo-*")
		if err != nil {
			logger.Log.WithError(err).Error("Failed to create temporary directory")
			os.Exit(1)
		}
		defer os.RemoveAll(tempDir)

		logger.Log.Infof("Cloning repository from %s to %s", repoPathOrURI, tempDir)
		err = fetch.CloneRepository(repoPathOrURI, tempDir)
		if err != nil {
			logger.Log.WithError(err).Error("Failed to clone repository using fetch package")
			os.Exit(1)
		}

		// Open the cloned repository
		repo, err = git.PlainOpen(tempDir)
		if err != nil {
			logger.Log.WithError(err).Error("Failed to open cloned repository")
			os.Exit(1)
		}
	} else {
		// Open the local repository
		repo, err = git.PlainOpen(repoPathOrURI)
		if err != nil {
			logger.Log.WithError(err).Error("Failed to open local repository")
			os.Exit(1)
		}
	}

	firstCommit, err := getCommit(repo, firstCommitSHA)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to retrieve first_commit")
		os.Exit(1)
	}

	if secondCommitSHA == "" {
		secondCommitSHA, err = getLastCommitSHA(repo)
		if err != nil {
			logger.Log.WithError(err).Error("Failed to get the latest commit")
			os.Exit(1)
		}
	}

	secondCommit, err := getCommit(repo, secondCommitSHA)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to retrieve second_commit")
		os.Exit(1)
	}

	// Ensure first_commit is older
	if firstCommit.Committer.When.After(secondCommit.Committer.When) {
		firstCommit, secondCommit = secondCommit, firstCommit
	}

	result := compareCommits(firstCommit, secondCommit)
	result.CommitDetails[0] = CommitDetails{Hash: firstCommit.Hash.String(), Timestamp: firstCommit.Committer.When}
	result.CommitDetails[1] = CommitDetails{Hash: secondCommit.Hash.String(), Timestamp: secondCommit.Committer.When}

	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		logger.Log.WithError(err).Error("Failed to marshal result to JSON")
		os.Exit(1)
	}

	// Print JSON to stdout
	fmt.Println(string(output))

	logger.Log.Info("Diff operation completed successfully")
}

// isRemoteURI checks if the provided path is a remote URI.
func isRemoteURI(uri string) bool {
	return strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") || strings.HasPrefix(uri, "git@")
}

// getCommit retrieves the commit object by its SHA-1 hash.
func getCommit(repo *git.Repository, sha string) (*object.Commit, error) {
	commitHash := plumbing.NewHash(sha)
	return repo.CommitObject(commitHash)
}

// getLastCommitSHA gets the SHA of the latest commit on the default branch.
func getLastCommitSHA(repo *git.Repository) (string, error) {
	ref, err := repo.Head()
	if err != nil {
		return "", err
	}
	return ref.Hash().String(), nil
}

// compareCommits analyzes the differences between two commits.
func compareCommits(firstCommit, secondCommit *object.Commit) ComparisonResultGrouped {
	tree1, err := firstCommit.Tree()
	if err != nil {
		logger.Log.WithError(err).Error("Failed to get tree for first_commit")
		return ComparisonResultGrouped{}
	}
	tree2, err := secondCommit.Tree()
	if err != nil {
		logger.Log.WithError(err).Error("Failed to get tree for second_commit")
		return ComparisonResultGrouped{}
	}

	patch, err := tree1.Patch(tree2)
	if err != nil {
		logger.Log.WithError(err).Error("Failed to create patch between commits")
		return ComparisonResultGrouped{}
	}

	modifiedFiles := make(map[string][]string)
	createdFiles := make(map[string][]string)
	deletedFiles := make(map[string][]string)
	unmodifiedFiles := make(map[string][]string)
	changedFoldersSet := make(map[string]bool)

	changedFilesSet := make(map[string]bool)
	filePatches := patch.FilePatches()
	for _, filePatch := range filePatches {
		from, to := filePatch.Files()
		if from == nil && to != nil {
			// File was created
			createdPath := to.Path()
			parentFolder := getParentFolder(createdPath)
			createdFiles[parentFolder] = append(createdFiles[parentFolder], filepath.Base(createdPath))
			changedFilesSet[createdPath] = true
			changedFoldersSet[parentFolder] = true
		} else if from != nil && to == nil {
			// File was deleted
			deletedPath := from.Path()
			parentFolder := getParentFolder(deletedPath)
			deletedFiles[parentFolder] = append(deletedFiles[parentFolder], filepath.Base(deletedPath))
			changedFilesSet[deletedPath] = true
			changedFoldersSet[parentFolder] = true
		} else if from != nil && to != nil {
			// File was modified
			modifiedPath := from.Path()
			parentFolder := getParentFolder(modifiedPath)
			modifiedFiles[parentFolder] = append(modifiedFiles[parentFolder], filepath.Base(modifiedPath))
			changedFilesSet[modifiedPath] = true
			changedFoldersSet[parentFolder] = true
		}
	}

	// Collect all files from both trees
	allFilesSet := make(map[string]bool)
	err = tree1.Files().ForEach(func(f *object.File) error {
		allFilesSet[f.Name] = true
		return nil
	})
	if err != nil {
		logger.Log.WithError(err).Error("Failed to iterate over files in first_commit tree")
		return ComparisonResultGrouped{}
	}
	err = tree2.Files().ForEach(func(f *object.File) error {
		allFilesSet[f.Name] = true
		return nil
	})
	if err != nil {
		logger.Log.WithError(err).Error("Failed to iterate over files in second_commit tree")
		return ComparisonResultGrouped{}
	}

	// Identify unmodified files
	for file := range allFilesSet {
		if !changedFilesSet[file] {
			parentFolder := getParentFolder(file)
			unmodifiedFiles[parentFolder] = append(unmodifiedFiles[parentFolder], filepath.Base(file))
		}
	}

	// Convert changedFoldersSet to a slice
	changedFolders := make([]string, 0, len(changedFoldersSet))
	for folder := range changedFoldersSet {
		changedFolders = append(changedFolders, folder)
	}

	return ComparisonResultGrouped{
		ModifiedFiles:   modifiedFiles,
		CreatedFiles:    createdFiles,
		DeletedFiles:    deletedFiles,
		UnmodifiedFiles: unmodifiedFiles,
		ChangedFolders:  changedFolders,
	}
}

// getParentFolder returns the parent folder of a given file path.
// If the file is at the root, it returns "root".
func getParentFolder(filePath string) string {
	dir := filepath.Dir(filePath)
	if dir == "." || dir == "/" {
		return "root"
	}
	return filepath.ToSlash(dir)
}
