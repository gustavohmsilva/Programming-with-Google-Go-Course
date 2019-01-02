package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputSentence string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputSentence = scanner.Text()
	}
	inputSentence = strings.ToLower(inputSentence)
	if strings.HasPrefix(inputSentence, "i") && strings.HasSuffix(inputSentence, "n") && strings.Contains(inputSentence, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
