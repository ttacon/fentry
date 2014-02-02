package main

import (
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

var (
	recursiveMode = false
	duration      time.Duration
)

type syncAlteredFiles struct {
	*sync.RWMutex
	files []string
}

type Notifier struct {
	data *syncAlteredFiles
	dur  time.Duration
}

type dirWatcher struct {
	changes chan<- []string
	dir     string
}

func GetChangedFiles(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	fileQueue := &FileQueue{}
	fileQueue.EnqueueAll(files)
	var changedFiles []string
	for !fileQueue.IsEmpty() {
		file, err := fileQueue.Dequeue()
		if err != nil {
			panic(err)
		}
		if file.IsDir() && recursiveMode && !strings.HasPrefix(file.Name(), ".") {
			dirFiles, err := ioutil.ReadDir(file.Name())
			if err != nil {
				panic(err)
			}
			fileQueue.EnqueueAll(dirFiles)
		} else {
			if file.ModTime().After(time.Now().Add(-1 * duration)) {
				changedFiles = append(changedFiles, file.Name())
			}
		}
	}
	return changedFiles
}
