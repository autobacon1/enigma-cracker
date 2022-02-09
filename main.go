package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	plugboardPairs := generate()
	fmt.Println(plugboardPairs)

	returnLetter := swapLetter(plugboardPairs, plugboardPairs[0])
	fmt.Println(returnLetter)
}

func generate() []string {
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	pairs := make([]string, 20)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 20; i++ {
		index := rand.Intn(len(letters))
		pairs[i] = letters[index]
		letters = append(letters[:index], letters[index+1:]...)
	}
	return pairs
}

func generate2(letters []string) []string {
	//letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	pairs := make([]string, 20)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 20; i++ {
		index := rand.Intn(len(letters))
		pairs[i] = letters[index]
		letters = append(letters[:index], letters[index+1:]...)
	}
	return pairs
}

func swapLetter(plugboard []string, character string) string {
	for i := 0; i < len(plugboard); i++ { //find index of input character
		if plugboard[i] == character { //find index of character in plugboard
			if (i % 2) == 0 { //find whether index is odd or even
				return plugboard[i+1] //if even, return the next character
			} else {
				return plugboard[i-1] //if odd, return the previous character
			}
		}
	}
	return character //not used on plugboard
}
