# Go Employee Management Cli

## Overview
This is a simple command-line Employee Management System written in Go. It allows users to add new employees to a JSON file and list all employees stored in the system.

## Features
- Add employees with details including ID, Name, Age, Position, Department, and Salary
- List all employees with their complete details
- Persist employee data in a JSON file (`data/employees.json`)
- Automatic ID generation for new employees

## Prerequisites
- Go 1.16 or higher
- Basic understanding of command-line interfaces

## Installation
1. Clone or download this repository
2. Ensure you have Go installed on your system
3. Create a `data` directory in the project root
4. Create an empty `employees.json` file in the `data` directory or the program will create one automatically

## Usage
The program supports two subcommands: `add` and `list`

### Adding an Employee
To add a new employee, use the `add` subcommand with the required flags:
```bash
go run main.go add --name "John Doe" --age 30 --position "Developer" --department "IT" --salary 75000
```

Required flags:
- `--name`: Employee's full name (string)
- `--age`: Employee's age (integer)
- `--position`: Employee's job title (string)

Optional flags:
- `--department`: Employee's department (string)
- `--salary`: Employee's annual salary (integer)

### Listing Employees
To list all employees:
```bash
go run main.go list
```

This will display all employees with their details in the format:
```
ID: [ID], Name: [Name], Age: [Age], Position: [Position], Department: [Department], Salary: [Salary]
```

## Project Structure
```
employee-management/
├── data/
│   └── employees.json
├── employee/
│   └── employee.go
├── utils/
│   └── utils.go
└── main.go
```

## Error Handling
- The program validates required flags for the `add` command
- It handles file I/O errors when reading/writing to the JSON file
- If the JSON file doesn't exist, it will be created automatically
- Invalid subcommands will display an error message

## Example
```bash
# Add a new employee
go run main.go add --name "Jane Smith" --age 25 --position "Designer" --department "Creative" --salary 65000

# List all employees
go run main.go list
```

## Notes
- Employee data is stored in `data/employees.json`
- The program automatically generates unique IDs for new employees
- Salary is stored as a float64 but accepts integer input
- The program exits with appropriate error codes when errors occur