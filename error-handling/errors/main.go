package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
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
  errors.New("message")           // simple error
  fmt.Errorf("format", args...)   // formatted error
  custom types implementing error  // rich error info

BEST PRACTICES:
  - Always check errors, never ignore them
  - Add context when wrapping errors
  - Use errors.Is/As for comparison
  - Sentinel errors for specific conditions

*/

// CUSTOM ERROR TYPES

// ValidationError represents a validation failure
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Message)
}

// NotFoundError with context
type NotFoundError struct {
	Resource string
	ID       interface{}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %v not found", e.Resource, e.ID)
}

// SENTINEL ERRORS (known error values)
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidInput = errors.New("invalid input")
)

func main() {
	// 1. BASIC ERROR HANDLING

	fmt.Println("--- Basic Error Handling ---")

	// Division with error check
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
	username := "bob"
	err2 := fmt.Errorf("user %q not found", username)
	fmt.Println("fmt.Errorf:", err2)

	// 3. CUSTOM ERROR TYPES

	fmt.Println("\n--- Custom Error Types ---")

	err = validateUser("", -5)
	if err != nil {
		fmt.Println("Validation error:", err)

		// Type assertion to get details
		if ve, ok := err.(*ValidationError); ok {
			fmt.Printf("  Field: %s, Message: %s\n", ve.Field, ve.Message)
		}
	}

	// 4. ERROR WRAPPING (Go 1.13+)

	fmt.Println("\n--- Error Wrapping ---")

	err = readConfig("config.json")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Unwrap to get original error
	unwrapped := errors.Unwrap(err)
	if unwrapped != nil {
		fmt.Println("Unwrapped:", unwrapped)
	}

	// 5. ERRORS.IS AND ERRORS.AS

	fmt.Println("\n--- errors.Is and errors.As ---")

	// Check if error matches a specific error
	err = wrapError()
	fmt.Println("Error:", err)

	if errors.Is(err, ErrNotFound) {
		fmt.Println("  -> Is ErrNotFound: true")
	}

	// Extract specific error type
	err = getUserError()
	fmt.Println("\nError:", err)

	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		fmt.Printf("  -> As NotFoundError: Resource=%s, ID=%v\n", nfe.Resource, nfe.ID)
	}

	// 6. SENTINEL ERRORS

	fmt.Println("\n--- Sentinel Errors ---")

	_, err = fetchData(123)
	handleError(err)

	_, err = fetchData(0)
	handleError(err)

	_, err = fetchData(-1)
	handleError(err)

	// 7. ERROR HANDLING PATTERNS

	fmt.Println("\n--- Error Handling Patterns ---")

	// Pattern 1: Check and return
	if err := processStep1(); err != nil {
		fmt.Println("Step 1 failed:", err)
	}

	// Pattern 2: Check with context
	if err := processStep2(); err != nil {
		fmt.Println("Step 2 failed:", fmt.Errorf("in processStep2: %w", err))
	}

	// Pattern 3: Multiple returns
	data, count, err := fetchMultiple()
	if err != nil {
		fmt.Println("Fetch failed:", err)
	} else {
		fmt.Printf("Got %d items: %v\n", count, data)
	}

	// 8. PRACTICAL: FILE OPERATIONS

	fmt.Println("\n--- Practical: File Operations ---")

	content, err := readFileSafe("nonexistent.txt")
	if err != nil {
		fmt.Println("Read error:", err)

		// Check specific error type
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("  -> File does not exist")
		}
	} else {
		fmt.Println("Content:", content)
	}

	// 9. PRACTICAL: PARSING WITH ERRORS

	fmt.Println("\n--- Practical: Parsing ---")

	nums := []string{"42", "invalid", "100", "", "3.14"}
	for _, s := range nums {
		n, err := parseNumber(s)
		if err != nil {
			fmt.Printf("  Parse %q: ERROR - %v\n", s, err)
		} else {
			fmt.Printf("  Parse %q: %d\n", s, n)
		}
	}

	// 10. ERROR AGGREGATION

	fmt.Println("\n--- Error Aggregation ---")

	errs := validateForm(map[string]string{
		"name":  "",
		"email": "invalid",
		"age":   "-5",
	})

	if len(errs) > 0 {
		fmt.Println("Form validation failed:")
		for _, e := range errs {
			fmt.Println("  -", e)
		}
	}

	// 11. MUST PATTERN (PANIC ON ERROR)

	fmt.Println("\n--- Must Pattern ---")

	// Use when error should never happen
	config := mustLoadConfig()
	fmt.Println("Config loaded:", config)

	// 12. DEFER FOR CLEANUP ON ERROR

	fmt.Println("\n--- Cleanup on Error ---")

	err = processWithCleanup()
	if err != nil {
		fmt.Println("Process failed:", err)
	}
}

