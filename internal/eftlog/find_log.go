package eftlog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FindLog finds the latest log file.
func (l *Log) FindLog() (string, error) {
	// Set a default log path if not provided.
	l.Path = defaultIfEmpty(l.Path, LogPathDefault)

	// List all directories in the specified log path.
	files, err := listDirs(l.Path)
	if err != nil {
		return "", err
	}

	// Find the latest directory in the list.
	latestDir, err := findLatestDir(l.Path, files)
	if err != nil {
		return "", err
	}

	// Generate the path to the latest log file.
	logName := fmt.Sprintf(
		"%s application.log",
		strings.TrimPrefix(
			latestDir,
			"log_",
		),
	)

	filePath := filepath.Join(
		l.Path,
		latestDir,
		logName,
	)

	return filePath, nil
}

// defaultIfEmpty returns the default value if the input string is empty.
func defaultIfEmpty(s, defaultValue string) string {
	if s == "" {
		return defaultValue
	}
	return s
}

// listDirs returns a list of directories in the specified log path.
func listDirs(logPath string) (
	[]string, error) {
	// Open the log path directory.
	dir, err := os.Open(logPath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	// Read all files and directories in the log path.
	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	// Collect directory names.
	var dirs []string
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}

	if len(dirs) == 0 {
		return nil, fmt.Errorf("no folders found under %s", logPath)
	}

	return dirs, nil
}

// findLatestDir finds the latest directory from a list of directories.
func findLatestDir(logPath string, dirs []string) (string, error) {
	var latestDir string

	// Iterate through directories and find the latest one.
	for _, dir := range dirs {
		if latestDir == "" || isNewer(logPath, dir, latestDir) {
			latestDir = dir
		}
	}

	return latestDir, nil
}

// isNewer checks if a directory is newer based on modification time.
func isNewer(logPath, newDir, oldDir string) bool {
	newTime, _ := getDirModTime(logPath, newDir)
	oldTime, _ := getDirModTime(logPath, oldDir)
	return newTime.After(oldTime)
}

// getDirModTime returns the modification time of a directory.
func getDirModTime(logPath string, dir string) (time.Time, error) {
	fileInfo, err := os.Stat(filepath.Join(logPath, dir))
	if err != nil {
		return time.Time{}, err
	}
	return fileInfo.ModTime(), nil
}
