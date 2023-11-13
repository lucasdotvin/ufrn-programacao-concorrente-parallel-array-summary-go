package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
	"time"
)

const resultsDir = "./resultados"

func parseArgs() (n int, t int, err error) {
	if len(os.Args) != 4 {
		return 0, 0, errors.New("expected exactly 2 arguments")
	}

	rawN := os.Args[1]
	n, err = strconv.Atoi(rawN)

	if err != nil {
		return 0, 0, errors.New("expected a number for the first argument")
	}

	rawT := os.Args[2]
	t, err = strconv.Atoi(rawT)

	if err != nil {
		return 0, 0, errors.New("expected a number for the second argument")
	}

	return n, t, nil
}

func worker(id int, tasks <-chan *Task, results chan<- *Result) {
	r := &Result{
		totalSum:                        0,
		grouppedTotalSum:                newGroupTotalSum(),
		tasksWithTotalLesserThan5Count:  0,
		tasksWithTotalGreaterThan5Count: 0,
	}

	for t := range tasks {
		r.totalSum += uint64(t.total)
		r.grouppedTotalSum[t.group] += uint64(t.total)

		if t.total < 5 {
			r.tasksWithTotalLesserThan5Count++
		} else {
			r.tasksWithTotalGreaterThan5Count++
		}
	}

	results <- r
}

func main() {
	n, t, err := parseArgs()

	if err != nil {
		panic(err)
	}

	executionTimestamp := time.Now().Format("2006-01-02T15-04-05")

	resultsFilePath := path.Join(resultsDir, executionTimestamp)
	resultsFile, err := os.OpenFile(resultsFilePath, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(resultsFile, "n: %d, t: %d\n", n, t)

	tasksCount := uint64(math.Pow10(n))
	tasks := make([]*Task, tasksCount)

	for i := uint64(0); i < tasksCount; i++ {
		tasks[i] = newRandomTask(i)
	}

	processingStart := time.Now()

	processingDuration := time.Since(processingStart)

	fmt.Fprintf(resultsFile, "processing duration: %s\n", processingDuration)

	if err := resultsFile.Close(); err != nil {
		panic(err)
	}
}
