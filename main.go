package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/tjarratt/babble"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

var ErrWrongGuess = errors.New("incorrect guess, try again")

var numberOfGuesses int = 5
var guessIndex []int
var guessedLetters []string

func Input() rune {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Guess a letter. Note: Only the first letter is read.")
	fmt.Printf("You have guessed the letters %v\n", guessedLetters)
	input, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	guessedLetters = append(guessedLetters, string(input))

	return input
}

func GenerateWord() string {
	babbler := babble.NewBabbler()
	babbler.Count = 1
	return strings.ToLower(babbler.Babble())
}

func ParseGuess(guess rune, word string) ([]int, error) {
	return FindLetters(guess, word)
}

func ConvertWord(s string) string {
	return strings.Repeat("_", utf8.RuneCountInString(s))
}

func FindLetters(r rune, word string) ([]int, error) {
	for i, c := range word {
		if r == c {
			guessIndex = append(guessIndex, i)
		}
	}
	if len(guessIndex) == 0 {
		return nil, ErrWrongGuess
	}
	return guessIndex, nil
}

func ReplaceLetters(s []int, g rune, underscores string) string {
	out := []rune(underscores)
	for _, c := range s {
		out[c] = g
	}
	return string(out)
}

func Initialize() (string, string) {
	word := GenerateWord()
	underscored := ConvertWord(word)
	return word, underscored
}

func GameLogic(word string, underscored string) {
	if numberOfGuesses == 0 {
		log.Fatalf("You are out of guesses. Game over. The correct words was %s", word)
	}
	DetermineWin(underscored)

	fmt.Println(underscored)

	guess := Input()
	indexSlice, err := ParseGuess(guess, word)

	if err != nil {
		fmt.Println(err)
		numberOfGuesses -= 1
		fmt.Printf("You have %d tries left.\n", numberOfGuesses)
		GameLogic(word, underscored)
	}

	replaced := ReplaceLetters(indexSlice, guess, underscored)
	fmt.Println(replaced)

	guessIndex = guessIndex[:0]

	GameLogic(word, replaced)
}

func DetermineWin(underscored string) {
	if strings.Index(underscored, "_") == -1 {
		log.Fatal("Congratulations. You have won the game!")
	}
}

func ParseFlags() {
	wordPtr := flag.String("d", "medium", "difficulty: easy, medium, or hard")

	flag.Parse()

	switch *wordPtr {
	case "easy":
		numberOfGuesses = 7
	case "medium":
		numberOfGuesses = 5
	case "hard":
		numberOfGuesses = 3
	}
}

func main() {
	ParseFlags()
	word, underscore := Initialize()
	GameLogic(word, underscore)
}
