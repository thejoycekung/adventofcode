package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func threeSum(filename string) (a, b, c int) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	// create an array of all the lines
	var nums []int
	mapNums := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
		}
		nums = append(nums, num)
		mapNums[num] = 1
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	// sort the array
	sort.Ints(nums)

	for i, curA := range nums {
		for j, curB := range nums {
			need := 2020 - curA - curB
			fmt.Println("looking for:", need)
			if _, ok := mapNums[need]; ok {
				// we found a match!
				a, b, c = nums[i], nums[j], need
				return
			}
		}
	}
	return
}

func main() {
	filename := "../input"
	ind1, ind2, ind3 := threeSum(filename)
	fmt.Println(ind1 * ind2 * ind3)
}
