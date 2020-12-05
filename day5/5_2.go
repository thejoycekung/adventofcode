package main

import (
	"bufio"
	"fmt"
	"os"
)

type seat struct {
	row int
	col int
	id  int
}

func getSeatInfo(parse string) (seatID, rowNum, colNum int) {
	rows := parse[:len(parse)-3]
	cols := parse[len(parse)-3:]
	row := []int{0, 127}
	col := []int{0, 7}
	for i, c := range rows {
		switch c {
		case 'F':
			row = []int{row[0], (row[1] + row[0]) / 2}
		case 'B':
			row = []int{((row[1] + row[0]) / 2) + 1, row[1]}
		}
		if i == (len(rows) - 1) {
			if c == 'F' {
				rowNum = row[0]
			} else {
				rowNum = row[1]
			}
		}
	}
	for i, c := range cols {
		switch c {
		case 'L':
			col = []int{col[0], (col[1] + col[0]) / 2}
		case 'R':
			col = []int{((col[1] + col[0]) / 2) + 1, col[1]}
		}
		if i == (len(cols) - 1) {
			if c == 'L' {
				colNum = col[0]
			} else {
				colNum = col[1]
			}
		}
	}
	seatID = rowNum*8 + colNum
	return
}

func findMySeat(seatMap [][]int) (seatID int) {
	// 	However, there's a catch: some of the seats at the very front and back of the plane don't exist on this aircraft, so they'll be missing from your list as well.

	// Your seat wasn't at the very front or back, though; the seats with IDs +1 and -1 from yours will be in your list.
	for i, r := range seatMap {
		for j, seat := range r {
			if (i != 0 && i != (len(seatMap)-1)) && seat == 0 {
				seatID = i*8 + j
				if (j != 0 && j != (len(r)-1)) && ((seatMap[i][j-1] != 0) && (seatMap[i][j+1] != 0)) {
					return
				}
			}
		}
	}
	seatID = 0
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
	// set up seatMap
	seatMap := make([][]int, 128)
	for i := range seatMap {
		seatMap[i] = make([]int, 8)
	}
	for scanner.Scan() {
		line := scanner.Text()
		seatID, rowNum, colNum := getSeatInfo(line)
		if maxSeat <= seatID {
			maxSeat = seatID
		}
		seatMap[rowNum][colNum] = seatID
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := range seatMap {
		fmt.Println(seatMap[i])
	}

	// where's my seat?
	mySeat := findMySeat(seatMap)
	fmt.Println(maxSeat)
	fmt.Println(mySeat)
}
