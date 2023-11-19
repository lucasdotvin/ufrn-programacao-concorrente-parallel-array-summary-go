package main

import "math/rand"

type Task struct {
	id    uint32
	group uint8
	total uint8
}

func newRandomTask(id uint32) Task {
	return Task{
		id:    id,
		group: uint8(rand.Intn(5)) + 1,
		total: uint8(rand.Intn(11)),
	}
}
