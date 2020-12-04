package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var result int
	var curPos int
	// open file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err.Error())
	}

	// we only need to keep the *current* line in memory
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for len(line) <= curPos {
			line = strings.Repeat(line, 2)
		}
		if rune(line[curPos]) == rune('#') {
			result++
		}
		curPos += 3
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}
