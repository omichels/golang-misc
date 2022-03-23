package main

import (
	"fmt"
)

type speed float64

const (
	MPH speed = 1
	KPH       = 1.60934
)

type color string

const (
	BlueColor  color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type wheels string

const (
	SportsWheels wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(c color) *Builder
	Wheels(w wheels) *Builder
	TopSpeed(s speed) *Builder
	Build() *Car
}

type Interface interface {
	Drive() error
	Stop() error
}

type Car struct {
	topSpeed speed
	color    string
	wheels   string
}

func (a *Car) String() string {
	return fmt.Sprintf("The Car specs: %f\t%s\t%s", a.topSpeed, a.color, a.wheels)
}

type CarColorBuilder struct {
	Builder
}

func (cb *CarColorBuilder) Color(c color) *CarColorBuilder {
	fmt.Println(cb.Builder)
	return cb
}

type AssemblyDirector struct {
	b *Builder
}

func NewAssemblyDirector() *AssemblyDirector {
	c := &AssemblyDirector{}
	return c
}

func (a *AssemblyDirector) Color(c color) *Builder {

	return a.b
}

func (a *ConcreteBuilder) Wheels(w wheels) *Builder {

	return a.b
}
func (a *ConcreteBuilder) TopSpeed(s speed) *Builder {

	return a.b
}

func (a *ConcreteBuilder) Build() *Car {

	fmt.Println("Build")
	return nil
}

func main() {

	fmt.Println()
	//assembly := NewBuilder().Paint(RedColor)
	//assembly := NewAssembly().Color(BlueColor).TopSpeed(10).Wheels(SteelWheels).Build()
	assembly := NewCarBuilder()
	aCar := assembly.Color(BlueColor).Wheels(SteelWheels).TopSpeed(100).Build()
	fmt.Println(aCar)
	//
	//familyCar := assembly.Wheels(SportsWheels).TopSpeed(50 * car.MPH).Build()
	//familyCar.Drive()
}
