package main

// The Dependency Inversion Principle states that high-level modules should not depend on low-level modules.
// Both should depend on abstractions.
// Abstractions should not depend on details.
// Details should depend on abstractions.

// In the good example, BusinessLogic depends on the Database interface rather than a specific implementation.
// This allows us to easily switch between different database types without changing the BusinessLogic code.

// ----------------------------------------------------------------------------

// BAD example: violates the Dependency Inversion Principle
type MySQL struct{}

func (m *MySQL) Query(sql string) Result {
	// Perform MySQL query
}

type BusinessLogic struct {
	DB *MySQL
}

func (b *BusinessLogic) ProcessData() {
	result := b.DB.Query("SELECT * FROM users")
	// Process result
}

// ----------------------------------------------------------------------------

// GOOD example: follows the Dependency Inversion Principle
type Database interface {
	Query(string) Result
}

type MySQL struct{}

func (m *MySQL) Query(sql string) Result {
	// Perform MySQL query
}

type PostgreSQL struct{}

func (p *PostgreSQL) Query(sql string) Result {
	// Perform PostgreSQL query
}

type BusinessLogic struct {
	DB Database
}

func (b *BusinessLogic) ProcessData() {
	result := b.DB.Query("SELECT * FROM users")
	// Process result
}

// ----------------------------------------------------------------------------

// This example demonstrates the Dependency Inversion Principle by:

// Defining a Database interface that abstracts the query operation.
// Implementing this interface for different database types (MySQL and PostgreSQL).
// Having BusinessLogic depend on the Database interface rather than a concrete implementation.

// The benefits of this approach:

// BusinessLogic is decoupled from specific database implementations.
// It works with any type that satisfies the Database interface.
// We can easily switch between different database types without changing BusinessLogic.
// It's easier to test BusinessLogic by using a mock database that implements the Database interface.
// New database types can be added without affecting existing code.

// This design inverts the traditional dependency relationship.
// Instead of high-level modules depending on low-level modules, both depend on abstractions.
// This makes the system more flexible and easier to maintain and test.
