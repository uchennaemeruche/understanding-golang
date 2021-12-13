package main

import "fmt"

func main() {
	PrintCars()
}

func PrintCars() {
	cars := []Toyota{
		{color: "red", model: "2007", speed: 70},
		{color: "blue", model: "2010", speed: 120},
	}

	for _, car := range cars {
		fmt.Println(car.Drive())
		fmt.Println(car.Park())
		fmt.Println(car.Reverse())
	}

}
