package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const N = 25

	numsFull := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		numsFull = append(numsFull, n)
	}
	fmt.Println(len(numsFull))

	numsArray := make([]int, 0, 0)
	numsMap := make(map[int]bool, 0)
	for i := 0; len(numsMap) < N; i++ {
		n := numsFull[i]
		numsArray = append(numsArray, n)
		numsMap[n] = true
	}
	fmt.Println(numsMap)
	fmt.Println(len(numsMap))

	var result int
	out:
	for i := N; i < len(numsFull); i++ {
		n := numsFull[i]
		hasMatch := false
		for _, k := range numsArray {
			other := int(math.Abs(float64(n - k)))
			fmt.Println(n, k, other, hasMatch)
			if other != k && numsMap[other] == true {
				hasMatch = true
				break
			}
		}
		if !hasMatch {
			result = n
			break out
		}
		del := numsArray[0]
		delete(numsMap, del)
		numsArray = numsArray[1:]
		numsArray = append(numsArray, n)
		numsMap[n] = true
		fmt.Println(len(numsArray), len(numsMap))
	}
	fmt.Println(result)
	fmt.Println("----part2")
	// pretty slow

	var finalSlice []int
	finish:
	for i := 0; i < len(numsFull) - 2; i++ {
		for j := i + 2; j < len(numsFull); j++ {
			a := numsFull[i:j]
			sum := 0
			for _, n := range a {
				sum += n
			}
			fmt.Println(a, sum)
			if sum == result {
				finalSlice = a
				break finish
			}
		}
	}
	fmt.Println(finalSlice)
	min := finalSlice[0]
	max := finalSlice[0]
	for _, n := range finalSlice {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println(min, max, min + max)
}
