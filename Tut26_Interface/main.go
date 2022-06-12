package main

import (
	"fmt"
	"math"
)

type rect struct {
	width  float32
	length float32
}

type circle struct {
	radius float32
}

type geometry interface {
	area() float32
	perim() float32
}

func (c *circle) area() float32 {

	return math.Pi * c.radius * c.radius
}

func (r *rect) area() float32 {

	return r.width * r.length
}

func (c *circle) perim() float32 {

	return 2 * math.Pi * c.radius
}

func measure(g geometry) {

	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	mycir := circle{radius: 5}

	//myrect := rect{width: 30, length: 20}

	measure(&mycir)
	//fmt.Println(myrect.area())

}
