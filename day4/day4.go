package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Room struct {
	name, checksum string
	sector         int
}

type RuneFreq struct {
	r    rune
	freq int
}
type RuneFreqs []RuneFreq

func (s RuneFreqs) Len() int {
	return len(s)
}
func (s RuneFreqs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s RuneFreqs) Less(i, j int) bool {
	if s[i].freq == s[j].freq {
		return s[i].r < s[j].r
	} else {
		return s[i].freq > s[j].freq
	}
}

func (room Room) validateChecksum() bool {
	runes := []rune(strings.Replace(room.name, "-", "", -1))
	// count up letter frequency
	counts := make(map[rune]int)
	for _, char := range runes {
		curr, ok := counts[char]
		if ok {
			counts[char] = curr + 1
		} else {
			counts[char] = 1
		}
	}
	// make an array of RuneFreq structs we can sort
	runeFreqs := make(RuneFreqs, len(counts))
	j := 0
	for k, v := range counts {
		runeFreqs[j] = RuneFreq{k, v}
		j++
	}
	sort.Sort(runeFreqs)
	// iterate through the checksum and see if it matches
	for i, r := range room.checksum {
		if runeFreqs[i].r != r {
			return false
		}
	}
	return true
}

func (room Room) decrypt() string {
	newRunes := []rune(room.name)
	for i, r := range room.name {
		if r == '-' {
			newRunes[i] = ' '
		} else {
			newRunes[i] = rune(97 + ((int(r) - 97 + room.sector) % 26))
		}
	}
	return string(newRunes)
}

func main() {
	input := getInput("input.txt")
	for _, room := range input {
		if room.validateChecksum() {
			if "northpole object storage" == room.decrypt() {
				fmt.Println(room.sector)
			}
		}
	}
}

func getInt(in string) int {
	n, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return n
}

func getInput(filename string) []Room {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	result := make([]Room, len(lines))
	matcher, _ := regexp.Compile("([a-z-]+)-([0-9]+)\\[([a-z]+)\\]")
	for i, roomID := range lines {
		match := matcher.FindStringSubmatch(roomID)
		sector, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		result[i] = Room{
			name:     match[1],
			sector:   sector,
			checksum: match[3],
		}
	}
	return result
}
