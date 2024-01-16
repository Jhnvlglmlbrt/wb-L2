package main

import "fmt"

type State interface {
	Handle()
}

type Context struct {
	state State
}

func (c *Context) setState(state State) {
	c.state = state
}

func (c *Context) request() {
	c.state.Handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("Handling request in State A")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("Handling request in State B")
}

func main() {
	ctx := new(Context)

	ctx.setState(&ConcreteStateA{})
	ctx.request()

	ctx.setState(&ConcreteStateB{})
	ctx.request()
}
