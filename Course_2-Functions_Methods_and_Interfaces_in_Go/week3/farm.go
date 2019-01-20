package main

import "fmt"

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	for {
		inputAnimal := ""
		inputAction := ""
		fmt.Print(">")
		fmt.Scanf("%s %s", &inputAnimal, &inputAction)
		switch inputAnimal {
		case "cow":
			selectAction(inputAction, cow)
		case "snake":
			selectAction(inputAction, snake)
		case "bird":
			selectAction(inputAction, bird)
		default:
			fmt.Println("Animal not in the database...")
		}
	}
}

//Class Animal
type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

//Function selectAction to decrease complexity of main function
func selectAction(input string, animal Animal) {
	switch input {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Action not in the database...")
	}
}
