package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wordList := loadWords("wordlist.txt") // Load 5-letter words from a file
	solver := NewWordleSolver(wordList)
	solver.Solve()
}

// Load words from a file
func loadWords(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToUpper(scanner.Text())
		if len(word) == 5 {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return words
}

// WordleSolver struct
type WordleSolver struct {
	wordList []string
}

// NewWordleSolver creates a new WordleSolver
func NewWordleSolver(wordList []string) *WordleSolver {
	return &WordleSolver{wordList: wordList}
}

// Solve starts the Wordle solving process
// Assuming feedback is a string like "gyygg" where 'g' is green, 'y' is yellow, and 'b' is gray
func (ws *WordleSolver) Solve() {
	var guess string
	for attempts := 0; attempts < 6; attempts++ {
		if attempts == 0 {
			guess = ws.wordList[0] // Starting with the first word in the list, can be randomized
		} else {
			guess = ws.makeGuess()
		}
		fmt.Printf("Guess #%d: %s\n", attempts+1, guess)

		var feedback string
		fmt.Println("Enter feedback (g for green, y for yellow, b for gray):")
		fmt.Scanln(&feedback)

		if feedback == "ggggg" {
			fmt.Println("Solved!")
			return
		}

		ws.refineWordList(guess, feedback)
	}
	fmt.Println("Failed to solve within 6 attempts.")
}

// makeGuess selects the next guess based on the current state of the word list
func (ws *WordleSolver) makeGuess() string {
	// For simplicity, just returning the first word in the list
	// Advanced logic would analyze the current word list to make a smarter guess
	return ws.wordList[0]
}

// refineWordList narrows down the word list based on the feedback
func (ws *WordleSolver) refineWordList(guess string, feedback string) {
	var newWordList []string

	for _, word := range ws.wordList {
		if matchesFeedback(word, guess, feedback) {
			newWordList = append(newWordList, word)
		}
	}

	ws.wordList = newWordList
}

// matchesFeedback checks if a word matches the given feedback for a guess
func matchesFeedback(word, guess, feedback string) bool {
	if len(word) != len(guess) || len(guess) != len(feedback) {
		return false
	}

	for i := range guess {
		switch feedback[i] {
		case 'g':
			if word[i] != guess[i] {
				return false
			}
		case 'y':
			if !strings.Contains(word, string(guess[i])) || word[i] == guess[i] {
				return false
			}
		case 'b':
			if strings.Contains(word, string(guess[i])) {
				return false
			}
		}
	}

	return true
}
