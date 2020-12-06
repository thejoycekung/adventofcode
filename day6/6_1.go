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
	curForm := make(map[rune]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			for key := range curForm {
				if curForm[key] == true {
					result++
				}
			}
			curForm = make(map[rune]bool)
			continue
		}
		for _, r := range line {
			curForm[rune(r)] = true
		}
	}
	for key := range curForm {
		if curForm[key] == true {
			result++
		}
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}
