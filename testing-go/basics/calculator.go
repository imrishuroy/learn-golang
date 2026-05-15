package main

import (
	"errors"
	"fmt"
)

/*
  TESTING IN GO

KEY POINTS:
  - Test files end with _test.go
  - Test functions start with Test and take *testing.T
  - Run tests with: go test ./...
  - Run with verbose: go test -v ./...
  - Run specific test: go test -run TestName

TEST TYPES:
  - Unit tests: Test individual functions
  - Table-driven tests: Multiple cases in one test
  - Benchmark tests: Measure performance (func BenchmarkX(b *testing.B))
  - Example tests: Documentation + verification (func ExampleX())

TESTING PACKAGE:
  t.Error(args...)      // Log error, continue
  t.Errorf(format, ...) // Formatted error, continue
  t.Fatal(args...)      // Log error, stop test
  t.Fatalf(format, ...) // Formatted error, stop test
  t.Log(args...)        // Log info (shown with -v)
  t.Skip(args...)       // Skip this test
  t.Helper()            // Mark as helper function

RUN THIS EXAMPLE:
  cd 06-testing/basics
  go test -v ./...
  go test -cover ./...
  go test -bench=. ./...

*/

// Calculator provides basic math operations
type Calculator struct {
	precision int
}

// NewCalculator creates a new calculator
func NewCalculator() *Calculator {
	return &Calculator{precision: 2}
}

// Add returns the sum of two numbers
func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference
func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product
func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient, error if dividing by zero
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// StringUtils for string operations
type StringUtils struct{}

// Reverse reverses a string
func (s *StringUtils) Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if string is a palindrome
func (s *StringUtils) IsPalindrome(str string) bool {
	return str == s.Reverse(str)
}

// User for mock example
type User struct {
	ID       int
	Name     string
	Email    string
	IsActive bool
}

// UserRepository interface for dependency injection
type UserRepository interface {
	GetByID(id int) (*User, error)
	Save(user *User) error
}

// UserService uses a repository
type UserService struct {
	repo UserRepository
}

// NewUserService creates a new service
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUser fetches a user by ID
func (s *UserService) GetUser(id int) (*User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.repo.GetByID(id)
}

// ActivateUser activates a user
func (s *UserService) ActivateUser(id int) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	user.IsActive = true
	return s.repo.Save(user)
}

func main() {
	fmt.Println("This package is meant to be tested!")
	fmt.Println("Run: go test -v ./...")
	fmt.Println()

	calc := NewCalculator()
	fmt.Println("5 + 3 =", calc.Add(5, 3))
	fmt.Println("5 - 3 =", calc.Subtract(5, 3))
	fmt.Println("5 * 3 =", calc.Multiply(5, 3))

	result, err := calc.Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	su := &StringUtils{}
	fmt.Println("Reverse 'hello':", su.Reverse("hello"))
	fmt.Println("IsPalindrome 'radar':", su.IsPalindrome("radar"))
}
