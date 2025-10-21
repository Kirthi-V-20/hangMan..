package main

import (
	"os"
	"strings"
	"testing"
)

func createDictFiles(words []string) (string, error) {
	f, err := os.CreateTemp("/tmp", "hangman-dict")
	data := strings.Join(words, "\n")
	_, err = f.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}
func TestSecretWordNoCapitals(t *testing.T) {
	wordList, err := createDictFiles([]string{"Lion", "Elephant", "monkey"})
	defer os.Remove(wordList)
	if err != nil {
		t.Errorf("Couldn't create word list.Can't proceed with test:%v", err)
	}
	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("Should get 'monkey'but got %s", secretWord)
	}
}
func TestSecretWordLength(t *testing.T) {
	wordList, err := createDictFiles([]string{"Lion", "it", "monkey"})
	defer os.Remove(wordList)
	if err != nil {
		t.Errorf("Couldn't create word list.Can't proceed with test:%v", err)
	}
	secretWord := getSecretWord(wordList)
	if len(secretWord) < 6 {
		t.Errorf("Expected word length 6 or greater than 6, but got %q (length %d)", secretWord, len(secretWord))
	}
}
func TestSecretWordNopunctuation(t *testing.T) {
	wordList, err := createDictFiles([]string{"Lion's", "Elephant's", "monkey"})
	defer os.Remove(wordList)
	if err != nil {
		t.Errorf("Couldn't create word list.Can't proceed with test:%v", err)
	}
	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("Should not get words with punctuations. Got %s", secretWord)
	}
}

func TestCorrectGuess(t *testing.T) {
	SecretWord := "pineapple"
	currentState := NewHangMan(SecretWord)
	user_Input := byte('l')
	newState := checkGuess(currentState, byte(user_Input))
	expected := HangMan{
		SecretWord:       SecretWord,
		Guesses:          append(currentState.Guesses, byte(user_Input)),
		ChancesRemaining: 7,
		CorrectGuesses:   append(currentState.CorrectGuesses, byte(user_Input)),
	}
	if newState.SecretWord != expected.SecretWord {
		t.Errorf("Secreat word is modified\n")
	}
	if string(newState.Guesses) != string(expected.Guesses) {
		t.Errorf("Guess should be [l] but got %v\n", newState.Guesses)
	}
	if string(newState.CorrectGuesses) != string(expected.CorrectGuesses) {
		t.Errorf("Guess should be [l] but got %v", newState.CorrectGuesses)
	}
	if newState.ChancesRemaining != expected.ChancesRemaining {
		t.Errorf("Chances left modified!\n")
	}
}

func TestCorrectGuess1(t *testing.T) {
	secretWord := "pineapple"
	currentState := HangMan{
		SecretWord:       secretWord,
		Guesses:          []byte{'x', 'y'},
		ChancesRemaining: 5,
		CorrectGuesses:   []byte{},
	}
	user_Input := byte('p')
	newState := checkGuess(currentState, byte(user_Input))
	expected := HangMan{
		SecretWord:       secretWord,
		Guesses:          append(currentState.Guesses, byte(user_Input)),
		ChancesRemaining: 5,
		CorrectGuesses:   append(currentState.CorrectGuesses, byte(user_Input)),
	}
	if newState.SecretWord != expected.SecretWord {
		t.Errorf("Secreat word is modified\n")
	}
	if string(newState.Guesses) != string(expected.Guesses) {
		t.Errorf("Guess should be [e] but got %v\n", newState.Guesses)
	}
	if string(newState.CorrectGuesses) != string(expected.CorrectGuesses) {
		t.Errorf("Guess should be [e] but got %v", newState.CorrectGuesses)
	}
	if newState.ChancesRemaining != expected.ChancesRemaining {
		t.Errorf("Chances left modified!\n")
	}
}

