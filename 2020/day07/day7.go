package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstColorRegex := regexp.MustCompile(`(?P<container>\w* \w*) bags contain`)
	childColorsRegex := regexp.MustCompile(`((?P<count>\d) (?P<child>\w* \w*) bag)+`)
	rules := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		match := firstColorRegex.FindStringSubmatch(line)
		color := match[1]
		if childColorsRegex.MatchString(line) {
			childMatches := childColorsRegex.FindAllStringSubmatch(line, -1)
			for _, childMatch := range childMatches {
				childCount, _ := strconv.Atoi(childMatch[2])
				for i := 0; i < childCount; i++ {
					rules[color] = append(rules[color], childMatch[3])
				}
			}
		} else {
			rules[color] = []string{}
		}
	}
	//fmt.Println(rules)

	count := 0
	// very slow because I duplicate colors in map to save setup for part 2
	for k, _ := range rules {
		if k != "shiny gold" && containsBag(k, rules, "shiny gold") {
			count++
		}
	}
	fmt.Println(count)
	count = 0
	for _, c := range rules["shiny gold"] {
		count += countBag(c, rules,)
	}
	fmt.Println(count)
}

func containsBag(key string, rules map[string][]string, target string) bool {
	if key == target {
		return true
	}
	if len(rules[key]) == 0 {
		return false
	}
	result := false
	for _, v := range rules[key] {
		result = result || containsBag(v, rules, target)
	}
	return result
}

func countBag(color string, rules map[string][]string) int {
	if len(rules[color]) == 0 {
		return 1
	}
	result := 1
	for _, v := range rules[color] {
		result += countBag(v, rules)
	}
	return result
}
