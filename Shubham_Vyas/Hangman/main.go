package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Requried fields for Hangman
type Hangman struct {
	Entries     map[string]bool
	Placeholder []string
	Chances     int
}

// Response of get api used to fetch random word
type VercelResponse struct {
	Word          string `json:"word"`
	Defination    string `json:"defination"`
	Pronunciation string `json:"pronunciation"`
}

// Get a random word
func getWord() (word string, err error) {
	response, err := http.Get("https://random-words-api.vercel.app/word")
	if err != nil {
		return
	}
	// Close response Body before leaving function
	defer response.Body.Close()
	// Read response body
	resBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body")
		return
	}

	// Response will be a list of vercelresponse object
	var randomWords []VercelResponse

	// Response will be in json need to parse it
	err = json.Unmarshal(resBytes, &randomWords)
	if err != nil {
		fmt.Println("Failed to parse response")
		return
	}

	word = randomWords[0].Word
	return
}

// Return keys from map
func getKeys(entries map[string]bool) (keys []string) {
	for k := range entries {
		keys = append(keys, k)
	}
	return
}

// Start Hangman game
func (h *Hangman) start(word string) error {
	for {
		userInput := strings.Join(h.Placeholder, "")
		if h.Chances == 0 && userInput != word {
			fmt.Println("Game Over! Try Again")
			fmt.Println("You failed to guess: ", word)
			break
		}

		if userInput == word {
			fmt.Println("You win!")
			break
		}

		// Write message
		fmt.Println()
		fmt.Println(h.Placeholder)
		fmt.Println("Guesses: ", getKeys(h.Entries))
		fmt.Println("Chances: ", h.Chances)

		fmt.Printf("Guess a letter or a word: ")

		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			return err
		}
		// If word or letter already guessed no need to waste a chance
		if _, ok := h.Entries[input]; ok {
			continue
		}
		// if not a char lets check for word
		if len(input) > 1 {
			if input == word {
				fmt.Println("You win")
				break
			} else {
				h.Entries[input] = false
				h.Chances -= 1
				continue
			}
		}
		doesStringMatch := false
		for i, v := range word {
			// as v is type rune
			if string(v) == input {
				h.Placeholder[i] = input
				doesStringMatch = true
			}
		}
		if !doesStringMatch {
			h.Entries[input] = false
			h.Chances -= 1
		}
	}
	return nil
}

func main() {
	word, err := getWord()
	// terminate if cant get word
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(2)
	}

	var h Hangman
	// repeating "_"
	// setting capacity and length as per word length
	h.Placeholder = make([]string, len(word), len(word))
	for i := 0; i < len(h.Placeholder); i++ {
		h.Placeholder[i] = "_"
	}
	h.Chances = len(word)
	h.Entries = make(map[string]bool)

	err = h.start(word)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(2)
	}
}
