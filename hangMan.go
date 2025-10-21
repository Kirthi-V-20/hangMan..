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
	if currentState.ChancesRemaining > 0 && isContainletter && !isAlreadyGuessed {
		currentState = HangMan{
			SecretWord:       currentState.SecretWord,
			Guesses:          append(currentState.Guesses, user_Input),
			CorrectGuesses:   append(currentState.CorrectGuesses, user_Input),
			ChancesRemaining: currentState.ChancesRemaining,
		}
	}
	if currentState.ChancesRemaining > 0 && !isContainletter && !isAlreadyGuessed {
		currentState = HangMan{
			SecretWord:       currentState.SecretWord,
			Guesses:          append(currentState.Guesses, user_Input),
			CorrectGuesses:   currentState.CorrectGuesses,
			ChancesRemaining: currentState.ChancesRemaining - 1,
		}
	}
	return currentState
}

func getUserInput(s string) byte {
	fmt.Print(s)
	reader := bufio.NewReader(os.Stdin)
	letter, _ := reader.ReadByte()
	reader.ReadByte()
	return letter
}

func checkWon(game HangMan) bool {
	for _, ch := range game.SecretWord {
		if !bytes.Contains(game.CorrectGuesses, []byte{byte(ch)}) {
			return false
		}
	}
	return true
}

func displayWord(state HangMan) string {
	display := ""

	for _, ch := range state.SecretWord {
		if bytes.Contains(state.CorrectGuesses, []byte{byte(ch)}) {
			display += string(ch) + " "
		} else {
			display += "_ "
		}
	}

	return strings.TrimSpace(display)
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
