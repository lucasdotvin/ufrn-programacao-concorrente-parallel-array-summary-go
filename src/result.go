package main

type groupTotalSum map[uint8]uint64

type Result struct {
	totalSum                        uint64
	grouppedTotalSum                groupTotalSum
	tasksWithTotalLesserThan5Count  uint64
	tasksWithTotalGreaterThan5Count uint64
}

func newResult() *Result {
	return &Result{
		totalSum:                        0,
		grouppedTotalSum:                newGroupTotalSum(),
		tasksWithTotalLesserThan5Count:  0,
		tasksWithTotalGreaterThan5Count: 0,
	}
}

func newGroupTotalSum() groupTotalSum {
	groups := make(groupTotalSum)

	for g := 1; g < 6; g++ {
		groups[uint8(g)] = 0
	}

	return groups
}

func (r *Result) merge(other *Result) {
	r.totalSum += other.totalSum

	for g, sum := range other.grouppedTotalSum {
		r.grouppedTotalSum[g] += sum
	}

	r.tasksWithTotalLesserThan5Count += other.tasksWithTotalLesserThan5Count
	r.tasksWithTotalGreaterThan5Count += other.tasksWithTotalGreaterThan5Count
}
