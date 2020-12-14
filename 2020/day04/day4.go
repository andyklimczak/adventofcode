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
	r := regexp.MustCompile(`\S*:\S*`)
	temp := make(map[string]string)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			vals := r.FindAllString(line, -1)
			for _, val := range vals {
				split := strings.Split(val, ":")
				k, v := split[0], split[1]
				temp[k] = v
			}
		} else {
			//fmt.Println(temp)
			if validate("byr", temp["byr"]) &&
				validate("iyr", temp["iyr"]) &&
				validate("eyr", temp["eyr"]) &&
				validate("hgt", temp["hgt"]) &&
				validate("hcl", temp["hcl"]) &&
				validate("ecl", temp["ecl"]) &&
				validate("pid", temp["pid"]) {
				//fmt.Println("valid")
				//fmt.Println(len(temp))
				count++
			}
			temp = make(map[string]string)
		}
	}
	fmt.Println("count", count)
}

func validate(key, val string) (result bool) {
	result = false
	switch key {
	case "byr":
		year, _ := strconv.Atoi(val)
		result = len(val) == 4 &&  year >= 1920 && year <= 2002
	case "iyr":
		year, _ := strconv.Atoi(val)
		result = len(val) == 4 &&  year >= 2010 && year <= 2020
	case "eyr":
		year, _ := strconv.Atoi(val)
		result = len(val) == 4 &&  year >= 2020 && year <= 2030
	case "hgt":
		if len(val) > 0 {
			height, _ := strconv.Atoi(val[:len(val)-2])
			if strings.Contains(val, "in") {
				result = height >= 59 && height <= 76
			} else {
				result = height >= 150 && height <= 193
			}
		}
	case "hcl":
		r := regexp.MustCompile(`#[a-z0-9]{6}`)
		result = r.MatchString(val)
	case "ecl":
		validEyeColours := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
		result = validEyeColours[val]
	case "pid":
		r := regexp.MustCompile(`^[0-9]{9}$`)
		result = r.MatchString(val)
	}
	//fmt.Println(key, val, result)
	return
}