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
