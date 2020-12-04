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
		str := []rune(parts[1])

		// in the policy, retrieve the letter
		policy := strings.Split(parts[0], " ")
		letter := []rune(policy[1])

		// retrieve the indices
		// no concept of "index 0" so i will -1
		nums := strings.Split(policy[0], "-")
		bot, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println(err.Error())
		}
		top, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println(err.Error())
		}
		bot--
		top--

		// check both indices
		isBot := str[bot] == letter[0]
		isTop := str[top] == letter[0]

		// XOR them
		if isBot != isTop {
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
	filename := "input"
	result := getValidPswds(filename)
	fmt.Println(result)
}
