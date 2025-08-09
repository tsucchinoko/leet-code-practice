# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview
This is a LeetCode practice repository for studying algorithms and data structures. Each problem is self-contained in its own directory under `/problem/` with language-specific implementations (currently Go).

## Project Structure
```
problem/
├── {problem-name}/
│   └── go/
│       ├── README.md     # Problem explanation in Japanese
│       ├── go.mod        # Independent Go module
│       ├── main.go       # Solution implementation
│       └── main_test.go  # Unit tests
```

## Common Commands

### Testing a Solution
Navigate to the problem's Go directory first:
```bash
cd problem/{problem-name}/go/

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover

# Run a specific test
go test -v -run TestFunctionName
```

### Running a Solution
```bash
cd problem/{problem-name}/go/
go run main.go
```

### Adding a New Problem
1. Create directory structure: `problem/{problem-name}/go/`
2. Initialize Go module: `go mod init github.com/tsucchinoko/{abbreviated-name}`
3. Create `main.go` with solution and `main_test.go` with tests
4. Optionally add `README.md` explaining the approach

## Module Naming Convention
Each problem uses an independent Go module with naming pattern:
- Module: `github.com/tsucchinoko/{abbreviated-problem-name}`
- Examples: `twosum`, `best-buy`, `create-hello-world-fn`

## Code Conventions
- Function names match LeetCode problem names (e.g., `TwoSum`, `MaxProfit`)
- Include Japanese comments explaining algorithm logic
- Use table-driven tests where appropriate
- Each solution should have comprehensive test coverage
- Document time and space complexity in comments

## Key Architecture Decisions
- **Independent Modules**: Each problem is a separate Go module (no shared dependencies)
- **Language Separation**: Problems organized by language under `/problem/{name}/{language}/`
- **Self-Contained**: Each problem includes implementation, tests, and documentation
- **Educational Focus**: Detailed Japanese documentation explaining approaches and complexity