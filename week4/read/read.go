package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type name struct {
	fname string
	lname string
}

func main() {
	var userInput string
	scanBuffer := bufio.NewScanner(os.Stdin)
	fmt.Printf("Inform data file to be used: ")
	if scanBuffer.Scan() {
		userInput = scanBuffer.Text()
	}
	file, err := os.Open(userInput)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the data file: %s", userInput))
	}
	defer file.Close()
	sliceOfNames := make([]name, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp := name{}
		temp.fname = strings.Split(scanner.Text(), " ")[0]
		temp.lname = strings.Split(scanner.Text(), " ")[1]
		sliceOfNames = append(sliceOfNames, temp)
	}
	for _, v := range sliceOfNames {
		fmt.Println(v.fname, v.lname)
	}
}
