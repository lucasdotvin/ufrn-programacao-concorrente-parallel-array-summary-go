package main

import "math/rand"

type Task [5]uint8

func newRandomTask(id uint32) Task {
	group := uint8(rand.Intn(5)) + 1
	total := uint8(rand.Intn(11))
	groupAndTotal := group<<4 | total

	idPieces := [4]uint8{
		uint8(id & 0b00000000000000000000000011111111),
		uint8((id & 0b00000000000000001111111100000000) >> 8),
		uint8((id & 0b00000000111111110000000000000000) >> 16),
		uint8((id & 0b11111111000000000000000000000000) >> 24),
	}

	return Task{
		idPieces[0],
		idPieces[1],
		idPieces[2],
		idPieces[3],
		groupAndTotal,
	}
}

func (t Task) id() uint32 {
	return uint32(t[0]) | uint32(t[1])<<8 | uint32(t[2])<<16 | uint32(t[3])<<24
}

func (t Task) group() uint8 {
	return t[4] >> 4
}

func (t Task) total() uint8 {
	return t[4] & 0b00001111
}
