package main

import (
	"errors"
	"fmt"
)

/*
  ERROR HANDLING IN GO

KEY POINTS:
  - Go uses explicit error returns instead of exceptions
  - error is a built-in interface: Error() string
  - Convention: error is the last return value
  - Check errors immediately after function calls
  - nil error means success

ERROR INTERFACE:
  type error interface {
      Error() string
  }

CREATING ERRORS:
  errors.New("message")           // Simple error
  fmt.Errorf("format", args...)   // Formatted error
  fmt.Errorf("wrap: %w", err)     // Wrap error (Go 1.13+)
  Custom types implementing error // Rich error info

CHECKING ERRORS:
  errors.Is(err, target)          // Check if err matches target
  errors.As(err, &target)         // Extract error of specific type
  errors.Unwrap(err)              // Get wrapped error

BEST PRACTICES:
  - Always check errors
  - Add context when wrapping
  - Use sentinel errors for known conditions
  - Use custom types for rich error info

*/

// SENTINEL ERRORS (predefined error values)
var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrUnauthorized = errors.New("unauthorized")
)

// CUSTOM ERROR TYPE with rich information
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

// DivisionError for division by zero
type DivisionError struct {
	Dividend int
	Divisor  int
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}

func main() {
	// 1. BASIC ERROR HANDLING

	fmt.Println("--- Basic Error Handling ---")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 2. CREATING ERRORS

	fmt.Println("\n--- Creating Errors ---")

	// errors.New
	err1 := errors.New("something went wrong")
	fmt.Println("errors.New:", err1)

	// fmt.Errorf (formatted)
	name := "config.json"
	err2 := fmt.Errorf("failed to load %s", name)
	fmt.Println("fmt.Errorf:", err2)

	// 3. CUSTOM ERROR TYPES

	fmt.Println("\n--- Custom Error Types ---")

	err = validateUser("", -5)
	if err != nil {
		fmt.Println("Validation failed:", err)

		// Type assertion to get details
		if ve, ok := err.(*ValidationError); ok {
			fmt.Printf("  Field: %s\n", ve.Field)
			fmt.Printf("  Message: %s\n", ve.Message)
		}
	}

	// 4. ERROR WRAPPING (Go 1.13+)

	fmt.Println("\n--- Error Wrapping ---")

	err = readConfig("missing.json")
	if err != nil {
		fmt.Println("Error:", err)

		// Unwrap to get original
		unwrapped := errors.Unwrap(err)
		if unwrapped != nil {
			fmt.Println("Unwrapped:", unwrapped)
		}
	}

	// 5. ERRORS.IS - Check Error Identity

	fmt.Println("\n--- errors.Is ---")

	err = fetchUser(999)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("User not found (caught with errors.Is)")
	}

	// Works with wrapped errors too
	wrappedErr := fmt.Errorf("database error: %w", ErrNotFound)
	if errors.Is(wrappedErr, ErrNotFound) {
		fmt.Println("Found ErrNotFound in wrapped error")
	}

	// 6. ERRORS.AS - Extract Error Type

	fmt.Println("\n--- errors.As ---")

	err = processData(-10)
	var divErr *DivisionError
	if errors.As(err, &divErr) {
		fmt.Println("Caught DivisionError:")
		fmt.Printf("  Dividend: %d\n", divErr.Dividend)
		fmt.Printf("  Divisor: %d\n", divErr.Divisor)
	}

	// 7. SENTINEL ERRORS PATTERN

	fmt.Println("\n--- Sentinel Errors ---")

	testCases := []int{1, 0, -1, 42}
	for _, id := range testCases {
		_, err := getItem(id)
		switch {
		case err == nil:
			fmt.Printf("ID %d: Success\n", id)
		case errors.Is(err, ErrNotFound):
			fmt.Printf("ID %d: Not found\n", id)
		case errors.Is(err, ErrInvalidInput):
			fmt.Printf("ID %d: Invalid input\n", id)
		case errors.Is(err, ErrUnauthorized):
			fmt.Printf("ID %d: Unauthorized\n", id)
		default:
			fmt.Printf("ID %d: Unknown error: %v\n", id, err)
		}
	}

	// 8. PANIC AND RECOVER

	fmt.Println("\n--- Panic and Recover ---")

	// Safe division with recover
	safeResult := safeDivide(10, 0)
	fmt.Println("Safe divide 10/0:", safeResult)

	safeResult = safeDivide(10, 2)
	fmt.Println("Safe divide 10/2:", safeResult)

	// 9. DEFER WITH ERROR HANDLING

	fmt.Println("\n--- Defer with Errors ---")

	err = processWithCleanup()
	if err != nil {
		fmt.Println("Process error:", err)
	}

	// 10. MULTIPLE ERROR AGGREGATION

	fmt.Println("\n--- Multiple Errors ---")

	errs := validateForm(map[string]string{
		"name":  "",
		"email": "invalid",
		"age":   "-5",
	})

	if len(errs) > 0 {
		fmt.Println("Validation errors:")
		for _, e := range errs {
			fmt.Println("  -", e)
		}
	}

	// 11. ERROR HANDLING PATTERNS

	fmt.Println("\n--- Error Handling Patterns ---")

	// Pattern 1: Early return
	if err := doStep1(); err != nil {
		fmt.Println("Step 1 failed:", err)
	}

	// Pattern 2: Add context when propagating
	if err := doStep2(); err != nil {
		wrapped := fmt.Errorf("in step 2: %w", err)
		fmt.Println("Step 2 failed:", wrapped)
	}

	// Pattern 3: Handle specific errors differently
	if err := doStep3(); err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("Step 3: Resource not found, using default")
		} else {
			fmt.Println("Step 3 failed:", err)
		}
	}
}

// HELPER FUNCTIONS

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

func validateUser(name string, age int) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "cannot be empty"}
	}
	if age < 0 {
		return &ValidationError{Field: "age", Message: "cannot be negative"}
	}
	return nil
}

func readConfig(filename string) error {
	// Simulate file not found
	originalErr := errors.New("file not found")
	// Wrap with context using %w
	return fmt.Errorf("failed to read %s: %w", filename, originalErr)
}

func fetchUser(id int) error {
	if id == 999 {
		return ErrNotFound
	}
	return nil
}

func processData(value int) error {
	if value < 0 {
		return &DivisionError{Dividend: value, Divisor: 0}
	}
	return nil
}

func getItem(id int) (string, error) {
	switch {
	case id <= 0:
		return "", ErrInvalidInput
	case id == 42:
		return "", ErrUnauthorized
	case id > 100:
		return "", ErrNotFound
	default:
		return fmt.Sprintf("item-%d", id), nil
	}
}

func safeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
			result = 0
		}
	}()

	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func processWithCleanup() (err error) {
	fmt.Println("  Acquiring resources...")

	defer func() {
		fmt.Println("  Cleaning up resources...")
	}()

	fmt.Println("  Processing...")
	return errors.New("simulated processing error")
}

func validateForm(fields map[string]string) []error {
	var errs []error

	if fields["name"] == "" {
		errs = append(errs, errors.New("name is required"))
	}
	if fields["email"] == "invalid" {
		errs = append(errs, errors.New("email format is invalid"))
	}
	if fields["age"] == "-5" {
		errs = append(errs, errors.New("age cannot be negative"))
	}

	return errs
}

func doStep1() error { return errors.New("step 1 error") }
func doStep2() error { return errors.New("underlying error") }
func doStep3() error { return ErrNotFound }
