package main

import "math"

// The Open Closed Principle states that software entities (classes, modules, functions, etc.) should be open for extension but closed for modification.

// In the good example, we define a Shape interface and implement it for different shapes.
// This allows us to add new shapes without modifying the CalculateArea function.

// ----------------------------------------------------------------------------

// BAD example: violates the Open Closed Principle
type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func CalculateArea(shapes ...interface{}) float64 {
	var area float64
	for _, shape := range shapes {
		switch s := shape.(type) {
		case Rectangle:
			area += s.Width * s.Height
		case Circle:
			area += math.Pi * s.Radius * s.Radius
		}
	}
	return area
}

// ----------------------------------------------------------------------------

// GOOD example: follows the Open Closed Principle
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func CalculateArea(shapes ...Shape) float64 {
	var area float64
	for _, shape := range shapes {
		area += shape.Area()
	}
	return area
}

// ----------------------------------------------------------------------------

// This example demonstrates the Open Closed Principle by:

// Defining a Shape interface with an Area() method. This is the abstraction that allows for extension.
// Implementing the Shape interface for different types (Rectangle and Circle).
// Creating a CalculateArea function that works with any Shape.

// The key benefits are:

// To add a new shape (e.g., Triangle), we don't need to modify CalculateArea.
// We just need to implement the Shape interface for the new type.
// The CalculateArea function is "closed for modification" but "open for extension" through the Shape interface.

// This design allows the system to be easily extended with new shapes without changing existing code, reducing the risk of introducing bugs in working code.
