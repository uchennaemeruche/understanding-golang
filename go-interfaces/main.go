package main

import "fmt"

func main() {
	PrintCars()
}

func PrintCars() {

	camry07 := Toyota{color: "red", model: "Camry 2007", speed: 70}
	highlander010 := Toyota{color: "blue", model: "Highlander 2010", speed: 120}

	accord011 := Honda{color: "Gray", model: "Accord 2011", speed: 110}
	pilot018 := Honda{color: "white", model: "Pilot 2018", speed: 150}

	cars := []Car{camry07, highlander010, accord011, pilot018}

	for _, car := range cars {
		fmt.Println(car.Drive())
		fmt.Println(car.Park())
		fmt.Println(car.Reverse())
	}

}
