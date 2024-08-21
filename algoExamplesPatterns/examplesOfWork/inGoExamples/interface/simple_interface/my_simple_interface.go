package main

import "fmt"

type ElectroCar struct {
	EngineSound string
}

func (e ElectroCar) EngineOn() {
	fmt.Println(e.EngineSound)
}

type Car struct {
	EngineSound string
}

func (c Car) EngineOn() {
	fmt.Println(c.EngineSound)
}

type FighterPlane struct {
	EngineSound string
}

func (f FighterPlane) EngineOn() {
	fmt.Println(f.EngineSound)
}

type Engine interface {
	EngineOn()
}

func main() {
	c := Car{
		EngineSound: "WRRRRRRRR",
	}

	e := ElectroCar{
		EngineSound: "shhh...",
	}

	f := FighterPlane{
		EngineSound: "WRWWRRWRWRWWRRWRWWRWR WW3 IS COMING, I HOPE NO",
	}

	Engine.EngineOn(c)
	Engine.EngineOn(e)
	Engine.EngineOn(f)
}
