// Package log provides functionality for extracting the IP address from the EFT log file.
package log

import (
	"fmt"
	"path/filepath"
	"strings"
)

// DefaultLogPath is the default path to the EFT log files.
const DefaultLogPath = "C:\\Battlestate Games\\EFT\\Logs"

// Config holds configuration options for log extraction.
type Config struct {
	LogPath string
}

// Log holds the log instance.
type Log struct {
	config Config
}

// NewLog creates a new Log instance with the given configuration.
func NewLog(config Config) *Log {
	return &Log{config: config}
}

// GetIP gets the IP address from the log file specified in the configuration.
func (l *Log) GetIP() (string, error) {
	filePath, err := l.findLogFile()
	if err != nil {
		return "", err
	}
	return extractIPFromLog(filePath)
}

// findLogFile locates the log file based on the configuration.
func (l *Log) findLogFile() (string, error) {
	logPath := defaultIfEmpty(l.config.LogPath, DefaultLogPath)

	dirs, err := ListDirectories(logPath)
	if err != nil {
		return "", err
	}

	latestDir := GetLatestDirectory(logPath, dirs)
	if latestDir == "" {
		return "", ErrNoLogFilesFound
	}

	logName := fmt.Sprintf("%s application.log", strings.TrimPrefix(latestDir, "log_"))

	return filepath.Join(logPath, latestDir, logName), nil
}

// defaultIfEmpty returns the defaultValue if the string is empty.
func defaultIfEmpty(s, defaultValue string) string {
	if s == "" {
		return defaultValue
	}
	return s
}
