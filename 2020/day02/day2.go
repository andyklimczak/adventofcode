package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(`(?P<Min>\d*)-(?P<Max>\d*) (?P<Rune>\S): (?P<Password>\w*)`)

	validPassword := 0
	otherValidPassword := 0
	for scanner.Scan() {
		line := scanner.Text()
		match := r.FindStringSubmatch(line)
		results := map[string]string{}
		for i, name := range match {
			if name != line {
				results[r.SubexpNames()[i]] = name
			}
		}
		password := results["Password"]
		max, _ := strconv.Atoi(results["Max"])
		min, _ := strconv.Atoi(results["Min"])
		rune := results["Rune"]

		count := strings.Count(password, rune)
		fmt.Printf("%#v\n", results)
		fmt.Println(count)
		if count >= min && count <= max {
			validPassword++
		}
		if (string(password[min - 1]) == rune || string(password[max - 1]) == rune) && password[min - 1] != password[max - 1] {
			otherValidPassword++
		}
	}
	fmt.Println(validPassword)
	fmt.Println(otherValidPassword)
}
