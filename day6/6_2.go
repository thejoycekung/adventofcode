package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// open file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err.Error())
	}

	var result int
	scanner := bufio.NewScanner(file)
	curForm := make(map[rune]int)
	var groupSize int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			for key := range curForm {
				if curForm[key] == groupSize {
					result++
				}
			}
			curForm = make(map[rune]int)
			groupSize = 0
			continue
		}
		for _, r := range line {
			curForm[rune(r)]++
		}
		groupSize++
	}
	// check the last
	for key := range curForm {
		if curForm[key] == groupSize {
			result++
		}
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}
