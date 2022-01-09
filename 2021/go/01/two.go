package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	f, err := os.Open(args[0])
	check(err)

	var nums []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		check(err)
		nums = append(nums, n)
	}
	check(scanner.Err())

	increases := 0
	for i, n := range nums {
		if i < 3 {
			continue
		}

		if n > nums[i-3] {
			increases += 1
		}
	}
	fmt.Println(increases)
}
