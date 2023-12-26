/*
In Go, the Factory Pattern is used to create objects without exposing
the instantiation logic to the client. Here's an example of how you might
implement a simple Factory Pattern:
*/

package main

import (
	"fmt"
)

// Interface for the product
type Shape interface {
	Draw() string
}

// Concrete implementations of Shape interface
type Circle struct{}

func (c Circle) Draw() string {
	return "Circle drawn"
}

type Square struct{}

func (s Square) Draw() string {
	return "Square drawn"
}

// Factory function to create shapes
func CreateShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return Circle{}
	case "square":
		return Square{}
	default:
		return nil
	}
}
func main() {
	// Create shapes using the factory function
	circle := CreateShape("circle")
	square := CreateShape("square")
	// Use the created shapes
	fmt.Println(circle.Draw()) // Output: Circle drawn
	fmt.Println(square.Draw()) // Output: Square drawn
}

/*
Circle drawn
Square drawn
*/
