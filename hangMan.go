package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

type HangMan struct {
	SecretWord       string
	Guesses          []byte
	ChancesRemaining uint
	CorrectGuesses   []byte
}

func NewHangMan(secretWord string) HangMan {
	return HangMan{
		SecretWord:       secretWord,
		Guesses:          []byte{},
		ChancesRemaining: 7,
		CorrectGuesses:   []byte{},
	}
}

func isAllLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func getSecretWord(wordFileName string) string {
	allowedWords := []string{}
	file, err := os.Open(wordFileName)
	if err != nil {
		errMessage := fmt.Sprintf("Can't open file %s : %v\n", wordFileName, err)
		panic(errMessage)

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if word == strings.ToLower(word) && len(word) >= 6 && isAllLetters(word) {
			allowedWords = append(allowedWords, word)
		}
	}
	randomNum := rand.Intn(len(allowedWords))
	return allowedWords[randomNum]
}

func checkGuess(currentState HangMan, user_Input byte) HangMan {
	isContainletter := strings.ContainsRune(currentState.SecretWord, rune(user_Input))
	isAlreadyGuessed := bytes.Contains(currentState.Guesses, []byte{user_Input})
	if currentState.ChancesRemaining > 1 && isContainletter && !isAlreadyGuessed {
		currentState = HangMan{
			SecretWord:       currentState.SecretWord,
			Guesses:          append(currentState.Guesses, user_Input),
			CorrectGuesses:   append(currentState.CorrectGuesses, user_Input),
			ChancesRemaining: currentState.ChancesRemaining,
		}
	}
	return currentState
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
