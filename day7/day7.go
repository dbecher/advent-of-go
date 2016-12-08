package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func StringHasABBA(s string) bool {
	for i := range s {
		if i < 3 {
			continue
		}
		if s[i-3] == s[i] && s[i-2] == s[i-1] && s[i] != s[i-1] {
			return true
		}
	}
	return false
}

func StringABAs(s string) []string {
	matches := make([]string, 0)
	for i := range s {
		if i < 2 {
			continue
		}
		if s[i-2] == s[i] && s[i] != s[i-1] {
			matches = append(matches, string(s[i-2:i+1]))
		}
	}
	return matches
}

type IPv7Address struct {
	address, hypernet []string
}

func (addr IPv7Address) HasABBA() bool {
	addressHasABBA := false
	for _, str := range addr.address {
		addressHasABBA = StringHasABBA(str)
		if addressHasABBA {
			break
		}
	}
	if !addressHasABBA {
		return false
	}
	for _, str := range addr.hypernet {
		if StringHasABBA(str) {
			return false
		}
	}
	return true
}

func (addr IPv7Address) HasSSL() bool {
	addrABAs := make([]string, 0)
	for _, a := range addr.address {
		// flatten arrays
		addrABAs = append(addrABAs, StringABAs(a)...)
	}
	for _, aba := range addrABAs {
		for _, h := range addr.hypernet {
			if strings.Contains(h, string([]byte{aba[1], aba[0], aba[1]})) {
				return true
			}
		}
	}
	return false
}

func main() {
	addresses := getInput("input.txt")
	cnt := 0
	for _, address := range addresses {
		if address.HasSSL() {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func GetLinesFromFile(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(dat)), "\n")
}

func getInput(filename string) []IPv7Address {
	lines := GetLinesFromFile(filename)
	result := make([]IPv7Address, len(lines))
	addressMatch, _ := regexp.Compile("(?:^|\\])([a-z]+)(?:\\[|$)")
	hypernetMatch, _ := regexp.Compile("\\[([a-z]+)\\]")
	for i, line := range lines {
		addr := IPv7Address{}
		addressMatches := addressMatch.FindAllStringSubmatch(line, -1)
		addr.address = make([]string, len(addressMatches))
		for j, m := range addressMatches {
			addr.address[j] = m[1]
		}
		hypernetMatches := hypernetMatch.FindAllStringSubmatch(line, -1)
		addr.hypernet = make([]string, len(hypernetMatches))
		for j, m := range hypernetMatches {
			addr.hypernet[j] = m[1]
		}
		result[i] = addr
	}
	return result
}
