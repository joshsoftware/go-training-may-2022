package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "elephant"
	//loookup for entries made by user
	entries := map[string]bool{}
	//list of "_" corrosponding to number of letters in word. [_ _ _ _]
	placeholders := []string{}
	//get length of the word and initialize the slize with each element as "_"
	for i := 0; i < len(word); i++ {
		placeholders = append(placeholders, "_")
	}
	fmt.Println("entries  are -->", entries)
	chances := 8
	for {
		//If user gueses a wrong input then they loss a chance.
		userInput := strings.Join(placeholders, "")
		if chances == 0 && userInput != word {
			fmt.Println("Game over!! Please try again.")
			break
		}
		//evaluate a win
		if userInput == word {
			fmt.Print("You win")
			break
		}
		fmt.Println("chances are --> ", chances)
		keys := get_keys(entries)
		fmt.Println("Previous Guesses --> ", keys)
		fmt.Println("Guess the word")
		//scan the input
		str := ""
		fmt.Scanln(&str)
		if str == word {
			fmt.Print("You win")
			break
		}
		//checking if the entry for the letter is already present in the map
		_, ok := entries[str]
		//if the letter exists, skip the iteration
		if ok {
			continue
		}
		//if the letter doen't exists,update the entries map
		entries[str] = true
		exists := false
		//common approach of iterating the word
		//check if letter exists in the word
		//tmp := strings.Split(word, "")
		/* for i, v := range tmp {
			//if exists, update the placeholder indices
			if v == str {
				exists = true
				placeholders[i] = str
			}
		} */
		//Correct approach of iterating the word
		for i, v := range word {
			if str == string(v) {
				exists = true
				placeholders[i] = str
			}
		}
		//if doesn't exists, decrement the chances
		if !exists {
			fmt.Println("Letter does not exists")
			chances -= 1
		}
		fmt.Println("placeholders -->", placeholders)
	}

}

func get_keys(entries map[string]bool) (keys []string) {
	for k, _ := range entries {
		keys = append(keys, k)
	}
	return
}
