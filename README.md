# Test-Driven Development in Golang

A practical demonstration of **Test-Driven Development (TDD)** in Go, showcasing how to build an HTTP server by writing tests first. This repository exemplifies the Red-Green-Refactor cycle and TDD best practices using Go's standard testing library.

## 🎯 What is Test-Driven Development?

TDD is a software development approach where you write tests **before** writing the implementation code. The process follows three simple steps:

1. **🔴 Red**: Write a failing test that defines the desired behavior
2. **🟢 Green**: Write the minimum code to make the test pass
3. **♻️ Refactor**: Improve the code while keeping tests passing

This project demonstrates TDD principles through:
- HTTP handler tests written before implementation
- User management logic driven by test cases
- Comprehensive edge case coverage
- Clear separation of test concerns

## 📚 TDD Concepts Demonstrated

- ✅ **Red-Green-Refactor Cycle**: See how each feature starts with a failing test
- ✅ **Test-Driven API Design**: API endpoints designed through test specifications
- ✅ **Edge Case Testing**: Error handling and validation driven by test cases
- ✅ **Go Testing Best Practices**: Using `httptest`, table-driven tests, and proper assertions
- ✅ **Test Organization**: Separate test files demonstrating clean test structure

## Table of Contents

- [What is TDD?](#-what-is-test-driven-development)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [TDD Workflow Examples](#tdd-workflow-examples)
- [Project Structure](#project-structure)
- [Running Tests](#running-tests)
- [TDD Best Practices](#tdd-best-practices-demonstrated)
- [API Documentation](#api-documentation)
- [Development Guide](#development-guide)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

- Go 1.25 or higher
- Basic understanding of Go testing (`go test`)
- Familiarity with HTTP handlers in Go

## Installation

### Clone the Repository

```bash
git clone https://github.com/kunalkumar-1/go-http.git
cd go-http
```

### Install Dependencies

This project uses only Go's standard library, so no external dependencies need to be installed:

```bash
go mod download
```

## TDD Workflow Examples

### Example 1: Root Handler (GET /)

**Step 1: Red - Write the failing test**
```go
func TestHandleRoot(t *testing.T) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodGet, "/health", nil)
    handleRoot(w, r)

    desiredCode := http.StatusOK
    if w.Code != desiredCode {
        t.Errorf("bad response code: expected %d, got %d", desiredCode, w.Code)
    }

    expectedMessage := []byte("Welcome to our HomePage!\n")
    if !bytes.Equal(w.Body.Bytes(), expectedMessage) {
        t.Errorf("bad response body: expected %s, got %s", 
            string(expectedMessage), string(w.Body.Bytes()))
    }
}
```
*Run: `go test` - Test fails because `handleRoot` doesn't exist yet*

**Step 2: Green - Implement minimum code**
```go
func handleRoot(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to our HomePage!\n"))
}
```
*Run: `go test` - Test passes! ✅*

**Step 3: Refactor - Improve while keeping tests green**
```go
func handleRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Requested path:", r.URL.Path)
    _, err := w.Write([]byte("Welcome to our HomePage!\n"))
    if err != nil {
        slog.Error("Error serving the health_handler err: " + err.Error())
        return
    }
}
```
*Run: `go test` - Still passes! ✅*

### Example 2: User Management (TDD for Business Logic)

**Step 1: Red - Test adding a user**
```go
func TestAddUser(t *testing.T) {
    testManager := NewManager()
    
    err := testManager.AddUser("jhon", "smith", "foo@bar.com")
    if err != nil {
        t.Fatalf("failed to add user: %v", err)
    }
    
    if len(testManager.users) != 1 {
        t.Fatalf("expected 1 user, got %v", len(testManager.users))
    }
}
```

**Step 2: Green - Implement AddUser**
```go
func (m *Manager) AddUser(firstName string, lastName string, email string) error {
    // Minimum implementation to pass
    newUser := User{
        FirstName: firstName,
        LastName:  lastName,
        Email:     mail.Address{Address: email},
    }
    m.users = append(m.users, newUser)
    return nil
}
```

**Step 3: Refactor - Add validation tests, then implement**
```go
// New test for validation
func TestAddUserInvalidEmail(t *testing.T) {
    testManager := NewManager()
    err := testManager.AddUser("jhon", "smith", "foobar") // Invalid email
    if err == nil {
        t.Errorf("expected error for invalid email")
    }
}
```

*Now implement validation in AddUser to make this test pass*

## Project Structure

```
go-http/
├── cmd/
│   └── server/
│       ├── main.go          # Implementation (written AFTER tests)
│       └── main_test.go     # HTTP handler tests (drives implementation)
├── internal/
│   └── users/
│       ├── users.go         # User logic (written AFTER tests)
│       └── users_test.go    # User tests (drives implementation)
├── go.mod
└── README.md
```

### Key TDD Structure Points

- **`*_test.go` files**: Tests written first, define the API contract
- **Implementation files**: Code written to satisfy tests
- **Test organization**: Each package has corresponding tests alongside
- **Test-driven design**: API signatures emerge from test needs

## API Documentation

> **Note**: All API endpoints were designed through TDD. The test files (`*_test.go`) serve as the API specification and contract.

### How Tests Define the API

Each endpoint's behavior is defined by its tests. For example, the `/json` endpoint's behavior is specified by:

- `TestHandleJSON` - Defines successful JSON handling
- `TestHandleJSONEmptyBody` - Defines error handling for empty body
- `TestHandleJSONEmptyNameFeild` - Defines validation for Name field

Looking at the tests tells you exactly what the API does!

### Endpoints

#### `GET /`
Returns a welcome message.

**Test**: `TestHandleRoot` in `cmd/server/main_test.go`

**Example:**
```bash
curl http://localhost:4000/
```

**Response:**
```
Welcome to our HomePage!
```

---

#### `GET /goodbye`
Returns a goodbye message.

**Test**: `TestHandleGoodbye` in `cmd/server/main_test.go`

**Example:**
```bash
curl http://localhost:4000/goodbye
```

**Response:**
```
Goodbye world is served at goodbye
```

---

#### `GET /hello/`
Returns a personalized hello message using query parameters.

**Tests**: 
- `TestHandleHelloParameterized` - With user parameter
- `TestHandleHelloNoParameterized` - Without user parameter (defaults)
- `TestHandleHelloWrongParameterized` - Invalid parameter

**Query Parameters:**
- `user` (optional): The username to greet. Defaults to "User" if not provided.

**Example:**
```bash
curl "http://localhost:4000/hello/?user=John"
```

**Response:**
```
Hello John!
```

---

#### `GET /responses/{user}/hello/`
Returns a personalized hello message using path variables.

**Test**: `TestHandleUserResponsesHello` in `cmd/server/main_test.go`

**Path Parameters:**
- `user`: The username to greet (required)

**Example:**
```bash
curl http://localhost:4000/responses/Alice/hello/
```

**Response:**
```
Hello Alice!
```

---

#### `GET /user/hello`
Returns a personalized hello message using HTTP headers.

**Tests**:
- `TestHandleHelloHeader` - Valid header
- `TestHandleHelloNoHeader` - Missing header (error case)

**Headers:**
- `user`: The username to greet (required)

**Example:**
```bash
curl -H "user: Bob" http://localhost:4000/user/hello
```

**Response:**
```
Hello Bob!
```

**Error Response (400 Bad Request):**
If the `user` header is missing (as tested in `TestHandleHelloNoHeader`):
```
invalid username provided
```

---

#### `POST /json`
Accepts a JSON payload and returns a personalized hello message.

**Tests**:
- `TestHandleJSON` - Valid JSON payload
- `TestHandleJSONEmptyBody` - Empty body error
- `TestHandleJSONEmptyNameFeild` - Missing Name field error

**Request Body:**
```json
{
  "Name": "Charlie"
}
```

**Example:**
```bash
curl -X POST http://localhost:4000/json \
  -H "Content-Type: application/json" \
  -d '{"Name":"Charlie"}'
```

**Response:**
```
Hello Charlie!
```

**Error Responses:**

- **400 Bad Request** - Empty request body (from `TestHandleJSONEmptyBody`):
  ```
  empty request body
  ```

- **400 Bad Request** - Invalid JSON or missing Name field (from `TestHandleJSONEmptyNameFeild`):
  ```
  invalid request body!
  ```

### Running the Server

To test the API manually:

```bash
# Start the server
go run ./cmd/server/main.go

# The server runs on port 4000
# Test endpoints using curl or your API client
```

## Project Structure

```
go-http/
├── cmd/
│   └── server/
│       ├── main.go          # Main application entry point
│       └── main_test.go     # HTTP handler tests
├── internal/
│   └── users/
│       ├── users.go         # User management logic
│       └── users_test.go    # User management tests
├── go.mod                   # Go module definition
└── README.md               # This file
```

### Package Overview

- **`cmd/server`**: HTTP server implementation with route handlers
- **`internal/users`**: User management package with CRUD operations and email validation

## Running Tests

### Run All Tests

```bash
go test ./...
```

### Run Tests with Coverage

```bash
go test -cover ./...
```

This will show you how much of your code is covered by tests. In TDD, high coverage is a natural byproduct.

### Run Tests in Watch Mode (Recommended for TDD)

For a true TDD experience, use a file watcher to run tests automatically:

```bash
# Install air (hot reload tool)
go install github.com/cosmtrek/air@latest

# Or use entr (Unix/Mac)
brew install entr
ls *.go **/*.go | entr -r go test ./...
```

### Run Tests with Verbose Output

```bash
go test -v ./...
```

This shows which tests pass/fail, helpful during the Red-Green cycle.

### Run Specific Package Tests

```bash
# Test HTTP handlers only
go test ./cmd/server

# Test user management only
go test ./internal/users

# Run a specific test
go test ./cmd/server -run TestHandleRoot
```

### Generate HTML Coverage Report

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

This opens a visual coverage report in your browser.

### TDD Workflow Commands

**Typical TDD cycle:**
```bash
# 1. Write test, see it fail (Red)
go test ./cmd/server -v -run TestNewHandler

# 2. Write implementation, see it pass (Green)
# ... implement code ...
go test ./cmd/server -v -run TestNewHandler

# 3. Refactor, ensure tests still pass
go test ./cmd/server -v
```

## TDD Best Practices Demonstrated

### 1. Test Names Describe Behavior

```go
TestHandleRoot                    // What does the handler do
TestAddUser                       // What does the function do
TestAddUserInvalidEmail           // Edge case clearly named
TestHandleHelloNoHeader           // Error case clearly named
```

**Why**: Test names serve as documentation of expected behavior.

### 2. Table-Driven Tests

The `users_test.go` demonstrates table-driven tests for comprehensive coverage:

```go
tests := map[string]struct {
    first         string
    last          string
    expected      *User
    expectedError error
}{
    "simple lookup": {
        first: "foo",
        last: "bar",
        expected: &testManager.users[0],
        expectedError: nil,
    },
    "no match lookup": {
        first: "nonexistent",
        last: "user",
        expected: nil,
        expectedError: ErrNoResultFound,
    },
}

for name, test := range tests {
    t.Run(name, func(t *testing.T) {
        result, err := testManager.GetUserByName(test.first, test.last)
        // assertions...
    })
}
```

**Why**: One test function covers multiple scenarios efficiently.

### 3. Testing Edge Cases First

Notice how error cases are tested alongside happy paths:

- `TestHandleHelloNoHeader` - Tests missing header
- `TestHandleJSONEmptyBody` - Tests empty request body
- `TestAddUserInvalidEmail` - Tests email validation
- `TestAddUserDuplicateName` - Tests duplicate prevention

**Why**: TDD ensures edge cases are considered upfront, not as an afterthought.

### 4. Using httptest Package

All HTTP handler tests use `net/http/httptest`:

```go
func TestHandleRoot(t *testing.T) {
    w := httptest.NewRecorder()        // Mock response writer
    r := httptest.NewRequest(http.MethodGet, "/", nil)  // Mock request
    handleRoot(w, r)                   // Call handler
    // Assertions on w.Code and w.Body
}
```

**Why**: Fast, isolated unit tests without a real server.

### 5. Test-Driven Error Handling

Error handling is driven by tests:

```go
// Test defines expected error behavior
func TestHandleHelloNoHeader(t *testing.T) {
    r := httptest.NewRequest(http.MethodGet, "/user/hello", nil)
    w := httptest.NewRecorder()
    handleHelloNoHeader(w, r)
    
    desiredCode := http.StatusBadRequest  // Test expects 400
    expectedMessage := []byte("invalid username provided\n")
    // assertions...
}
```

**Why**: Error handling is intentional, not accidental.

### 6. One Assertion Per Test (When Appropriate)

While some tests have multiple assertions, critical behaviors are tested separately:

```go
TestAddUser               // Tests successful addition
TestAddUserInvalidEmail   // Tests email validation
TestAddUserFirstName      // Tests first name validation
TestAddUserLastName       // Tests last name validation
```

**Why**: Clear failure messages and focused tests.

## Development Guide

### Following TDD for New Features

**Step 1: Write the Test First**

1. Create a test in the appropriate `*_test.go` file
2. Write the test to describe the desired behavior
3. Run `go test` - it should fail (🔴 Red)

```go
// cmd/server/main_test.go
func TestHandleHealth(t *testing.T) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodGet, "/health", nil)
    handleHealth(w, r)
    
    if w.Code != http.StatusOK {
        t.Errorf("expected status 200, got %d", w.Code)
    }
    
    expected := `{"status":"healthy"}`
    if w.Body.String() != expected {
        t.Errorf("expected %s, got %s", expected, w.Body.String())
    }
}
```

**Step 2: Write Minimal Implementation**

1. Implement just enough to make the test pass
2. Run `go test` - it should pass (🟢 Green)

```go
// cmd/server/main.go
func handleHealth(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`{"status":"healthy"}`))
}

// In main():
mux.HandleFunc("/health", handleHealth)
```

**Step 3: Refactor**

1. Improve code quality
2. Ensure all tests still pass
3. Run `go test ./...` after each refactor

```go
func handleHealth(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status": "healthy",
    })
}
```

### Code Style

This project follows Go's standard formatting:

```bash
# Format code
go fmt ./...

# Or using goimports
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .
```

### User Management API

The `internal/users` package demonstrates TDD for business logic:

#### Add User

**Tests**:
- `TestAddUser` - Successful user addition
- `TestAddUserInvalidEmail` - Email validation
- `TestAddUserFirstName` - First name validation
- `TestAddUserLastName` - Last name validation
- `TestAddUserDuplicateName` - Duplicate prevention

#### Get User by Name

**Test**: `TestGetUserByName` with table-driven tests covering:
- Simple lookup
- Last element lookup
- No match lookup
- Partial match lookup
- Empty name handling

See `internal/users/users_test.go` for complete test examples.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Contribution Guidelines

- Write tests for new features
- Ensure all tests pass (`go test ./...`)
- Follow Go code style guidelines
- Update documentation as needed
- Write clear commit messages

## License

This project is open source and available under the [MIT License](LICENSE).

## Acknowledgments

- Built with Go's excellent standard library
- Inspired by clean architecture principles

## Support

For issues, questions, or contributions, please open an issue on the [GitHub repository](https://github.com/kunalkumar-1/go-http/issues).

---

**Note**: This is a demonstration project. For production use, consider adding:
- Database persistence (PostgreSQL, MongoDB, etc.)
- Authentication and authorization
- Request validation middleware
- Rate limiting
- Monitoring and observability (Prometheus, Grafana)
- API versioning
- OpenAPI/Swagger documentation
