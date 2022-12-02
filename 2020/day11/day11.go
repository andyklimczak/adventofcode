package main

import (
	"bufio"
	"log"
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

	input := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		input = append(input, n)
	}
}
