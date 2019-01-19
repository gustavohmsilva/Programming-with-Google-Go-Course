package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	accel := GetF64AndValidate("Acceleration")
	initVel := GetF64AndValidate("Initial Velocity")
	initDisp := GetF64AndValidate("Initial Displacement")
	displacement := GenDisplaceFn(accel, initVel, initDisp)
	for {
		var rawInput string
		fmt.Print("Time/Duration: ")
		fmt.Scanln(&rawInput)
		duration, err := strconv.ParseFloat(rawInput, 64)
		if err == nil {
			fmt.Println(displacement(duration))
		}

	}

}

func GenDisplaceFn(acc, initV, initD float64) func(t float64) float64 {
	rFunc := func(duration float64) float64 {
		return 0.5*acc*math.Pow(duration, 2.0) + initV*duration + initD
	}
	return rFunc
}

func GetF64AndValidate(nameOfValue string) float64 {
	for {
		var rawInput string
		fmt.Printf("%s: ", nameOfValue)
		fmt.Scanln(&rawInput)
		converted, err := strconv.ParseFloat(rawInput, 64)
		if err == nil {
			return converted
		}
		fmt.Println("Not a floating point value! Try again!")
	}
}
