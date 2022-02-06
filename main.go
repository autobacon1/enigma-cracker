package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hello world")

	rand.Seed(time.Now().UnixNano())

	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	pairs := [20]string{}

	start := time.Now().UTC()

	for i := 0; i < len(pairs); i++ {

		newLetters, character := generatePairs(letters[:])
		pairs[i] = character
		letters = newLetters

	}

	time.Sleep(time.Second * 2)

	end := time.Now().UTC()

	elapsed := end.Sub(start)

	//elapsed := time.Since(start)

	fmt.Println(pairs)

	fmt.Println(elapsed)

}

func generatePairs(letters []string) ([]string, string) {

	index := rand.Intn(len(letters) - 1)

	character := letters[index]

	var subLetters []string

	subLetters = append(letters[1+index:], subLetters...)

	subLetters = append(letters[:index], subLetters...)

	return subLetters, character

	//pairs := []string{}
	//print the sub slice starting from index 1(included) to index 4(excluded)
	//fmt.Println("numbers[1:4] ==", numbers[1:4])

}

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
