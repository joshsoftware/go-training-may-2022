package main

import "fmt"

// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strings"
// 	"time"
// )

// func main() {
// 	wordList := [10]string{
// 		"chair",
// 		"table",
// 		"mouse",
// 		"keyboard",
// 	}

// 	rand.Seed(time.Now().UnixNano())
// 	word := wordList[rand.Intn(10)]

// 	blankList := make([]string, len(word))

// 	for i := 0; i < len(blankList); i++ {
// 		blankList[i] = "_"
// 	}

// 	fmt.Println(blankList)
// 	chances := len(word) + 3
// 	for {
// 		chances--
// 		reader := bufio.NewReader(os.Stdin)
// 		fmt.Println("Guess any letter!")
// 		guess, err := reader.ReadString('\n')
// 		guess = guess[:1]
// 		if err != nil {
// 			fmt.Println("An error occured while reading input. Please try again", err)
// 			continue
// 		}
// 		// logic goes after this, to add  guess if it exists in word to blanklist
// 		slicedWord := word
// 		lenSlicedOut := 0
// 		for i := 0; i < strings.Count(slicedWord, guess); i++ {
// 			blankList[strings.IndexAny(slicedWord, guess)+lenSlicedOut] = guess
// 			slicedWord = word[strings.IndexAny(slicedWord, guess):]
// 			lenSlicedOut = len(word) - len(slicedWord)
// 			fmt.Println("Sliced out: ", lenSlicedOut)
// 			fmt.Println("Sliced word: ", slicedWord)

// 		}

// 		if strings.Join(blankList, "") == word {
// 			break
// 		}

// 		fmt.Println("Your guess is", guess)
// 		fmt.Println(blankList)
// 		if chances == 0 {
// 			break
// 		}

// 	}
// 	arrToString := strings.Join(blankList, "")
// 	fmt.Println(arrToString)
// 	if arrToString == word {
// 		fmt.Println("Great Work! You guessed the word correctly.")
// 	} else {
// 		fmt.Println("Oops, You ran out of chances! The word was actually ", word, ".")
// 	}

// }

func foo() {
	fmt.Println("H")
}
func main() {

}
