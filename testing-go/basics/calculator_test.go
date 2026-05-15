package main

import (
	"errors"
	"testing"
)

// 1. BASIC TEST

func TestAdd(t *testing.T) {
	calc := NewCalculator()

	result := calc.Add(2, 3)
	expected := 5.0

	if result != expected {
		t.Errorf("Add(2, 3) = %f; want %f", result, expected)
	}
}

// 2. TABLE-DRIVEN TESTS (PREFERRED PATTERN)

func TestCalculator_Add_TableDriven(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", -2, 3, 1},
		{"zeros", 0, 0, 0},
		{"decimals", 1.5, 2.5, 4.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Divide_TableDriven(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		expectError bool
	}{
		{"normal division", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
		{"negative numbers", -10, 2, -5, false},
		{"decimals", 7.5, 2.5, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Divide(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Errorf("Divide(%f, %f) expected error, got nil", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%f, %f) unexpected error: %v", tt.a, tt.b, err)
				}
				if result != tt.expected {
					t.Errorf("Divide(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}

// 3. SUBTESTS

func TestCalculator(t *testing.T) {
	calc := NewCalculator()

	t.Run("Add", func(t *testing.T) {
		if calc.Add(1, 2) != 3 {
			t.Error("Add failed")
		}
	})

	t.Run("Subtract", func(t *testing.T) {
		if calc.Subtract(5, 3) != 2 {
			t.Error("Subtract failed")
		}
	})

	t.Run("Multiply", func(t *testing.T) {
		if calc.Multiply(4, 3) != 12 {
			t.Error("Multiply failed")
		}
	})
}

// 4. STRING UTILS TESTS

func TestReverse(t *testing.T) {
	su := &StringUtils{}

	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"Go语言", "言语oG"},
	}

	for _, tt := range tests {
		result := su.Reverse(tt.input)
		if result != tt.expected {
			t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	su := &StringUtils{}

	tests := []struct {
		input    string
		expected bool
	}{
		{"radar", true},
		{"hello", false},
		{"", true},
		{"a", true},
		{"racecar", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := su.IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// 5. MOCK IMPLEMENTATION FOR TESTING

type MockUserRepository struct {
	users     map[int]*User
	saveError error
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[int]*User),
	}
}

func (m *MockUserRepository) GetByID(id int) (*User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (m *MockUserRepository) Save(user *User) error {
	if m.saveError != nil {
		return m.saveError
	}
	m.users[user.ID] = user
	return nil
}

func TestUserService_GetUser(t *testing.T) {
	repo := NewMockUserRepository()
	repo.users[1] = &User{ID: 1, Name: "Alice", Email: "alice@example.com"}

	service := NewUserService(repo)

	t.Run("existing user", func(t *testing.T) {
		user, err := service.GetUser(1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if user.Name != "Alice" {
			t.Errorf("expected Alice, got %s", user.Name)
		}
	})

	t.Run("non-existing user", func(t *testing.T) {
		_, err := service.GetUser(999)
		if err == nil {
			t.Error("expected error for non-existing user")
		}
	})

	t.Run("invalid ID", func(t *testing.T) {
		_, err := service.GetUser(-1)
		if err == nil {
			t.Error("expected error for invalid ID")
		}
	})
}

func TestUserService_ActivateUser(t *testing.T) {
	repo := NewMockUserRepository()
	repo.users[1] = &User{ID: 1, Name: "Bob", IsActive: false}

	service := NewUserService(repo)

	err := service.ActivateUser(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	user, _ := repo.GetByID(1)
	if !user.IsActive {
		t.Error("user should be active after ActivateUser")
	}
}

// 6. HELPER FUNCTIONS

func assertEqual(t *testing.T, got, want float64) {
	t.Helper() // marks this as helper, so line numbers point to caller
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestWithHelper(t *testing.T) {
	calc := NewCalculator()
	assertEqual(t, calc.Add(1, 2), 3)
	assertEqual(t, calc.Multiply(2, 3), 6)
}

// 7. SKIP TESTS

func TestSkipExample(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	// Long running test...
}

// 8. BENCHMARKS (run with: go test -bench=.)

func BenchmarkAdd(b *testing.B) {
	calc := NewCalculator()
	for i := 0; i < b.N; i++ {
		calc.Add(1, 2)
	}
}

func BenchmarkReverse(b *testing.B) {
	su := &StringUtils{}
	for i := 0; i < b.N; i++ {
		su.Reverse("hello world")
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	su := &StringUtils{}
	for i := 0; i < b.N; i++ {
		su.IsPalindrome("racecar")
	}
}

// 9. EXAMPLE TESTS (documentation + verification)

func ExampleCalculator_Add() {
	calc := NewCalculator()
	result := calc.Add(2, 3)
	println(result)
	// Output: +5.000000e+000
}

func ExampleStringUtils_Reverse() {
	su := &StringUtils{}
	result := su.Reverse("hello")
	println(result)
	// Output: olleh
}

// 10. TEST SETUP AND TEARDOWN

func TestWithSetup(t *testing.T) {
	// Setup
	calc := NewCalculator()

	// Teardown (runs even if test fails)
	t.Cleanup(func() {
		// cleanup code here
	})

	// Test
	if calc.Add(1, 1) != 2 {
		t.Error("1 + 1 should equal 2")
	}
}
