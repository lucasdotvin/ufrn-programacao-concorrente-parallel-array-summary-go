package main

type groupTotalSum map[uint8]uint64

type Result struct {
	totalSum                        uint64
	grouppedTotalSum                groupTotalSum
	tasksWithTotalLesserThan5Count  uint64
	tasksWithTotalGreaterThan5Count uint64
}

func newGroupTotalSum() groupTotalSum {
	groups := make(groupTotalSum)

	for g := 0; g < 5; g++ {
		groups[uint8(g)] = 0
	}

	return groups
}
