package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	array := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		array = append(array, []rune(line))
	}

	treeCount := treeCountParam(3,1,array)

	fmt.Println(treeCount)
	fmt.Println("part 2 ---")

	vals := []int{
		treeCountParam(1,1,array),
		treeCountParam(3,1,array),
		treeCountParam(5,1,array),
		treeCountParam(7,1,array),
		treeCountParam(1,2,array),
	}
	res := 1
	for _, val := range vals {
		res *= val
	}
	fmt.Println(res)

}

func treeCountParam(right, down int, array [][]rune) (treeCount int) {
	treeCount = 0
	y := 1
	for x, index := 0, 0; x < len(array); x, index = x+down, index + 1 {
		line := array[x]
		if x > 0 {
			y = (index * right) % len(line)
			if line[y] == '#' {
				treeCount++
			}
		}
	}
	return
}