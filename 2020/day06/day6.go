package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	part1 := partOne()
	fmt.Println(part1)
	part2 := partTwo()
	fmt.Println(part2)
}

func partOne() int {

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	answers := make(map[rune]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			for _, s := range line {
				answers[s] = true
			}
		} else {
			count += len(answers)
			answers = make(map[rune]bool)
		}
	}
	return count
}

func partTwo() int {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	peopleCount := 0
	answers := make(map[rune]int , 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			peopleCount++
			for _, s := range line {
				answers[s] += 1
			}
		} else {
			allAnswered := 0
			for _, v := range answers {
				if v == peopleCount {
					allAnswered++
				}
			}
			count += allAnswered
			answers = make(map[rune]int)
			peopleCount = 0
		}
	}
	return count
}
