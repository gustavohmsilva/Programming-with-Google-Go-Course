package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		//cmd: newanimal or query
		//arg1: arbritary name
		//arg3: cow, bird or snake for newanimal; eat, move or speak for query
		cmd := ""
		arg1 := ""
		arg2 := ""
		fmt.Print(">")
		fmt.Scanf("%s %s %s", &cmd, &arg1, &arg2)
		cmd, arg1, arg2 = strings.ToLower(cmd), strings.ToLower(arg1), strings.ToLower(arg2)
		switch cmd {
		case "newanimal":
			if arg2 == "cow" || arg2 == "bird" || arg2 == "snake" {
				fmt.Println("valid new animal!")
			} else {
				fmt.Println("Invalid new animal! Accepted animals are:\n\t1. cow\n\t2. bird\n\t3. snake")
			}
		case "query":
			if arg2 == "eat" || arg2 == "move" || arg2 == "speak" {
				fmt.Println("valid query!")
			} else {
				fmt.Println("Invalid Query! Accepted queries are:\n\t1. eat\n\t2. move\n\t3. speak")
			}
		default:
			fmt.Printf("\"%s %s %s\" wasn't recognized as a valid command...\n", cmd, arg1, arg2)
		}
	}
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow struct
type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

// Bird struct
type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

// Snake struct
type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}
