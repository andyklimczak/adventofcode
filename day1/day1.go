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

	file, err := os.Open("./day1")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    nums := make([]int, 0, 0)
    for scanner.Scan() {
        num, _ := strconv.Atoi(scanner.Text())
        nums = append(nums, num)
    }

    var result int
    for _, n := range nums {
    	for _, m := range nums {
    		if n + m == 2020 {
    			result = n * m
			}
		}
	}
	fmt.Println(result)

	sort.Ints(nums)
    out:
    for i := 0; i < len(nums); i++ {
    	l := i + 1
    	r := len(nums) - 1
    	for l < r {
    		sum := nums[i] + nums[l] + nums[r]
    		if sum == 2020 {
    			result = nums[i] * nums[l] * nums[r]
    			break out
			} else if sum < 2020 {
				l++
			} else {
				r--
			}
		}
	}
	fmt.Println(result)
}