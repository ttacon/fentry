package fentry

import (
	"fmt"
	"os"
	"path"
)

type fileQueueNode struct {
	value string
	next  *fileQueueNode
	prev  *fileQueueNode
}

type FileQueue struct {
	head *fileQueueNode
	tail *fileQueueNode
}

func (fq *FileQueue) Enqueue(s string) {
	if fq.head == nil {
		fq.head = &fileQueueNode{
			value: s,
		}
		fq.tail = fq.head
		return
	}
	n := &fileQueueNode{
		value: s,
	}
	n.prev = fq.tail
	fq.tail.next = n
	fq.tail = n
}

func (fq *FileQueue) EnqueueAll(baseDir string, ss []os.FileInfo) {
	for _, s := range ss {
		fq.Enqueue(path.Join(baseDir, s.Name()))
	}
}

func (fq *FileQueue) Dequeue() (string, error) {
	if fq.head == nil {
		return "", fmt.Errorf("queue has no item to dequeue")
	}
	toReturn := fq.head.value
	next := fq.head.next
	if next == nil {
		fq.head = nil
		fq.tail = nil
		return toReturn, nil
	}
	next.prev = nil
	fq.head.next = nil
	fq.head = next
	return toReturn, nil
}

func (fq *FileQueue) IsEmpty() bool {
	return fq.head == nil
}
