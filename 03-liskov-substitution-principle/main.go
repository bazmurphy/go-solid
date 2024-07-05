package main

import "fmt"

// The Liskov Substitution Principle states that objects of a superclass should be replaceable with objects of its subclasses without affecting the correctness of the program.

// In the good example, we've separated flying and walking behaviors into different interfaces.
// This ensures that we don't violate Liskov Substitution Principle by having a method that doesn't make sense for a subtype.

// ----------------------------------------------------------------------------

// BAD example: violates the Liskov Substitution Principle
type Bird interface {
	Fly()
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow flying")
}

type Ostrich struct{}

func (o Ostrich) Fly() {
	panic("Ostriches can't fly!")
}

// ----------------------------------------------------------------------------

// GOOD example: follows the Liskov Substitution Principle
type FlyingBird interface {
	Fly()
}

type WalkingBird interface {
	Walk()
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow flying")
}

type Ostrich struct{}

func (o Ostrich) Walk() {
	fmt.Println("Ostrich walking")
}

// ----------------------------------------------------------------------------

// This example demonstrates the Liskov Substitution Principle by:

// Creating separate interfaces for different behaviors: FlyingBird and WalkingBird.
// Implementing these interfaces only for the birds that can perform those actions.

// The benefits of this approach:

// It prevents nonsensical operations like calling Fly() on an ostrich.
// Any function expecting a FlyingBird can safely use a Sparrow without unexpected behavior.
// It accurately models the real-world capabilities of different birds.

// This design ensures that subtypes (specific birds) can be used interchangeably with their supertypes (bird interfaces) without breaking the program's logic.
