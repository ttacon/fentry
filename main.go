package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

var (
	sleepDuration     = flag.Int("sleepDur", 1, "the amount of time fentry should sleep")
	recursiveModeFlag = flag.Bool("rec", false, "whether directories should be watched recursively or not")
	isRunning         bool
)

type Fentry struct {
	isRunning bool
	dirs      []string
	Changes   chan []string
}

func main() {
	flag.Parse()
	duration = time.Second * time.Duration(*sleepDuration)
	recursiveMode = *recursiveModeFlag
	for {
		time.Sleep(duration)
		//		changes := GetAllChanges
	}

}

func NewFentry(dirs []string) *Fentry {
	return &Fentry{
		isRunning: false,
		dirs:      dirs,
		Changes:   make(chan []string),
	}
}

func (f *Fentry) SetDuration(d time.Duration) {
	if !f.IsRunning() {
		duration = d
	}
}

func (f *Fentry) SetRecursiveWatch(r bool) {
	if !f.IsRunning() {
		recursiveMode = r
	}
}

func (f *Fentry) IsRunning() bool {
	return f.isRunning
}

func (f *Fentry) Watch() {
	go func() {
		for f.IsRunning() {
			time.Sleep(duration)
			cs := GetAllChanges(f.dirs)
			if len(cs) > 0 {
				f.Changes <- cs
			}
		}
	}()

}

func GetAllChanges(dirs []string) []string {
	wg := sync.WaitGroup{}
	changes := make(chan []string)
	wg.Add(len(dirs))
	for _, dir := range dirs {
		go worker(&wg, dir, changes)
	}
	wg.Wait()
	var foundChanges []string
	for i := 0; i < len(dirs); i++ {
		res := <-changes
		if len(res) != 0 {
			foundChanges = append(foundChanges, res...)
		}
	}
	if len(foundChanges) != 0 {
		fmt.Printf("there were changes, files: %s\n", foundChanges)
	}
	return foundChanges
}

func worker(wg *sync.WaitGroup, dir string, changes chan<- []string) {
	changedFiles := GetChangedFiles(dir)
	changes <- changedFiles
	wg.Done()
}