func TestWrongGuess(t *testing.T) {
	secretWord := "pineapple"
	currentState := NewHangMan(secretWord)
	user_Input := byte('r')
	newState := checkGuess(currentState, byte(user_Input))
	expected := HangMan{
		SecretWord:       secretWord,
		Guesses:          append(currentState.Guesses, byte(user_Input)),
		ChancesRemaining: currentState.ChancesRemaining - 1,
		CorrectGuesses:   currentState.CorrectGuesses,
	}
	if newState.SecretWord != expected.SecretWord {
		t.Errorf("Secreat word is modified\n")
	}
	if string(newState.Guesses) != string(expected.Guesses) {
		t.Errorf("Guess should be [p] but got %v\n", newState.Guesses)
	}
	if string(newState.CorrectGuesses) != string(expected.CorrectGuesses) {
		t.Errorf("Guess should be [p] but got %v", newState.CorrectGuesses)
	}
	if newState.ChancesRemaining != expected.ChancesRemaining {
		t.Errorf("Chances left not decremented\n")
	}
}

func TestWrongGuess1(t *testing.T) {
	secretWord := "pineapple"
	currentState := HangMan{
		SecretWord:       secretWord,
		Guesses:          []byte{'x', 'y'},
		ChancesRemaining: 5,
		CorrectGuesses:   []byte{'t'},
	}
	user_Input := byte('f')
	newState := checkGuess(currentState, byte(user_Input))
	expected := HangMan{
		SecretWord:       secretWord,
		Guesses:          append(currentState.Guesses, byte(user_Input)),
		ChancesRemaining: currentState.ChancesRemaining - 1,
		CorrectGuesses:   currentState.CorrectGuesses,
	}
	if newState.SecretWord != expected.SecretWord {
		t.Errorf("Secreat word is modified\n")
	}
	if string(newState.Guesses) != string(expected.Guesses) {
		t.Errorf("Guess should be [a] but got %v\n", newState.Guesses)
	}
	if string(newState.CorrectGuesses) != string(expected.CorrectGuesses) {
		t.Errorf("Guess should be [a] but got %v", newState.CorrectGuesses)
	}
	if newState.ChancesRemaining != expected.ChancesRemaining {
		t.Errorf("Chances left not decremented\n")
	}
}

func TestAlreadyGuess(t *testing.T) {
	secretWord := "pineapple"
	currentState := NewHangMan(secretWord)
	user_Input := byte('a')
	newState := checkGuess(currentState, byte(user_Input))
	expected := HangMan{
		SecretWord:       secretWord,
		Guesses:          []byte{'a'},
		ChancesRemaining: 7,
		CorrectGuesses:   []byte{'a'},
	}
	if newState.SecretWord != expected.SecretWord {
		t.Errorf("Secreat word is modified\n")
	}
	if string(newState.Guesses) != string(expected.Guesses) {
		t.Errorf("Guess should be [e] but got %v\n", newState.Guesses)
	}
	if string(newState.CorrectGuesses) != string(expected.CorrectGuesses) {
		t.Errorf("Guess should be [e] but got %v", newState.CorrectGuesses)
	}
	if newState.ChancesRemaining != expected.ChancesRemaining {
		t.Errorf("Chances left modified!\n")
	}
}

func TestCheckWon(t *testing.T) {
	state := HangMan{
		SecretWord:       "pineapple",
		CorrectGuesses:   []byte{'p', 'i', 'n', 'e', 'a', 'p', 'l'},
		Guesses:          []byte{'p', 'i', 'n', 'e', 'a', 'p', 'l'},
		ChancesRemaining: 7,
	}
	if !checkWon(state) {

		t.Errorf("Expected true but got false")
	}
}

func TestCheckWon1(t *testing.T) {
	state := HangMan{
		SecretWord:       "pineapple",
		CorrectGuesses:   []byte{'p', 'n', 'e', 'a'},
		Guesses:          []byte{'p', 'n', 'e', 'a', 'g'},
		ChancesRemaining: 1,
	}
	if checkWon(state) {
		t.Errorf("Expected false but got true")
	}
}

func TestDisplayWord(t *testing.T) {
	state := HangMan{
		SecretWord:       "pineapple",
		CorrectGuesses:   []byte{},
		Guesses:          []byte{},
		ChancesRemaining: 7,
	}
	expected := "_ _ _ _ _ _ _ _ _"
	result := displayWord(state)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestDisplayWord1(t *testing.T) {
	state := HangMan{
		SecretWord:       "pineapple",
		CorrectGuesses:   []byte{'p', 'e'},
		Guesses:          []byte{'p', 'e'},
		ChancesRemaining: 7,
	}
	expected := "p _ _ e _ p p _ e"
	result := displayWord(state)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
