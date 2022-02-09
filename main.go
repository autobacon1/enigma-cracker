package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	plugboardPairs := make([]string, 20)
	fmt.Println(plugboardPairs)
	plugboardPairs = generate()
	fmt.Println(plugboardPairs)
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
