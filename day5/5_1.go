package main

import (
	"bufio"
	"fmt"
	"os"
)

func findSeat(parse string) (seatID int) {
	rows := parse[:len(parse)-3]
	cols := parse[len(parse)-3:]
	row := []int{0, 127}
	col := []int{0, 7}
	var rowNum, colNum int
	for i, c := range rows {
		switch c {
		case 'F':
			row = []int{row[0], (row[1] + row[0]) / 2}
		case 'B':
			row = []int{(row[1] + row[0]) / 2, row[1]}
		}
		if i == (len(rows) - 1) {
			if c == 'F' {
				rowNum = row[0]
			} else {
				rowNum = row[1]
			}
		}
	}
	fmt.Println(rowNum)
	for i, c := range cols {
		switch c {
		case 'L':
			col = []int{col[0], (col[1] + col[0]) / 2}
		case 'R':
			col = []int{(col[1] + col[0]) / 2, col[1]}
		}
		if i == (len(cols) - 1) {
			if c == 'L' {
				colNum = col[0]
			} else {
				colNum = col[1]
			}
		}
	}
	fmt.Println(colNum)
	seatID = rowNum*8 + colNum
	return
}

func main() {
	// open file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err.Error())
	}

	// we only need to keep the *current* line in memory
	scanner := bufio.NewScanner(file)
	var maxSeat int
	for scanner.Scan() {
		line := scanner.Text()
		seatID := findSeat(line)
		if maxSeat <= seatID {
			maxSeat = seatID
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(maxSeat)
}
