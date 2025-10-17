package main

import (
	"strings"
	"testing"
)

func TestSecretWordNoCapitals(t *testing.T) {
	wordList := "/usr/share/dict/words"
	secretWord := getSecretWord(wordList)
	if secretWord != strings.ToLower(secretWord) {
		t.Errorf("Should not get words with capital letters. Got %s", secretWord)
	}
}

func TestSecretWordLength(t *testing.T) {
	wordList := "/usr/share/dict/words"
	secretWord := getSecretWord(wordList)
	if len(secretWord) < 6 {
		t.Errorf("Expected word length 6 or greater than 6, but got %q (length %d)", secretWord, len(secretWord))
	}
}

func TestSecretWordNopunctuation(t *testing.T) {
	wordList := "/usr/share/dict/words"
	secretWord := getSecretWord(wordList)
	if !isAllLetters(secretWord) {
		t.Errorf("Should not get words with punctuations. Got %s", secretWord)
	}
}
