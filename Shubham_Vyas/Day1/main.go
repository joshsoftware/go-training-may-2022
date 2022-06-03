package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "elephant"
	// user entries would be here
	entries := make(map[string]bool)

	// setting length and capacity of slice
	placeholder := []string{}

	for i := 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")
	}

	chances := len(word)

	for {
		// evaluate a loss
		userInput := strings.Join(placeholder, "")
		if chances == 0 && userInput != word {
			fmt.Println("Game Over! Try again")
			break
		}

		// evaluate a win
		if userInput == word {
			fmt.Println("You win")
			break
		}

		// Console display
		fmt.Println()

		// render the placeholder
		fmt.Println(placeholder)

		// render the chances left
		fmt.Printf("Chances: %d\n", chances)

		keys := []string{}

		for k, _ := range entries {
			keys = append(keys, k)
		}

		fmt.Println(keys)
		fmt.Printf("Guess a letter or a word: ")

		// take input from user
		str := ""
		fmt.Scanln(&str)
	}
}
