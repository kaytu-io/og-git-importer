// main.go
package main

import (
	"mygitapp/diff"
	"mygitapp/fetch"
	"mygitapp/logger"
	"os"
)

func main() {
	// Set up logging level from environment variable
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info" // Default to "info" if not set
	}
	logger.SetupLogger(logLevel)

	if len(os.Args) < 2 {
		logger.Log.Error("No command provided")
		logger.Log.Error("Usage: <command> [options]")
		return
	}

	command := os.Args[1]

	switch command {
	case "fetch":
		// Allow optional targetDir
		if len(os.Args) < 3 || len(os.Args) > 4 {
			logger.Log.Error("Invalid arguments for fetch")
			logger.Log.Error("Usage: fetch <git URI> [targetDir]")
			return
		}
		gitRepoURI := os.Args[2]
		targetDir := ""
		if len(os.Args) == 4 {
			targetDir = os.Args[3]
		}
		if err := fetch.CloneRepository(gitRepoURI, targetDir); err != nil {
			logger.Log.WithError(err).Error("Fetch operation failed")
		} else {
			logger.Log.Info("Fetch operation completed successfully")
		}
	case "diff":
		diff.RunDiff(os.Args[2:])
	default:
		logger.Log.Error("Unknown command")
		logger.Log.Error("Available commands: fetch, diff")
	}
}
