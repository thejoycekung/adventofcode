package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getValidPswds(filename string) (result int) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	// read every line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// retrieve the string
		parts := strings.Split(scanner.Text(), ": ")
		str := parts[1]

		// in the policy, retrieve the letter
		policy := strings.Split(parts[0], " ")
		letter := policy[1]

		// retrieve the ranges
		nums := strings.Split(policy[0], "-")
		bot, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println(err.Error())
		}
		top, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println(err.Error())
		}

		// count how many times the letter appears
		// probably could've done this with a for and runes instead?
		curCount := strings.Count(str, letter)
		if curCount >= bot && curCount <= top {
			result++
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}

func main() {
	filename := "../input"
	result := getValidPswds(filename)
	fmt.Println(result)
}
