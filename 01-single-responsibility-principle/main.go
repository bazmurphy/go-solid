package main

// The Single Responsibility Principle states that a class or module should have only one reason to change.
// In other words, it should have only one job or responsibility.

// In the good example, we've separated the responsibilities into different structures.
// The User struct only holds data, while UserRepository handles database operations and EmailService handles email-related tasks.

// ----------------------------------------------------------------------------

// BAD example: violates the Single Responsibility Principle
type User struct {
	Name  string
	Email string
}

func (u *User) SaveToDatabase() error {
	// Save user to database
}

func (u *User) SendWelcomeEmail() error {
	// Send welcome email
}

// ----------------------------------------------------------------------------

// GOOD example: follows the Single Responsibility Principle
type User struct {
	Name  string
	Email string
}

type UserRepository struct {
	// Database connection details
}

func (r *UserRepository) Save(user User) error {
	// Save user to database
}

type EmailService struct {
	// Email service configuration
}

func (e *EmailService) SendWelcomeEmail(user User) error {
	// Send welcome email
}

// ----------------------------------------------------------------------------

// This example demonstrates the Single Responsibility Principle by separating concerns:

// The User struct is now only responsible for holding user data.
// It doesn't handle any operations.

// UserRepository is responsible for database operations.
// It has a Save method that takes a User and saves it to the database.
// This isolates all database-related logic.

// EmailService handles email-related tasks.
// It has a SendWelcomeEmail method that sends emails to users.

// By separating these responsibilities, we achieve several benefits:

// If we need to change how users are saved (e.g., switching databases), we only modify UserRepository.
// If email sending logic changes, we only update EmailService.
// The User struct can be used in various contexts without carrying unnecessary methods.

// This separation makes the code more maintainable and allows each component to evolve independently.
