// logger/logger.go
package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Log is the global logger instance
var Log = logrus.New()

// SetupLogger initializes the logger with the specified log level.
func SetupLogger(level string) {
	// Set the log output to stderr to separate it from stdout
	Log.Out = os.Stderr

	// Set the log format to JSON for better integration with logging systems
	Log.SetFormatter(&logrus.JSONFormatter{})

	// Parse and set the log level
	switch strings.ToLower(level) {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "info":
		Log.SetLevel(logrus.InfoLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	default:
		Log.SetLevel(logrus.InfoLevel) // Default to "info" level
	}
}
