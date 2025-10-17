package main

import (
	"bufio"
	"fmt"
	"os"
)

func getSecretWord(wordFileName string) string {
	file, err := os.Open(wordFileName)
	if err != nil {
		panic(fmt.Sprintf("The file could not be opened: %v", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return "Hii"
}
func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
