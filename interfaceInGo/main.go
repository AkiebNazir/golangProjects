package main

import "fmt"

type Calulate interface {
	Area()
}

type Circle struct {
	Radius float32
}

func (c *Circle) Area() {
	fmt.Printf("Area of a circle is %.2f\n", 3.14*c.Radius*c.Radius)
}

type Rectangle struct {
	Length float32
	Width  float32
}

func (r *Rectangle) Area() {
	fmt.Printf("Area of a rectangle is: %.2f\n", r.Length*r.Width)
}

func main() {

	var circle = Circle{
		Radius: 10,
	}
	var rect = Rectangle{
		Length: 20,
		Width:  40,
	}

	circle.Area()
	rect.Area()

}
