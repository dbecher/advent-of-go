package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	fmt.Println('7')
	doorKey := os.Args[1]
	password := make([]rune, 8)
	found := 0
	byteArr := []byte{}
	m := md5.New()
	var i uint64
MAIN:
	for i = 0; found < 8; i++ {
		byteArr = []byte(fmt.Sprintf("%v%v", doorKey, i))
		m.Reset()
		m.Write(byteArr)
		bs := m.Sum(nil)
		hash := fmt.Sprintf("%x", bs)
		// if position character is not "0" - "7"
		if hash[5] < '0' || hash[5] > '7' {
			continue MAIN
		}
		// get integer value of index position
		index := hash[5] - '0'
		// continue if we already have this position filled
		if password[index] > 0 {
			continue MAIN
		}
		// if hash doesnt start with 00000
		for j := 0; j < 5; j++ {
			if hash[j] != '0' {
				continue MAIN
			}
		}
		password[index] = rune(hash[6])
		found++
	}
	fmt.Println(string(password))
}
