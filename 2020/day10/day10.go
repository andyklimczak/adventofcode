package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		input = append(input, n)
	}

	diffs := map[int]int{1: 0, 2: 0, 3: 0}

	sort.Ints(input)
	input = append(input, input[len(input) - 1] + 3)

	fmt.Println(input)
	current := 0
	for _, n := range input {
		diff := n - current
		diffs[diff]++
		current = n
	}
	fmt.Println(diffs)
	fmt.Println(diffs[1] * diffs[3])
	max := input[len(input) - 1] + 3
	fmt.Println(max)

	fmt.Println("----part2")

	results := startFindCombos(input, max)
	fmt.Println(results)
	fmt.Println(len(results))
}

func startFindCombos(input []int, max int) {
	results := make([][]int, 0)
	result := make([]int, 0)
	results = findCombos(results, result, input, max, 0)
	return results
}

func findCombos(results [][]int, result []int, input []int, max int, current int) [][]int {
	if input[0] == max {
		return results
	}
	for i := 0; i < len(input); i++ {
		if input[i] - current < 3 {
			result
			results
		}
	}
}
