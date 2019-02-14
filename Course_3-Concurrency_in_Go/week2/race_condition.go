// In this example, where both goroutines, named hello and itsMe exist. There is no garantee
// that they will in fact execute in the order of the music. This is caused because Go don't
// garantee in this execution that the goroutines will be executed in the same order they are
// called. This will cause the code to sometime show:
// > Hello...
// > It's me you are looking for
//
// And sometimes the code can show:
// > It's me you are looking for
// > Hello...
//
// This happens because of race condition. The Goroutines will try to execute as fast as possible
// which not necessarily will match the order in which they are written in the code.

package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello...")
}

func itsMe(musicLine string) {
	fmt.Println(musicLine)
}

func main() {
	go hello()
	go itsMe("It's me you are looking for...")
	time.Sleep(time.Second)
}
