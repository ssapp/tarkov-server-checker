package log

import (
	"os"
	"path/filepath"
	"time"
)

// listDirectories returns a list of directories in the specified path.
func ListDirectories(path string) ([]string, error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}

	if len(dirs) == 0 {
		return nil, ErrNoLogDirectoriesFound
	}

	return dirs, nil
}

// GetLatestDirectory returns the latest directory from a list of directories.
func GetLatestDirectory(logPath string, dirs []string) string {
	var latestDir string
	for _, dir := range dirs {
		if latestDir == "" || newer(logPath, dir, latestDir) {
			latestDir = dir
		}
	}
	return latestDir
}

// newer checks if one directory is newer than another.
func newer(path, new, old string) bool {
	newTime, _ := getDirModTime(path, new)
	oldTime, _ := getDirModTime(path, old)
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
