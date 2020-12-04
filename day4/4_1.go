package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkNumFields(passport map[string]string) bool {
	necessary := []string{"eyr", "byr", "iyr", "hgt", "ecl", "hcl", "pid"}
	isValid := true
	for _, n := range necessary {
		if passport[n] == "" {
			isValid = false
		}
	}
	return isValid
}

func main() {
	// open file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err.Error())
	}

	var passports []map[string]string

	scanner := bufio.NewScanner(file)
	curpp := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isValid := checkNumFields(curpp)
			if isValid {
				passports = append(passports, curpp)
			}
			curpp = make(map[string]string)
			continue
		}
		parts := strings.Split(line, " ")
		for _, part := range parts {
			tokens := strings.Split(part, ":")
			curpp[tokens[0]] = tokens[1]
		}
	}
	// must check last passport
	if checkNumFields(curpp) {
		passports = append(passports, curpp)
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(len(passports))
}
