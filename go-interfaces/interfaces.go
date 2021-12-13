package main

import "fmt"

type Car interface {
	Drive() string
	Park() string
	Reverse() bool
}

type Toyota struct {
	color string
	model string
	speed int
}

func (t Toyota) Drive() string {
	str := fmt.Sprintf("Your %v Toyota car is running on %vkm/hr", t.model, t.speed)
	return str
}

func (t Toyota) Park() string {
	elapsedTime := 10.0
	distanceTravelled := float64(t.speed) * elapsedTime
	str := fmt.Sprintf("%v %v Toyota covered %vkm", t.color, t.model, distanceTravelled)
	return str
}

func (t Toyota) Reverse() bool {
	return true
}
