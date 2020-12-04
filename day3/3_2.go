package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkTrees(right int, down int, trees [][]rune) (result int) {
	var curPos int
	for row := 0; row < len(trees); {
		treeline := trees[row]
		if curPos >= len(treeline) {
			curPos = curPos % len(treeline)
		}
		if treeline[curPos] == rune('#') {
			result++
		}
		curPos += right
		row += down
	}
	return
}

func main() {
	// open file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err.Error())
	}

	scanner := bufio.NewScanner(file)
	var trees [][]rune
	for scanner.Scan() {
		line := []rune(scanner.Text())
		trees = append(trees, line)
	}

	one := checkTrees(1, 1, trees)
	three := checkTrees(3, 1, trees)
	five := checkTrees(5, 1, trees)
	seven := checkTrees(7, 1, trees)
	even := checkTrees(1, 2, trees)

	err = file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(one * three * five * seven * even)
}
