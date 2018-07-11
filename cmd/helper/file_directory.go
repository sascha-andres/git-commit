package helper

import "os"

// FileExists implements a lazy way to check for a file
func FileExists(path string) bool {
	if stat, err := os.Stat(path); err == nil || os.IsExist(err) {
		return !stat.IsDir()
	}
	return false
}

// DirectoryExists implements a lazy way to check for a directory
func DirectoryExists(path string) bool {
	if stat, err := os.Stat(path); err == nil || os.IsExist(err) {
		if stat.IsDir() {
			return true
		}
	}
	return false
}
