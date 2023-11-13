package main

import "math/rand"

type Task struct {
	id    uint64
	total uint8
	group uint8
}

func newRandomTask(id uint64) *Task {
	return &Task{
		id:    id,
		total: uint8(rand.Intn(10)),
		group: uint8(rand.Intn(5)),
	}
}
