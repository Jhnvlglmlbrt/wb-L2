package main

import "fmt"

// общий интерфейс для всех типов посетителей c методами, которые позволят добавлять функционал для кругов и квадратов
type Visitor interface {
	visitForCircle(*Circle)
	visitForSquare(*Square)
}

// метод принятия посетителей | абстрактный (класс) интерфейс для объектов, которые могут принимать посетителей
type Shape interface {
	accept(Visitor)
}

// конкретные типы реализуют методы принятия посетителя
// цель вызывать тот метод посетителя, который соответствует этому типу,
// так посетитель узнает с каким именно типом он работает
type Circle struct {
	radius float64
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

type Square struct {
	sideLength float64
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

// тип concrete visitor (посетитель)
type AreaCalculator struct {
	totalArea float64
}

// concrete visitors реализуют особенное поведение для всех типов, которые можно подать через методы интерфейса посетителя
func (a *AreaCalculator) visitForCircle(c *Circle) {
	area := 3.14 * c.radius * c.radius
	a.totalArea += area
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	area := s.sideLength * s.sideLength
	a.totalArea += area
}

func main() {
	shapes := []Shape{
		&Circle{radius: 5},
		&Square{sideLength: 4},
		&Circle{radius: 3},
	}

	areaVisitor := &AreaCalculator{}

	// в цикле все объекты принимают посетителя, что приводит к вызову соответствующих методов
	for _, shape := range shapes {
		shape.accept(areaVisitor)
	}

	fmt.Println("Total area of all shapes:", areaVisitor.totalArea)
}
