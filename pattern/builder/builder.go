package main

import "fmt"

type Director struct {
	builder Builder
}

type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
	Show() *Product
}

type ProductBuilder struct {
	product *Product
}

type Product struct {
	Content string
}

func (b *ProductBuilder) MakeHeader(str string) {
	b.product.Content += "<header>" + str + "</header>\n"
}

func (b *ProductBuilder) MakeBody(str string) {
	b.product.Content += "<article>" + str + "</article>\n"
}

func (b *ProductBuilder) MakeFooter(str string) {
	b.product.Content += "<footer>" + str + "</footer>\n"
}

func (b *ProductBuilder) Show() *Product {
	return b.product
}

func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func NewConcreteBuilder() *ProductBuilder {
	return &ProductBuilder{
		product: &Product{},
	}
}

func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

func main() {
	builder := NewConcreteBuilder()
	director := NewDirector(builder)
	director.Construct()

	result := builder.Show()
	fmt.Println(result.Content)
}
