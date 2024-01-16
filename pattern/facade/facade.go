package main

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import (
	"strings"
)

func NewManFacade() *ManFacade {
	return &ManFacade{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

type ManFacade struct {
	house *House
	tree  *Tree
	child *Child
}

func (m *ManFacade) Todo() string {
	result := []string{
		m.house.Build(),
		m.tree.Grow(),
		m.child.Born(),
	}
	return strings.Join(result, "\n")
}

type House struct {
}

func (h *House) Build() string {
	return "Build house"
}

type Tree struct {
}

func (t *Tree) Grow() string {
	return "Tree grow"
}

type Child struct {
}

func (c *Child) Born() string {
	return "Child born"
}

func main() {
	man := NewManFacade()

	result := man.Todo()
	println(result)
}
