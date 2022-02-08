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

	//fmt.Println("Generating Pairs")
	//result := generatePairs()
	//fmt.Println(result)
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

func generatePairs() [20]string {
	rand.Seed(time.Now().UnixNano())
	//letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	pairs := [20]string{}

	for i := 0; i < len(pairs); i++ {
		//letters, pairs[i] = letterSelector(letters[:])
	}
	return pairs
}

func letterSelector(input []string) ([]string, string) {

	//letterArray := make([]string, len(input))
	//reducedLetterArray := make([]string, 0)

	index := rand.Intn(len(input))

	_ = append(input[index+1:], input[:index]...)

	//copy(letterArray, input)
	//firstPart := input[:index]
	//_ = append(input[index+1:], firstPart...)
	//_ = append(input[:index], reducedLetterArray...)

	//reducedLetterArray = append(input[index+1:], reducedLetterArray...)
	//reducedLetterArray = append(input[:index], reducedLetterArray...)
	//return reducedLetterArray, input[index]
	return input[:len(input)-1], input[index-1]

}

//pairs := []string{}
//print the sub slice starting from index 1(included) to index 4(excluded)
//fmt.Println("numbers[1:4] ==", numbers[1:4])

//func randify(characterSet []string) string {
//index := rand.Intn(len(characterSet))
//character := characterSet[index]

//newArray := character[:index]

//newArray = append(newArray, character[index:]...)

//return newArray
//}

//array of 26 characters
//array of 20 empty characters
//take a random character from the first array. put in second.
//-1 from range of random generator
//use the first array without the one that was used.
//repeat

//function that takes in an array of ints
//generates an int, checks if that int already exists in the array
//if it doesnt. put in the array, go to next index.
//if it does, call function that takes in the array.

//var v [20]int
//	rand.Seed(time.Now().UnixNano())
//	for i := 0; i < 20; i++ {
//		tmp := rand.Intn(26)
//		for sr := range v {
//			if sr != tmp{
//				v[i] = tmp
//			} else {
//				tmp := rand.Intn(26)
//			}
//		}
//
//	}
