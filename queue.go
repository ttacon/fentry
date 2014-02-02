package main

import (
	"fmt"
	"os"
)

type fileQueueNode struct {
	value os.FileInfo
	next  *fileQueueNode
	prev  *fileQueueNode
}

type FileQueue struct {
	head *fileQueueNode
	tail *fileQueueNode
}

func (fq *FileQueue) Enqueue(s os.FileInfo) {
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

func (fq *FileQueue) EnqueueAll(ss []os.FileInfo) {
	for _, s := range ss {
		fq.Enqueue(s)
	}
}

func (fq *FileQueue) Dequeue() (os.FileInfo, error) {
	if fq.head == nil {
		return nil, fmt.Errorf("queue has no item to dequeue")
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
