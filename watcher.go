package fentry

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	RecursiveMode = false
	Duration      time.Duration
)

func GetChangedFiles(dir string) []string {
	var changedFiles []string
	f := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && !RecursiveMode {
			return filepath.SkipDir
		}

		if strings.HasPrefix(filepath.Base(path), ".") {
			return filepath.SkipDir
		}

		if !info.IsDir() && info.ModTime().After(time.Now().Add(Duration*-1)) {
			changedFiles = append(changedFiles, path)
		}
		return nil
	}
	filepath.Walk(dir, f)
	return changedFiles
}
