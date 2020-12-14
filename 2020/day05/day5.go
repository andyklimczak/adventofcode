package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var max int
	ids := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		rowInput := line[:7]
		columnInput := line[7:]

		rowVal := findVal(rowInput, 0, 127, 'F', 'B')
		columnVal := findVal(columnInput, 0, 7, 'L', 'R')

		res := rowVal * 8 + columnVal
		if res > max {
			max = res
		}
		//fmt.Println(rowInput, columnInput, rowVal, columnVal, res)
		ids = append(ids, res)
	}
	fmt.Println(max)


	sort.Ints(ids)
	var seatNum int
	for i := 1; i < len(ids) - 1; i++ {
		curr := ids[i]
		next := ids [i + 1]
		if curr + 1 != next {
			seatNum = curr + 1
			break
		}
	}
	fmt.Println(seatNum)
}

func findVal(input string, min, max int, lower, upper rune) int {
	if len(input) == 0 {
		return min
	}
	c := input[0]
	if rune(c) == lower {
		max = (max + min) / 2
	} else {
		min = (max + min + 1) / 2
	}
	//fmt.Println("evaluating", input, string(c), min, max, string(lower), string(upper))
	return findVal(input[1:], min, max, lower, upper)
}