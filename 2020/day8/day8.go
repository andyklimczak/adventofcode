package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Inst string
	Move int
	Visited bool
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructions := []Instruction{}
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " ")
		inst := s[0]
		move, _ := strconv.Atoi(s[1])
		temp := Instruction{
			Inst:    inst,
			Move:    move,
			Visited: false,
		}
		instructions = append(instructions, temp)
	}
	fmt.Println(instructions)
	fmt.Println(len(instructions))

	//var current *Instruction
	//currentPos := 0
	//count := 0
	//out:
	//for true {
	//	current = &instructions[currentPos]
	//	current.Visited = true
	//	var nextPos int
	//
	//	if current.Inst == "nop" {
	//		nextPos = currentPos + 1
	//	} else if current.Inst == "acc" {
	//		count += current.Move
	//		nextPos = currentPos + 1
	//	} else {
	//		nextPos = currentPos + current.Move
	//	}
	//
	//	fmt.Println(currentPos, nextPos, current, count)
	//	if instructions[nextPos].Visited == true {
	//		break out
	//	} else {
	//		currentPos = nextPos
	//	}
	//}
	//fmt.Println(count)
	count2, _ := runProgram(instructions, -1)
	fmt.Println(count2)

	fmt.Println(instructions)

	for i := 0; i < len(instructions); i++ {
		// reset
		for i := 0; i < len(instructions); i++ {
			s := &instructions[i]
			s.Visited = false
		}

		count, success := runProgram(instructions, i)

		if success {
			fmt.Println("success")
			fmt.Println(count)
			break
		}
	}
}

func runProgram(instructions []Instruction, pos int) (int, bool) {
	var current *Instruction
	currentPos := 0
	count := 0
	success := true

out:
	for currentPos < len(instructions){
		current = &instructions[currentPos]
		current.Visited = true
		var nextPos int
		op := current.Inst

		if currentPos == pos {
			switch op {
			case "jmp":
				op = "nop"
			case "nop":
				op = "jmp"
			}
		}

		if op == "nop" {
			nextPos = currentPos + 1
		} else if op == "acc" {
			count += current.Move
			nextPos = currentPos + 1
		} else {
			nextPos = currentPos + current.Move
		}

		fmt.Println(currentPos, nextPos, current, count)
		if nextPos < len(instructions) && instructions[nextPos].Visited == true {
			success = false
			break out
		} else {
			currentPos = nextPos
		}
	}
	return count, success
}