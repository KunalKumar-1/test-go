# Test-Driven Development with Golang

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Tests](https://img.shields.io/badge/Tests-Passing-success?style=flat)](https://github.com/kunalkumar-1/go-http)

## Table of Contents

- [Introduction](#introduction)
- [What is TDD?](#what-is-tdd)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Running the Service](#running-the-service)
- [API Endpoints](#api-endpoints)
    - [Root & Welcome](#root--welcome)
    - [Parameter-Based Routes](#parameter-based-routes)
    - [Header-Based Routes](#header-based-routes)
    - [JSON Routes](#json-routes)
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)
- [TDD Workflow](#tdd-workflow)
- [Service Methods](#service-methods)
- [Testing](#testing)
- [TDD Best Practices](#tdd-best-practices)
- [Development](#development)
- [Deployment](#deployment)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This repository demonstrates production-ready **Test-Driven Development (TDD)** practices in Golang. The project showcases how to build an HTTP server and user management system by writing tests first, following the Red-Green-Refactor cycle. Every feature in this codebase was developed test-first, making the tests serve as both specification and documentation.

## What is TDD?

Test-Driven Development is a software development methodology where tests are written before the implementation code. The process follows a simple three-step cycle:

1. **🔴 Red**: Write a failing test that defines the desired behavior
2. **🟢 Green**: Write the minimum code necessary to make the test pass
3. **♻️ Refactor**: Improve the code while ensuring all tests remain passing

### Why TDD?

- **Better Design**: Writing tests first leads to more modular, testable code
- **Documentation**: Tests serve as living documentation of expected behavior
- **Confidence**: Comprehensive test coverage from the start
- **Fewer Bugs**: Edge cases are considered upfront, not as afterthoughts
- **Faster Debugging**: When tests fail, you know exactly what broke

## Architecture

The architecture demonstrates TDD principles across multiple layers:

- **HTTP Handlers**: Route handlers driven by HTTP test specifications
- **User Management**: Business logic developed through unit tests
- **Validation Layer**: Input validation implemented test-first
- **Error Handling**: Error cases defined and tested before implementation
- **API Design**: API contracts established through test cases

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- [Golang](https://golang.org/doc/install) 1.25 or higher
- Basic understanding of Go testing framework
- Familiarity with HTTP concepts

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/kunalkumar-1/go-http.git
    cd go-http
    ```

2. Install dependencies:

    ```sh
    go mod download
    ```

3. Verify installation by running tests:

    ```sh
    go test ./...
    ```

### Running the Service

1. Start the HTTP server:

    ```sh
    go run ./cmd/server/main.go
    ```

2. The service will be available at:
    - API Server: `http://localhost:4000`

3. Test the server is running:

    ```sh
    curl http://localhost:4000/
    ```

## API Endpoints

### Root & Welcome

#### Get Welcome Message
```http
GET /
```

**Response:**
```
Welcome to our HomePage!
```

**Test Coverage:**
- `TestHandleRoot` - Validates welcome message and status code

---

#### Get Goodbye Message
```http
GET /goodbye
```

**Response:**
```
Goodbye world is served at goodbye
```

**Test Coverage:**
- `TestHandleGoodbye` - Validates goodbye message

---

### Parameter-Based Routes

#### Hello with Query Parameters
```http
GET /hello/?user=John
```

**Query Parameters:**
- `user` (optional): Username to greet. Defaults to "User" if not provided.

**Response:**
```
Hello John!
```

**Test Coverage:**
- `TestHandleHelloParameterized` - With user parameter
- `TestHandleHelloNoParameterized` - Without user parameter (defaults)
- `TestHandleHelloWrongParameterized` - Invalid parameter handling

**Example:**
```sh
# With parameter
curl "http://localhost:4000/hello/?user=Alice"

# Without parameter (uses default)
curl "http://localhost:4000/hello/"
```

---

#### Hello with Path Variables
```http
GET /responses/{user}/hello/
```

**Path Parameters:**
- `user` (required): Username to greet

**Response:**
```
Hello Alice!
```

**Test Coverage:**
- `TestHandleUserResponsesHello` - Path variable extraction and response

**Example:**
```sh
curl http://localhost:4000/responses/Bob/hello/
```

---

### Header-Based Routes

#### Hello with Header
```http
GET /user/hello
Headers:
  user: Charlie
```

**Headers:**
- `user` (required): Username to greet

**Response:**
```
Hello Charlie!
```

**Error Response (400 Bad Request):**
```
invalid username provided
```

**Test Coverage:**
- `TestHandleHelloHeader` - Valid header handling
- `TestHandleHelloNoHeader` - Missing header error case

**Example:**
```sh
# Success case
curl -H "user: Charlie" http://localhost:4000/user/hello

# Error case (missing header)
curl http://localhost:4000/user/hello
```

---

### JSON Routes

#### Hello with JSON Payload
```http
POST /json
Content-Type: application/json

{
  "Name": "David"
}
```

**Request Body:**
```json
{
  "Name": "string (required)"
}
```

**Response:**
```
Hello David!
```

**Error Responses:**

**400 Bad Request** - Empty request body:
```
empty request body
```

**400 Bad Request** - Invalid JSON or missing Name field:
```
invalid request body!
```

**Test Coverage:**
- `TestHandleJSON` - Valid JSON payload
- `TestHandleJSONEmptyBody` - Empty body error handling
- `TestHandleJSONEmptyNameFeild` - Missing Name field validation

**Example:**
```sh
# Success case
curl -X POST http://localhost:4000/json \
  -H "Content-Type: application/json" \
  -d '{"Name":"David"}'

# Error case (empty body)
curl -X POST http://localhost:4000/json \
  -H "Content-Type: application/json" \
  -d '{}'

# Error case (invalid JSON)
curl -X POST http://localhost:4000/json \
  -H "Content-Type: application/json" \
  -d 'invalid'
```

---

## Project Structure

```
go-http/
├── cmd/
│   └── server/
│       ├── main.go                # HTTP server implementation
│       └── main_test.go           # HTTP handler tests (written first)
├── internal/
│   └── users/
│       ├── users.go               # User management implementation
│       └── users_test.go          # User management tests (written first)
├── go.mod                         # Go module dependencies
├── go.sum                         # Go module checksums
└── README.md                      # Project documentation
```

### Key Files

- **`main_test.go`**: HTTP handler tests that drove the API design
- **`main.go`**: HTTP server implementation written to satisfy tests
- **`users_test.go`**: User management tests defining business logic
- **`users.go`**: User management implementation

## Technologies Used

- **Golang**: Primary language (1.25+)
- **net/http**: Standard library HTTP server
- **net/http/httptest**: HTTP testing utilities
- **testing**: Go's built-in testing framework
- **net/mail**: Email validation

## TDD Workflow

### Example: Building the Root Handler

**Step 1: 🔴 Red - Write a Failing Test**

```go
// cmd/server/main_test.go
func TestHandleRoot(t *testing.T) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodGet, "/", nil)
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

Run test: `go test ./cmd/server` - **FAILS** ❌ (handleRoot doesn't exist)

**Step 2: 🟢 Green - Write Minimum Code to Pass**

```go
// cmd/server/main.go
func handleRoot(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to our HomePage!\n"))
}
```

Run test: `go test ./cmd/server` - **PASSES** ✅

**Step 3: ♻️ Refactor - Improve Code Quality**

```go
// cmd/server/main.go
func handleRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Requested path:", r.URL.Path)
    _, err := w.Write([]byte("Welcome to our HomePage!\n"))
    if err != nil {
        slog.Error("Error serving the root handler: " + err.Error())
        return
    }
}
```

Run test: `go test ./cmd/server` - **STILL PASSES** ✅

### Example: Building User Management

**Step 1: 🔴 Red - Test Adding a User**

```go
// internal/users/users_test.go
func TestAddUser(t *testing.T) {
    testManager := NewManager()
    
    err := testManager.AddUser("John", "Doe", "john@example.com")
    if err != nil {
        t.Fatalf("failed to add user: %v", err)
    }
    
    if len(testManager.users) != 1 {
        t.Fatalf("expected 1 user, got %v", len(testManager.users))
    }
}
```

Run test: `go test ./internal/users` - **FAILS** ❌

**Step 2: 🟢 Green - Implement AddUser**

```go
// internal/users/users.go
func (m *Manager) AddUser(firstName string, lastName string, email string) error {
    newUser := User{
        FirstName: firstName,
        LastName:  lastName,
        Email:     mail.Address{Address: email},
    }
    m.users = append(m.users, newUser)
    return nil
}
```

Run test: `go test ./internal/users` - **PASSES** ✅

**Step 3: ♻️ Refactor - Add Validation**

First, write a test for email validation:

```go
// internal/users/users_test.go
func TestAddUserInvalidEmail(t *testing.T) {
    testManager := NewManager()
    err := testManager.AddUser("John", "Doe", "invalid-email")
    if err == nil {
        t.Errorf("expected error for invalid email")
    }
}
```

Then refactor AddUser to include validation:

```go
// internal/users/users.go
func (m *Manager) AddUser(firstName string, lastName string, email string) error {
    // Validate email
    _, err := mail.ParseAddress(email)
    if err != nil {
        return fmt.Errorf("invalid email address: %w", err)
    }
    
    newUser := User{
        FirstName: firstName,
        LastName:  lastName,
        Email:     mail.Address{Address: email},
    }
    m.users = append(m.users, newUser)
    return nil
}
```

Run tests: `go test ./internal/users` - **ALL PASS** ✅

## Service Methods

### HTTP Handler Methods

> ### Root Handlers (`cmd/server/main.go`)
```go
func handleRoot(w http.ResponseWriter, r *http.Request)
func handleGoodbye(w http.ResponseWriter, r *http.Request)
```

> ### Parameterized Handlers
```go
func handleHelloParameterized(w http.ResponseWriter, r *http.Request)
func handleUserResponsesHello(w http.ResponseWriter, r *http.Request)
```

> ### Header-Based Handlers
```go
func handleHelloNoHeader(w http.ResponseWriter, r *http.Request)
```

> ### JSON Handlers
```go
func handleJSON(w http.ResponseWriter, r *http.Request)
```

### User Management Methods

> ### User Operations (`internal/users/users.go`)
```go
func NewManager() *Manager
func (m *Manager) AddUser(firstName string, lastName string, email string) error
func (m *Manager) GetUserByName(firstName string, lastName string) (*User, error)
func (m *Manager) GetAllUsers() []User
func (m *Manager) DeleteUser(firstName string, lastName string) error
```

### Test Helper Methods

> ### HTTP Test Helpers (`cmd/server/main_test.go`)
```go
func setupTestRequest(method string, path string, body io.Reader) (*httptest.ResponseRecorder, *http.Request)
func assertStatusCode(t *testing.T, got int, want int)
func assertResponseBody(t *testing.T, got []byte, want []byte)
```

## Testing

### Running Tests

```sh
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run specific package tests
go test ./cmd/server
go test ./internal/users

# Run specific test
go test ./cmd/server -run TestHandleRoot

# Run tests with race detection
go test -race ./...
```

### Coverage Report

```sh
# Generate coverage profile
go test -coverprofile=coverage.out ./...

# View coverage in terminal
go tool cover -func=coverage.out

# Generate HTML coverage report
go tool cover -html=coverage.out
```

### Test Output Example

```
PASS
coverage: 95.2% of statements
ok      github.com/kunalkumar-1/go-http/cmd/server      0.156s
ok      github.com/kunalkumar-1/go-http/internal/users  0.089s
```

### Test Structure

```
go-http/
├── cmd/server/
│   ├── main.go
│   └── main_test.go           # HTTP handler tests
│       ├── TestHandleRoot
│       ├── TestHandleGoodbye
│       ├── TestHandleHelloParameterized
│       ├── TestHandleHelloNoParameterized
│       ├── TestHandleHelloWrongParameterized
│       ├── TestHandleUserResponsesHello
│       ├── TestHandleHelloHeader
│       ├── TestHandleHelloNoHeader
│       ├── TestHandleJSON
│       ├── TestHandleJSONEmptyBody
│       └── TestHandleJSONEmptyNameFeild
└── internal/users/
    ├── users.go
    └── users_test.go          # User management tests
        ├── TestAddUser
        ├── TestAddUserInvalidEmail
        ├── TestAddUserFirstName
        ├── TestAddUserLastName
        ├── TestAddUserDuplicateName
        ├── TestGetUserByName
        └── TestGetAllUsers
```

## TDD Best Practices

### 1. Test Names Document Behavior

Tests are named to clearly describe what behavior they're testing:

```go
TestHandleRoot                    // Root handler behavior
TestHandleHelloNoHeader           // Error case: missing header
TestAddUserInvalidEmail           // Validation: invalid email
TestHandleJSONEmptyBody           // Error case: empty request body
```

### 2. Table-Driven Tests

Use table-driven tests for comprehensive scenario coverage:

```go
tests := map[string]struct {
    firstName     string
    lastName      string
    expected      *User
    expectedError error
}{
    "simple lookup": {
        firstName: "John",
        lastName: "Doe",
        expected: &User{...},
        expectedError: nil,
    },
    "no match lookup": {
        firstName: "NonExistent",
        lastName: "User",
        expected: nil,
        expectedError: ErrNoResultFound,
    },
}

for name, test := range tests {
    t.Run(name, func(t *testing.T) {
        result, err := testManager.GetUserByName(test.firstName, test.lastName)
        // assertions...
    })
}
```

### 3. Test Edge Cases First

Error cases and edge cases are tested alongside happy paths:

- Missing required fields
- Invalid input formats
- Empty request bodies
- Duplicate entries
- Boundary conditions

### 4. Use httptest for HTTP Testing

All HTTP handlers use `net/http/httptest` for fast, isolated tests:

```go
func TestHandleJSON(t *testing.T) {
    w := httptest.NewRecorder()
    r := httptest.NewRequest(http.MethodPost, "/json", bytes.NewBuffer(payload))
    handleJSON(w, r)
    
    // Assertions on w.Code and w.Body
}
```

### 5. One Test Per Behavior

Each test focuses on a single behavior or scenario:

```go
TestAddUser               // Happy path: successful addition
TestAddUserInvalidEmail   // Error path: email validation
TestAddUserFirstName      // Validation: first name requirements
TestAddUserDuplicateName  // Error path: duplicate prevention
```

### 6. Tests as Documentation

The test files serve as the specification:

- `main_test.go` defines the HTTP API contract
- `users_test.go` defines user management behavior
- Reading tests tells you exactly what the system does

## Development

### TDD Development Workflow

**For New Features:**

1. **Write the Test First**
   ```sh
   # Create test in appropriate *_test.go file
   # Run: go test ./path/to/package
   # Should FAIL (🔴 Red)
   ```

2. **Implement Minimum Code**
   ```sh
   # Write just enough code to pass
   # Run: go test ./path/to/package
   # Should PASS (🟢 Green)
   ```

3. **Refactor**
   ```sh
   # Improve code quality
   # Run: go test ./...
   # All tests should still PASS (♻️ Refactor)
   ```

### Live Reload with Air

Install Air for development with live reload:

```sh
go install github.com/cosmtrek/air@latest
```

Run with live reload:

```sh
air
```

### Watch Mode for Tests

Run tests automatically on file changes:

```sh
# Using entr (Unix/Mac)
brew install entr
ls *.go **/*.go | entr -r go test ./...

# Or use gotestsum
go install gotest.tools/gotestsum@latest
gotestsum --watch
```

### Code Formatting

```sh
# Format all Go files
go fmt ./...

# Using goimports (recommended)
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .
```

### Linting

```sh
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run
```

## Deployment

### Docker Deployment

Create a `Dockerfile`:

```dockerfile
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 4000
CMD ["./server"]
```

Build and run:

```sh
# Build image
docker build -t go-http-tdd .

# Run container
docker run -p 4000:4000 go-http-tdd
```

### Docker Compose

Create a `docker-compose.yml`:

```yaml
version: '3.8'
services:
  app:
    build: .
    ports:
      - "4000:4000"
    environment:
      - PORT=4000
    restart: unless-stopped
```

Run with Docker Compose:

```sh
docker-compose up -d --build
```

### Production Considerations

- **Reverse Proxy**: Use nginx or Caddy as reverse proxy
- **HTTPS/TLS**: Configure SSL certificates (Let's Encrypt)
- **Environment Variables**: Externalize configuration
- **Graceful Shutdown**: Implement proper shutdown handling
- **Health Checks**: Add `/health` endpoint for monitoring
- **Logging**: Structured logging with log levels
- **Rate Limiting**: Add rate limiting middleware
- **CORS**: Configure CORS for API access
- **Metrics**: Add Prometheus metrics endpoint

## Roadmap

- [x] TDD-driven HTTP handlers
- [x] User management with validation
- [x] Comprehensive test coverage
- [x] Table-driven tests
- [x] HTTP testing with httptest
- [ ] Middleware support (logging, recovery)
- [ ] Database integration (PostgreSQL)
- [ ] JWT authentication
- [ ] API versioning
- [ ] OpenAPI/Swagger documentation
- [ ] Integration tests
- [ ] Benchmark tests
- [ ] CI/CD pipeline (GitHub Actions)
- [ ] Docker multi-stage builds
- [ ] Kubernetes deployment manifests
- [ ] Monitoring and observability (Prometheus/Grafana)
- [ ] Load testing suite
- [ ] Security scanning

## Contributing

Contributions are welcome! When contributing, please follow TDD principles:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. **Write tests first** for your feature
4. Implement the feature to make tests pass
5. Ensure all tests pass (`go test ./...`)
6. Commit your changes (`git commit -m 'Add some amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Contribution Guidelines

- **Always write tests first** - Follow the Red-Green-Refactor cycle
- All tests must pass before submitting PR
- Maintain or improve test coverage
- Follow Go code style guidelines (`gofmt`, `goimports`)
- Write meaningful commit messages
- Update documentation for new features
- Include examples in tests

### Code Review Checklist

- [ ] Tests written before implementation
- [ ] All tests passing
- [ ] Code coverage maintained/improved
- [ ] No commented-out code
- [ ] Clear, descriptive test names
- [ ] Edge cases covered
- [ ] Documentation updated

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Built with Test-Driven Development** 🔴 🟢 ♻️

*Every line of code in this project was written to satisfy a test.*