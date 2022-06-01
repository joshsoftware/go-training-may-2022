package practice

// getting babble repository for random word generation
import "github.com/tjarratt/babble"

func ReturnRandomWord() string {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	return babbler.Babble()
}
