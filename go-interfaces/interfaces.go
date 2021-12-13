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

type Honda struct {
	color string
	model string
	speed int
}

func (t Toyota) Drive() string {
	str := fmt.Sprintf("Your Toyota %v car is running on %vkm/hr", t.model, t.speed)
	return str
}

func (t Toyota) Park() string {
	elapsedTime := 10.0
	distanceTravelled := float64(t.speed) * elapsedTime
	str := fmt.Sprintf("%v Toyota %v covered %vkm", t.color, t.model, distanceTravelled)
	return str
}

func (t Toyota) Reverse() bool {
	return true
}
func (h Honda) Drive() string {
	str := fmt.Sprintf("Your Honda %v car is running on %vkm/hr", h.model, h.speed)
	return str
}

func (h Honda) Park() string {
	elapsedTime := 10.0
	distanceTravelled := float64(h.speed) * elapsedTime
	str := fmt.Sprintf("%v Honda %v covered %vkm", h.color, h.model, distanceTravelled)
	return str
}

func (h Honda) Reverse() bool {
	return true
}
