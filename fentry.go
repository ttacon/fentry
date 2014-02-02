package fentry

import "time"

type Fentry struct {
	isRunning bool
	dirs      []string
	Changes   chan []string
}

func NewFentry(dirs []string) *Fentry {
	return &Fentry{
		isRunning: false,
		dirs:      dirs,
		Changes:   make(chan []string),
	}
}

func (f *Fentry) SetDuration(d time.Duration) *Fentry {
	if !f.IsRunning() {
		Duration = d
	}
	return f
}

func (f *Fentry) SetRecursiveWatch(r bool) *Fentry {
	if !f.IsRunning() {
		RecursiveMode = r
	}
	return f
}

func (f *Fentry) IsRunning() bool {
	return f.isRunning
}

func (f *Fentry) Watch() *Fentry {
	f.isRunning = true

	go func() {
		for f.IsRunning() {
			time.Sleep(Duration)
			cs := GetAllChanges(f.dirs)
			if len(cs) > 0 {
				f.Changes <- cs
			}
		}
	}()
	return f
}

func GetAllChanges(dirs []string) []string {
	changes := make(chan []string)
	done := make(chan bool)

	for _, dir := range dirs {
		go worker(dir, changes, done)
	}

	var foundChanges []string
	for i := 0; i < len(dirs); i++ {
		res := <-changes
		done <- true
		if len(res) != 0 {
			foundChanges = append(foundChanges, res...)
		}
	}
	return foundChanges
}

func worker(dir string, changes chan<- []string, done chan bool) {
	changedFiles := GetChangedFiles(dir)
	changes <- changedFiles

	select {
	case _ = <-done:
		return
	}
}
