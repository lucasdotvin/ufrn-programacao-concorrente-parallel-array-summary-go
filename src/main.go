package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func parseArgs() (n int, t int, err error) {
	if len(os.Args) != 3 {
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

func worker(tasks []Task, results chan<- *Result) {
	r := newResult()

	for _, t := range tasks {
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

func chunkTasks(tasks []Task, chunksCount int) [][]Task {
	tasksCount := len(tasks)
	chunkSize := tasksCount / chunksCount
	chunks := make([][]Task, chunksCount)

	for i := 0; i < chunksCount; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == chunksCount-1 {
			end = tasksCount
		}

		chunks[i] = tasks[start:end]
	}

	return chunks
}

func main() {
	n, t, err := parseArgs()

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("Parâmetros recebidos:")
	fmt.Printf("n: %d, t: %d: %d\n", n, t)

	tasksCount := uint32(math.Pow10(n))
	tasks := make([]Task, tasksCount)

	for i := uint32(0); i < tasksCount; i++ {
		tasks[i] = newRandomTask(i)
	}

	fmt.Printf("Total de tarefas: %d\n", tasksCount)
	fmt.Println("")

	processingStart := time.Now()

	results := make(chan *Result, t)
	chunks := chunkTasks(tasks, t)

	fmt.Println("Iniciando workers...")

	for i := 0; i < t; i++ {
		go worker(chunks[i], results)
	}

	fmt.Println("Workers iniciados")
	fmt.Println("")
	fmt.Println("Aguardando resultados...")

	r := newResult()

	for i := 0; i < t; i++ {
		r.merge(<-results)
	}

	processingDuration := time.Since(processingStart)

	fmt.Println("Resultados recebidos")
	fmt.Println("")

	fmt.Printf("Soma total: %d\n", r.totalSum)
	fmt.Printf("Tarefas com total menor do que 5: %d\n", r.tasksWithTotalLesserThan5Count)
	fmt.Printf("Tarefas com total maior ou igual a 5: %d\n", r.tasksWithTotalGreaterThan5Count)

	fmt.Printf("Soma por grupo: %d\n", r.totalSum)

	for g := 1; g < 6; g++ {
		fmt.Printf("Grupo %d: %d\n", g, r.grouppedTotalSum[uint8(g)])
	}

	fmt.Println("")

	fmt.Printf("Tempo de processamento total: %s\n", processingDuration)
}
