package main

import (
	"fmt"
)

func getSecretWord(wordFileName string) string {
	return "Hii"
}
func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
