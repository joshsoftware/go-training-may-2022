package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "elephant"
	hint := "biggest animal"
	chances := 10
	guesses := []string{}

	for i := 0; i < len(word); i++ {
		guesses = append(guesses, "_")
	}

	for {
		//add condition show his loss
		userGuess := strings.Join(guesses, "")
		if chances == 0 && userGuess != word {
			fmt.Println("Failed to guess please try again")
			break
		}

		if chances > 0 && userGuess == word {
			fmt.Println("Congratulations")
			break
		}

		fmt.Println(hint)
		fmt.Print(guesses)
		fmt.Println("Enter a your guess")
		guess := ""
		fmt.Scanln(&guess)

	}

}
