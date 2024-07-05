package main

import "fmt"

// The Interface Segregation Principle states that no client should be forced to depend on methods it does not use.
// It suggests breaking down larger interfaces into smaller, more specific ones.

// In the good example, we've segregated the interfaces so that Robot only needs to implement the Worker interface, while Human implements the more comprehensive LivingWorker interface.

// ----------------------------------------------------------------------------

// BAD example: violates the Interface Segregation Principle
type Worker interface {
	Work()
	Eat()
	Sleep()
}

type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Robot working")
}

func (r Robot) Eat() {
	// Robots don't eat, so this method is empty
}

func (r Robot) Sleep() {
	// Robots don't sleep, so this method is empty
}

// ----------------------------------------------------------------------------

// GOOD example: follows the Interface Segregation Principle
type Worker interface {
	Work()
}

type LivingWorker interface {
	Worker
	Eat()
	Sleep()
}

type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Robot working")
}

type Human struct{}

func (h Human) Work() {
	fmt.Println("Human working")
}

func (h Human) Eat() {
	fmt.Println("Human eating")
}

func (h Human) Sleep() {
	fmt.Println("Human sleeping")
}

// ----------------------------------------------------------------------------

// This example demonstrates the Interface Segregation Principle by:

// Defining a minimal Worker interface with just a Work() method.
// Creating a more comprehensive LivingWorker interface that extends Worker and adds Eat() and Sleep() methods.
// Implementing Worker for Robot, which only needs to work.
// Implementing LivingWorker for Human, which needs to work, eat, and sleep.

// The benefits of this approach:

// Robot isn't forced to implement unnecessary methods like Eat() or Sleep().
// Human can implement all necessary methods without leaving any unused.
// Client code can depend on the minimal interface it needs (Worker for tasks, LivingWorker for more complex scenarios).

// This design prevents clients from depending on methods they don't use, making the system more flexible and reducing the impact of changes.
