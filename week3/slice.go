package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	valuesList := make([]int, 3)
	var userInput string
	for i := 0; i > -1; i++ {
		scanBuffer := bufio.NewScanner(os.Stdin)
		if scanBuffer.Scan() {
			userInput = scanBuffer.Text()
		}
		userInputInt, err := strconv.Atoi(userInput)
		if err != nil {
			break
		} else {
			if i < 3 {
				valuesList[i] = userInputInt
				t := make([]int, len(valuesList))
				copy(t, valuesList)
				sort.Ints(t)
				fmt.Println(t)
			} else {
				valuesList = append(valuesList, userInputInt)
				sort.Ints(valuesList)
				fmt.Println(valuesList)
			}

		}
	}
}
