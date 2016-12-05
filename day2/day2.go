package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// 10, 11, 12, 13 = A, B, C, D
var keypad = [][]int{
	{-1, -1, 1, -1, -1},
	{-1, 2, 3, 4, -1},
	{5, 6, 7, 8, 9},
	{-1, 10, 11, 12, -1},
	{-1, -1, 13, -1, -1},
}

func main() {
	input := getInput("input.txt")
	fmt.Println(input)
	row, col := 1, 1
	for i := range input {
		for _, move := range input[i] {
			if move == 'U' && row > 0 && keypad[row-1][col] > -1 {
				row = row - 1
			}
			if move == 'D' && row < len(keypad)-1 && keypad[row+1][col] > -1 {
				row = row + 1
			}
			if move == 'L' && col > 0 && keypad[row][col-1] > -1 {
				col = col - 1
			}
			if move == 'R' && col < len(keypad[0])-1 && keypad[row][col+1] > -1 {
				col = col + 1
			}
		}
		fmt.Println(keypad[row][col])
	}
}

func getInput(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(dat)), "\n")
}
