package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	l := ReadSliceOfInts(10)
	BubbleSort(l)
	fmt.Println(l)

}

// BubbleSort does exactly what the name says. Receives a slice.
func BubbleSort(receivedList []int) {
	amountOfLoops := 1
	for generalLoop := 0; generalLoop < amountOfLoops; generalLoop++ {
		needMoreLoops := false
		for bubbleI := 0; bubbleI < len(receivedList)-1; bubbleI++ {
			if receivedList[bubbleI] > receivedList[bubbleI+1] {
				Swap(receivedList, bubbleI)
				needMoreLoops = true
			}
			if needMoreLoops {
				amountOfLoops++
				needMoreLoops = false
			}
		}
	}
}

// Swap exchange values if first value of slice is bigger than second.
func Swap(sliceToSort []int, i int) {
	sliceToSort[i], sliceToSort[i+1] = sliceToSort[i+1], sliceToSort[i]
}

// ReadSliceOfInts receives from the user a slice of values and returns it to the program.
func ReadSliceOfInts(qtLoops int) []int {
	listagem := make([]int, 0)
	for i := 0; i < qtLoops; i++ {
		fmt.Printf("Type the next number on the table or X to finish: ")
		leituraBf := bufio.NewScanner(os.Stdin)
		var leitura string
		if leituraBf.Scan() {
			leitura = leituraBf.Text()
		}
		if strings.ToLower(leitura) == "x" {
			break
		}
		n, err := strconv.Atoi(leitura)
		if err != nil {
			fmt.Println("Entered data can't be converted to integer value")
			i--
		} else {
			listagem = append(listagem, n)
		}
	}
	return listagem
}
