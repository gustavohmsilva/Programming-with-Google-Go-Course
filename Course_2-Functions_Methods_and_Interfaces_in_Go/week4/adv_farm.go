package main

import (
	"fmt"
	"strings"
)

func main() {
	listOfAnimals := make([]Animal, 0)
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
				var a Animal
				switch arg2 {
				case "cow":
					cow := Cow{arg1, "grass", "walk", "moo"}
					a = cow
				case "bird":
					bird := Bird{arg1, "worms", "fly", "peep"}
					a = bird
				case "snake":
					snake := Snake{arg1, "mice", "slither", "hss"}
					a = snake
				}
				listOfAnimals = append(listOfAnimals, a)
				fmt.Println("Sucess!")
			} else {
				fmt.Println("Invalid new animal! Accepted animals are:\n\t1. cow\n\t2. bird\n\t3. snake")
			}
		case "query":
			if arg2 == "eat" || arg2 == "move" || arg2 == "speak" {
				for _, v := range listOfAnimals {
					if v.Name() == arg1 {
						switch arg2 {
						case "eat":
							v.Eat()
						case "move":
							v.Move()
						case "speak":
							v.Speak()
						}
					}
				}
			} else {
				fmt.Println("Invalid Query! Accepted queries are:\n\t1. eat\n\t2. move\n\t3. speak")
			}
		default:
			fmt.Printf("\"%s %s %s\" wasn't recognized as a valid command...\n", cmd, arg1, arg2)
		}
	}
}

type Animal interface {
	Name() string
	Eat()
	Move()
	Speak()
}

// Cow struct
type Cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (c Cow) Name() string {
	return c.name
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
	name       string
	food       string
	locomotion string
	noise      string
}

func (b Bird) Name() string {
	return b.name
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
	name       string
	food       string
	locomotion string
	noise      string
}

func (s Snake) Name() string {
	return s.name
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
