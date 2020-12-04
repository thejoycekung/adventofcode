package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func yearCheck(val string, lower int, higher int) bool {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println(err.Error)
	}
	return len(val) == 4 && (intVal <= higher && intVal >= lower)
}

func eyeCheck(val string) bool {
	eyes := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, eye := range eyes {
		if val == eye {
			return true
		}
	}
	return false
}

func heightCheck(val string) bool {
	unit := val[len(val)-2:]
	if unit != "cm" && unit != "in" {
		return false
	}
	num, err := strconv.Atoi(val[:len(val)-2])
	if err != nil {
		fmt.Println(err.Error)
		return false
	}
	return (unit == "cm" && (num <= 193 && num >= 150)) || (unit == "in" && (num <= 76 && num >= 59))
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

	// now that we have a list of valid passports ...
	hexCode := regexp.MustCompile("#[A-Fa-f0-9]{6}")
	var validPassports int
	for _, pp := range passports {
		byr := yearCheck(pp["byr"], 1920, 2002)
		if !byr {
			continue
		}
		eyr := yearCheck(pp["eyr"], 2020, 2030)
		if !eyr {
			continue
		}
		iyr := yearCheck(pp["iyr"], 2010, 2020)
		if !iyr {
			continue
		}
		ecl := eyeCheck(pp["ecl"])
		if !ecl {
			continue
		}
		hgt := heightCheck(pp["hgt"])
		if !hgt {
			continue
		}
		pid := len(pp["pid"]) == 9
		if !pid {
			continue
		}
		hcl := hexCode.MatchString(pp["hcl"])
		if !hcl {
			continue
		}
		validPassports++
	}
	fmt.Println(validPassports)
}
