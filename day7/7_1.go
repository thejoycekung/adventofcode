package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func findBags(bags map[string][]string, bagType string) (result int) {
	curLayer := []string{bagType}
	touched := map[string]bool{
		bagType: true,
	}
	for {
		var nextLayer []string
		for _, v := range curLayer {
			parents := bags[v]
			for _, p := range parents {
				if !touched[p] {
					result++
					touched[p] = true
					nextLayer = append(nextLayer, p)
				}
			}
		}

		if len(nextLayer) == 0 {
			break
		}
		curLayer = nextLayer
	}
	return
}

func countBags(bags map[string]map[string]int, bagType string) (result int) {
	curLayer := map[string]int{
		bagType: 1,
	}
	touched := make(map[string]int)
	for {
		nextLayer := make(map[string]int)
		for b, count := range curLayer {
			innards := bags[b]
			for bag, inside := range innards {
				touched[bag] += inside * count
				nextLayer[bag] += inside * count
			}
		}

		if len(nextLayer) == 0 {
			break
		}
		curLayer = nextLayer
	}
	for _, v := range touched {
		result += v
	}
	return
}

func main() {
	var result int

	// open file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err.Error())
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal(err)
	}
	bags := make(map[string][]string)
	innards := make(map[string]map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, " contain ")
		parts[0] = reg.ReplaceAllString(parts[0], "")
		parts[0] = strings.TrimSuffix(parts[0], "s")
		innards[parts[0]] = make(map[string]int)
		contents := strings.Split(parts[1], ", ")
		for i := range contents {
			bag := reg.ReplaceAllString(contents[i], "")
			if bag == "no other bags" {
				continue
			}
			num := bag[:1]
			bag = bag[2:]
			if bag[len(bag)-1] == 's' {
				bag = bag[:len(bag)-1]
			}
			fmt.Println(bag)
			innards[parts[0]][bag], err = strconv.Atoi(num)
			if err != nil {
				fmt.Println(err.Error())
			}
			bags[bag] = append(bags[bag], parts[0])
		}
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	result = findBags(bags, "shiny gold bag")
	fmt.Println(result)

	result = countBags(innards, "shiny gold bag")
	fmt.Println(result)
}
