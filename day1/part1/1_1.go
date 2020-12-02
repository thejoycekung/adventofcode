package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func twoSum(filename string) (a, b int) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	// map to hold the numbers
	nums := make(map[int]int)

	// read every line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// check whether its "pair" exists already
		cur, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
		}
		pair := 2020 - cur
		if _, ok := nums[pair]; ok {
			// we found a match!
			a, b = cur, pair
			return
		}
		nums[cur] = 1
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func main() {
	filename := "../input"
	ind1, ind2 := twoSum(filename)
	fmt.Println(ind1 * ind2)
}
