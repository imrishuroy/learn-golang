package main

import "fmt"

/*
  STRUCTS IN GO

KEY POINTS:
  - Structs group related data into a single type
  - Structs are VALUE types (assignment creates a copy)
  - Use pointers to modify or avoid copying large structs
  - Fields accessed with dot notation: s.FieldName
  - Go auto-dereferences pointers to structs

SYNTAX:
  type Name struct {
      Field1 Type1
      Field2 Type2
  }

CREATING STRUCTS:
  s := Name{val1, val2}           // positional (fragile)
  s := Name{Field1: val1}         // named fields (preferred)
  s := Name{}                     // zero values
  p := &Name{Field1: val1}        // pointer to struct

*/

// TYPE DEFINITIONS

// Person represents a basic person
type Person struct {
	Name string
	Age  int
}

// Address for nested struct example
type Address struct {
	Street  string
	City    string
	Country string
}

// Employee demonstrates nested structs
type Employee struct {
	Person  // embedded struct (anonymous field)
	Title   string
	Address Address // named field
}

// Point for method examples
type Point struct {
	X, Y float64
}

func main() {
	// 1. CREATING STRUCTS

	fmt.Println("--- Creating Structs ---")

	// Method 1: Positional (order matters - not recommended)
	p1 := Person{"Alice", 30}
	fmt.Println("Positional:", p1)

	// Method 2: Named fields (recommended)
	p2 := Person{Name: "Bob", Age: 25}
	fmt.Println("Named fields:", p2)

	// Method 3: Partial initialization (others get zero value)
	p3 := Person{Name: "Charlie"}
	fmt.Println("Partial:", p3) // Age = 0

	// Method 4: Zero value
	var p4 Person
	fmt.Println("Zero value:", p4) // Name = "", Age = 0

	// 2. ACCESSING AND MODIFYING FIELDS

	fmt.Println("\n--- Accessing Fields ---")

	person := Person{Name: "Diana", Age: 28}
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)

	// Modify fields
	person.Age = 29
	person.Name = "Diana Smith"
	fmt.Println("Modified:", person)

	// 3. STRUCTS ARE VALUE TYPES (COPYING)

	fmt.Println("\n--- Value Types (Copying) ---")

	original := Person{Name: "Eve", Age: 35}
	copied := original // full copy

	copied.Name = "Eve Modified"
	copied.Age = 36

	fmt.Println("Original:", original) // unchanged
	fmt.Println("Copied:", copied)     // modified

	// 4. POINTERS TO STRUCTS

	fmt.Println("\n--- Pointers to Structs ---")

	// Create pointer with &
	frank := Person{Name: "Frank", Age: 40}
	ptr := &frank

	// Access fields (Go auto-dereferences)
	fmt.Println("ptr.Name:", ptr.Name) // same as (*ptr).Name
	fmt.Println("ptr.Age:", ptr.Age)

	// Modify through pointer
	ptr.Age = 41
	fmt.Println("After ptr.Age = 41:", frank) // frank is modified!

	// Create with &Type{} (common pattern)
	grace := &Person{Name: "Grace", Age: 33}
	fmt.Println("Created with &:", *grace)

	// 5. PASSING STRUCTS TO FUNCTIONS

	fmt.Println("\n--- Passing to Functions ---")

	person1 := Person{Name: "Henry", Age: 50}

	// Pass by value (copy)
	birthdayValue(person1)
	fmt.Println("After birthdayValue:", person1) // unchanged

	// Pass by pointer
	birthdayPointer(&person1)
	fmt.Println("After birthdayPointer:", person1) // age + 1

	// 6. ANONYMOUS STRUCTS

	fmt.Println("\n--- Anonymous Structs ---")

	// Useful for one-off data structures
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Server: %s:%d\n", config.Host, config.Port)

	// Common in tests and JSON handling
	response := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "ok",
		Message: "Success",
	}
	fmt.Println("Response:", response)

	// 7. NESTED STRUCTS

	fmt.Println("\n--- Nested Structs ---")

	emp := Employee{
		Person: Person{Name: "Ivan", Age: 35},
		Title:  "Engineer",
		Address: Address{
			Street:  "123 Main St",
			City:    "NYC",
			Country: "USA",
		},
	}

	fmt.Println("Employee:", emp)
	fmt.Println("Name:", emp.Person.Name)
	fmt.Println("City:", emp.Address.City)

	// 8. EMBEDDED STRUCTS (COMPOSITION)

	fmt.Println("\n--- Embedded Structs ---")

	// With embedded Person, fields are "promoted"
	fmt.Println("emp.Name:", emp.Name) // same as emp.Person.Name
	fmt.Println("emp.Age:", emp.Age)   // same as emp.Person.Age

	// 9. STRUCT COMPARISON

	fmt.Println("\n--- Struct Comparison ---")

	a := Person{Name: "Jack", Age: 30}
	b := Person{Name: "Jack", Age: 30}
	c := Person{Name: "Jack", Age: 31}

	fmt.Println("a == b:", a == b) // true
	fmt.Println("a == c:", a == c) // false

	// Note: Structs with slices/maps/functions can't use ==

	// 10. METHODS ON STRUCTS

	fmt.Println("\n--- Methods on Structs ---")

	pt := Point{X: 3, Y: 4}

	// Value receiver method
	fmt.Println("Distance from origin:", pt.Distance())

	// Pointer receiver method
	fmt.Println("Before scale:", pt)
	pt.Scale(2)
	fmt.Println("After scale(2):", pt)

	// 11. CONSTRUCTOR PATTERN

	fmt.Println("\n--- Constructor Pattern ---")

	// Go doesn't have constructors, use factory functions
	kate := NewPerson("Kate", 28)
	fmt.Println("From NewPerson:", kate)

	// With validation
	validPerson, err := NewPersonValidated("Leo", 25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Valid person:", validPerson)
	}

	// 12. STRUCT TAGS

	fmt.Println("\n--- Struct Tags ---")

	type User struct {
		ID       int    `json:"id" db:"user_id"`
		Username string `json:"username" db:"user_name"`
		Email    string `json:"email,omitempty"` // omit if empty
		Password string `json:"-"`               // never include in JSON
	}

	user := User{ID: 1, Username: "mike", Email: "", Password: "secret"}
	fmt.Printf("User: %+v\n", user)
	fmt.Println("Tags are used by encoding/json, database/sql, etc.")
}

// FUNCTIONS

// Pass by value - receives copy
func birthdayValue(p Person) {
	p.Age++
}

// Pass by pointer - modifies original
func birthdayPointer(p *Person) {
	p.Age++
}

// Value receiver - doesn't modify
func (p Point) Distance() float64 {
	return p.X*p.X + p.Y*p.Y
}

// Pointer receiver - can modify
func (p *Point) Scale(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Factory function (constructor pattern)
func NewPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

// Factory with validation
func NewPersonValidated(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if age < 0 {
		return nil, fmt.Errorf("age cannot be negative")
	}
	return &Person{Name: name, Age: age}, nil
}
