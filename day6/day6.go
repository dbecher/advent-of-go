package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type CharFreq struct {
	r    rune
	freq int
}
type CharFreqs []CharFreq

func (s CharFreqs) Len() int {
	return len(s)
}
func (s CharFreqs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s CharFreqs) Less(i, j int) bool {
	if s[i].freq == s[j].freq {
		return s[i].r < s[j].r
	} else {
		return s[i].freq > s[j].freq
	}
}

func CharacterFrequencies(s string) CharFreqs {
	// count up letter frequency
	counts := make(map[rune]int)
	for _, char := range s {
		curr, ok := counts[char]
		if ok {
			counts[char] = curr + 1
		} else {
			counts[char] = 1
		}
	}
	// make an array of CharFreq structs we can sort
	runeFreqs := make(CharFreqs, len(counts))
	j := 0
	for k, v := range counts {
		runeFreqs[j] = CharFreq{k, v}
		j++
	}
	sort.Sort(runeFreqs)
	// iterate through the checksum and see if it matches
	return runeFreqs
}

func main() {
	input := getInput("input.txt")
	result := make([]rune, len(input))
	for i, line := range input {
		freqs := CharacterFrequencies(line)
		result[i] = freqs[len(freqs)-1].r
	}
	fmt.Println(string(result))
}

func getInput(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	result := make([]string, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		colChars := make([]rune, len(lines))
		for j, line := range lines {
			colChars[j] = rune(line[i])
		}
		result[i] = string(colChars)
	}
	return result
}
