package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// sliceReader just read a variable list of integers and return a slice with all of them. x breaks the capture
func sliceReader() []int {
	l := make([]int, 0)
	for {
		fmt.Printf("Type the next number on the table or X to finish: ")
		buffer := bufio.NewScanner(os.Stdin)
		var reader string
		if buffer.Scan() {
			reader = buffer.Text()
		}
		if strings.ToLower(reader) == "x" {
			break
		}
		n, err := strconv.Atoi(reader)
		if err != nil {
			fmt.Println("Entered data can't be converted to integer value")
		} else {
			l = append(l, n)
		}
	}
	return l
}

// orderSlice is the goroutine that will do a bubble sort with the slice received and return it through a channel
func orderSlice(sli []int, c chan []int, wg *sync.WaitGroup) {
	amountOfLoops := 1
	for generalLoop := 0; generalLoop < amountOfLoops; generalLoop++ {
		needMoreLoops := false
		for bubbleI := 0; bubbleI < len(sli)-1; bubbleI++ {
			if sli[bubbleI] > sli[bubbleI+1] {
				sli[bubbleI], sli[bubbleI+1] = sli[bubbleI+1], sli[bubbleI]
				needMoreLoops = true
			}
			if needMoreLoops {
				amountOfLoops++
				needMoreLoops = false
			}
		}
	}
	c <- sli
	wg.Done()
}

func main() {
	var waitG sync.WaitGroup   // wait group
	c := make(chan []int)      // communication channel for the goroutines
	stppr := 0                 // stepper (marks the beginning of the slice passed to each goroutine)
	list := sliceReader()      // reads integers inputted from the user
	step := len(list) / 4      // Amount of itens passed to goroutine (for the first three at least, fourth receives everything else)
	orderedL := make([]int, 0) // empty slice tha receive the combination of slices treated by the goroutines
	if len(list) > 4 {         // if user input has less than four ints, no use spreading though four goroutines
		waitG.Add(4)
		for i := 0; i < 4; i++ {
			if i > 2 { // for the equal pieces
				go orderSlice(list[stppr:], c, &waitG)
			} else {
				go orderSlice(list[stppr:stppr+step], c, &waitG)
			}
			stppr = stppr + step
			t := <-c
			orderedL = append(orderedL, t...)

		}
	} else {
		waitG.Add(1)
		go orderSlice(list, c, &waitG)
		t := <-c
		orderedL = append(orderedL, t...)
	}
	waitG.Wait()
	fmt.Println(orderedL)
}
