package main

import (
	"fmt"
	"strings"
)

func getKeys(entries map[string]bool) (keys []string) {
	for k := range entries {
		keys = append(keys, k)
	}
	return
}

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

		fmt.Println(getKeys(entries))
		fmt.Printf("Guess a letter or a word: ")

		// take input from user
		str := ""
		fmt.Scanln(&str)
		if len(str) > 2 {
			if str == word {
				fmt.Println("You win!")
				break
			} else {
				entries[str] = true
				chances -= 1
				continue
			}
		}

		// checking if letter exists in the word
		_, ok := entries[str]
		if ok {
			continue
		}

		found := false

		for i, v := range word {
			if string(v) == str {
				placeholder[i] = string(v)
				found = true
			}
		}
		if !found {
			chances -= 1
			entries[str] = true
		}
	}
}