// HELPER FUNCTIONS

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
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
	// Simulate file read error
	originalErr := os.ErrNotExist

	// Wrap with context using %w
	return fmt.Errorf("failed to read config %s: %w", filename, originalErr)
}

func wrapError() error {
	// Multiple layers of wrapping
	return fmt.Errorf("layer 2: %w",
		fmt.Errorf("layer 1: %w", ErrNotFound))
}

func getUserError() error {
	return &NotFoundError{Resource: "User", ID: 42}
}

func fetchData(id int) (string, error) {
	switch {
	case id == 0:
		return "", ErrInvalidInput
	case id < 0:
		return "", ErrUnauthorized
	case id == 123:
		return "", ErrNotFound
	default:
		return fmt.Sprintf("data-%d", id), nil
	}
}

func handleError(err error) {
	if err == nil {
		fmt.Println("Success")
		return
	}

	switch {
	case errors.Is(err, ErrNotFound):
		fmt.Println("Handle: Resource not found")
	case errors.Is(err, ErrUnauthorized):
		fmt.Println("Handle: Access denied")
	case errors.Is(err, ErrInvalidInput):
		fmt.Println("Handle: Bad input")
	default:
		fmt.Println("Handle: Unknown error:", err)
	}
}

func processStep1() error {
	return errors.New("step 1 error")
}

func processStep2() error {
	return errors.New("underlying error")
}

func fetchMultiple() ([]int, int, error) {
	data := []int{1, 2, 3}
	return data, len(data), nil
}

func readFileSafe(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("reading %s: %w", filename, err)
	}
	return string(content), nil
}

func parseNumber(s string) (int, error) {
	if s == "" {
		return 0, errors.New("empty string")
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid number %q: %w", s, err)
	}

	return n, nil
}

func validateForm(fields map[string]string) []error {
	var errs []error

	if fields["name"] == "" {
		errs = append(errs, errors.New("name is required"))
	}

	if fields["email"] != "" && !isValidEmail(fields["email"]) {
		errs = append(errs, errors.New("invalid email format"))
	}

	if age, err := strconv.Atoi(fields["age"]); err == nil && age < 0 {
		errs = append(errs, errors.New("age cannot be negative"))
	}

	return errs
}

func isValidEmail(email string) bool {
	// Simplified validation
	for _, c := range email {
		if c == '@' {
			return true
		}
	}
	return false
}

func mustLoadConfig() map[string]string {
	// In real code, this would load from file/env
	config := map[string]string{
		"app":  "myapp",
		"port": "8080",
	}
	return config
}

func processWithCleanup() (err error) {
	fmt.Println("  Acquiring resources...")
	resource := "some resource"

	defer func() {
		fmt.Println("  Cleaning up", resource)
		if err != nil {
			fmt.Println("  (cleanup due to error)")
		}
	}()

	fmt.Println("  Processing...")

	// Simulate error
	return errors.New("processing failed")
}

// REAL-WORLD PATTERN: Copy with error handling
func copyFile(src, dst string) (err error) {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("opening source: %w", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("creating destination: %w", err)
	}
	defer func() {
		closeErr := destFile.Close()
		if err == nil {
			err = closeErr
		}
	}()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("copying data: %w", err)
	}

	return nil
}
