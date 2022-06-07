package main

import (
	"fmt"
	"strings"
)

func main() {
	//
	word := "elephant"
	// give the hint
	hint := "bigest animal"
	// no of chances
	chances := 10
	//empty slice for the guesses
	guesses := []string{}
	// empty slice for wrong guesses
	wrongGuess := []string{}
	// append "_" into empty guesses slice
	for i := 0; i < len(word); i++ {
		guesses = append(guesses, "_")
	}

	for {
		userGuess := strings.Join(guesses, "")
		// condition for game loss
		if chances == 0 && userGuess != word {
			fmt.Println("You are failed to guess try again")
			break
		}
		// wining condition
		if chances > 0 && userGuess == word {
			fmt.Println("You win the game")
			break
		}

		matched := false

		fmt.Println(hint)
		// print users remaining chances
		fmt.Println("Remaining chances= ", chances)
		// print users guesses
		fmt.Println("correct guesses= ", guesses)
		// print users wrong guesses
		fmt.Println("Wrong guesses=  ", wrongGuess)
		// enter guess
		fmt.Println("enter your guess")
		// declare and initialize variable for user guess
		guess := ""
		// store user input
		fmt.Scanln(&guess)

		for index, char := range word {
			if string(char) == guess {
				guesses[index] = guess
				chances--
				matched = true
			}
		}

		if !matched && !strings.Contains(strings.Join(wrongGuess, ""), guess) {
			wrongGuess = append(wrongGuess, guess)
			chances = chances - 1
		}

	}

}
