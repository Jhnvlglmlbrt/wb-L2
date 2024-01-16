package main

import (
	"fmt"
	"log"
)

type form string

const (
	C form = "Circle"
	R form = "Rectangle"
)

type Creator interface {
	CreateShape(form form) Shape
}

type Shape interface {
	Draw() string
}

type ConcreteCreator struct{}

// NewCreator is the ConcreteCreator constructor.
func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (s *ConcreteCreator) CreateShape(form form) Shape {
	var shape Shape

	switch form {
	case C:
		shape = &ConcreteCircle{string(form)}
	case R:
		shape = &ConcreteRectangle{string(form)}
	default:
		log.Fatalln("Unknown Form")
	}

	return shape
}

type ConcreteCircle struct {
	form string
}

func (c *ConcreteCircle) Draw() string {
	return c.form
}

type ConcreteRectangle struct {
	form string
}

func (r *ConcreteRectangle) Draw() string {
	return r.form
}

func main() {
	assert := []string{"C", "R"}

	factory := NewCreator()

	shapes := []Shape{
		factory.CreateShape(C),
		factory.CreateShape(R),
	}

	for i, shape := range shapes {
		if form := shape.Draw(); form != assert[i] {
			fmt.Printf("Expect action to %s, have %s.\n", assert[i], form)
		}
	}
}
