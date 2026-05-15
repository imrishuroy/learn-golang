# Learn Go

A comprehensive, tutorial-style repository for learning Go through hands-on examples. Each topic includes detailed comments explaining concepts, practical examples, and common patterns.

## Project Structure

```
learn-golang/
│
├── variables/                    # Data types, constants, type conversion
├── flow-control/                 # Loops, if/else, switch, defer
├── arrays/                       # Fixed-size collections, iteration
├── slices/                       # Dynamic arrays, capacity, append
├── maps/                         # Key-value pairs, sets, nested maps
├── pointers/                     # Memory addresses, dereferencing
├── structs/                      # Custom types, composition, tags
├── collections/                  # Arrays, slices, maps (combined)
│
├── functions/                    # Higher-order, closures, variadic
├── functions-patterns/           # Function patterns
│   ├── basics/                   # Parameters, returns, variadic
│   ├── methods/                  # Receivers, value vs pointer
│   ├── anonymous_functions/      # Closures, callbacks, decorators
│   └── defer-panic-recover/      # Resource cleanup, error recovery
│
├── interfaces/                   # Polymorphism, type assertions
├── generics/                     # Type parameters, constraints
├── oop/                          # OOP patterns
│   ├── interfaces/               # Polymorphism, type assertions
│   └── generics/                 # Type parameters, constraints
│
├── concurrency/                  # Concurrent programming
│   ├── goroutines-basic/         # Lightweight threads, runtime
│   ├── waitgroup/                # Synchronization, worker pools
│   ├── channels/                 # Communication, select, pipelines
│   ├── mutex/                    # Locks, RWMutex, sync primitives
│   └── context/                  # Cancellation, timeouts, values
├── concurrency_and_channels/     # Goroutines, channels, mutex
│
├── error-handling/errors/        # Creating, wrapping, custom errors
├── error_handling/               # Error patterns, panic/recover
│
├── testing-go/basics/            # Unit tests, table-driven, benchmarks
├── data-type-loops-and-flow-control/  # Types, loops, conditionals
├── files/                        # File I/O, CSV, JSON, directories
│
├── projstruct/                   # Multi-package project (standalone)
├── ocppcli/                      # CLI application (standalone)
├── builder-pattern/              # Builder design pattern
├── logger/                       # Logging with logrus
└── viperconfig/                  # Configuration with viper
```

## Running Examples

```bash
# Fundamentals
go run ./variables
go run ./flow-control
go run ./arrays
go run ./slices
go run ./maps
go run ./pointers
go run ./structs
go run ./collections

# Functions
go run ./functions
go run ./functions-patterns/basics
go run ./functions-patterns/methods
go run ./functions-patterns/anonymous_functions
go run ./functions-patterns/defer-panic-recover

# OOP
go run ./interfaces
go run ./generics
go run ./oop/interfaces
go run ./oop/generics

# Concurrency
go run ./concurrency/goroutines-basic
go run ./concurrency/waitgroup
go run ./concurrency/channels
go run ./concurrency/mutex
go run ./concurrency/context
go run ./concurrency_and_channels

# Error Handling
go run ./error-handling/errors
go run ./error_handling

# Other
go run ./data-type-loops-and-flow-control
go run ./files

# Testing
cd testing-go/basics && go test -v ./...

# Race detector
go run -race ./concurrency/channels
```

## Topics Covered

### Fundamentals
| Folder | Topics |
|--------|--------|
| `variables` | Data types, type inference, constants, iota, zero values |
| `flow-control` | for loops, range, if/else, switch, defer |
| `arrays` | Declaration, iteration, multidimensional, value semantics |
| `slices` | Creating, slicing, append, copy, capacity management |
| `maps` | CRUD operations, iteration, nested maps, sets |
| `pointers` | Address-of (&), dereference (*), nil safety, methods |
| `structs` | Fields, embedding, methods, constructors, tags |
| `collections` | Arrays, slices, maps combined tutorial |

### Functions
| Folder | Topics |
|--------|--------|
| `functions` | Higher-order functions, closures, variadic, filter |
| `functions-patterns/basics` | Parameters, multiple returns, named returns |
| `functions-patterns/methods` | Value/pointer receivers, chaining |
| `functions-patterns/anonymous_functions` | Closures, IIFE, callbacks, decorators |
| `functions-patterns/defer-panic-recover` | Resource cleanup, panic handling |

### OOP
| Folder | Topics |
|--------|--------|
| `interfaces` | Polymorphism, type switch, dependency injection |
| `generics` | Constraints, Stack, Pair, Cache, Map/Filter/Reduce |
| `oop/interfaces` | Implicit implementation, type assertions, Stringer |
| `oop/generics` | Type parameters, constraints, generic types |

### Concurrency
| Folder | Topics |
|--------|--------|
| `concurrency/goroutines-basic` | Creating goroutines, closure gotcha |
| `concurrency/waitgroup` | sync.WaitGroup, worker pools |
| `concurrency/channels` | Buffered/unbuffered, select, pipelines |
| `concurrency/mutex` | sync.Mutex, RWMutex, sync.Once, sync.Map |
| `concurrency/context` | WithCancel, WithTimeout, graceful shutdown |
| `concurrency_and_channels` | Combined concurrency patterns |

### Error Handling
| Folder | Topics |
|--------|--------|
| `error-handling/errors` | Creating, wrapping, errors.Is/As, custom types |
| `error_handling` | Error patterns, panic/recover |

### Testing & Other
| Folder | Topics |
|--------|--------|
| `testing-go/basics` | Unit tests, table-driven, mocks, benchmarks |
| `data-type-loops-and-flow-control` | Types, constants, loops, switch |
| `files` | File I/O, CSV, JSON, directories |

## Learning Path

**Recommended order:**

1. `variables` - Data types and constants
2. `flow-control` - Loops and conditionals
3. `arrays` - Fixed-size collections
4. `slices` - Dynamic arrays
5. `maps` - Key-value storage
6. `pointers` - Memory and references
7. `structs` - Custom types
8. `functions` - Function basics
9. `functions-patterns/methods` - Methods on types
10. `interfaces` - Polymorphism
11. `error-handling/errors` - Error patterns
12. `concurrency/goroutines-basic` - Goroutines
13. `concurrency/channels` - Channels
14. `concurrency/mutex` - Synchronization
15. `concurrency/context` - Cancellation
16. `testing-go/basics` - Writing tests
17. `generics` - Generic programming
18. `files` - File operations

## Standalone Modules

Run from inside their directory:

```bash
cd projstruct && go run .
cd ocppcli && go run .
```

## Resources

- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Proverbs](https://go-proverbs.github.io/)
- [Go Playground](https://go.dev/play/)
