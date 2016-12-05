package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := getInput("input.txt")
	fmt.Println(input)
	valid := 0
	for _, triangle := range input {
		sort.Ints(triangle)
		if triangle[2] < triangle[1]+triangle[0] {
			valid++
		}
	}
	fmt.Println(valid)
}

func getInt(in string) int {
	n, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return n
}

func getInput(filename string) [][]int {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	result := make([][]int, len(lines))
	cnt := -1
	for i := 0; i < len(lines); i += 3 {
		grouped := [][]string{
			strings.Fields(strings.TrimSpace(lines[i])),
			strings.Fields(strings.TrimSpace(lines[i+1])),
			strings.Fields(strings.TrimSpace(lines[i+2])),
		}
		for j := 0; j < 3; j++ {
			cnt = cnt + 1
			result[cnt] = make([]int, 3)
			result[cnt][0] = getInt(grouped[0][j])
			result[cnt][1] = getInt(grouped[1][j])
			result[cnt][2] = getInt(grouped[2][j])
		}
	}
	return result
}
