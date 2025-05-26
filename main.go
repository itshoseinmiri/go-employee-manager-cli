package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/itshoseinmiri/employee-management/employee"
	"github.com/itshoseinmiri/employee-management/utils"
)

func main() {
	manager := employee.Manager{}
	_ = utils.LoadFromFile("data/employees.json", &manager.Employees)

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	name := addCmd.String("name", "", "Employee Name")
	age := addCmd.Int("age", 0, "Employee Age")
	position := addCmd.String("position", "", "Employee Position")
	department := addCmd.String("department", "", "Department")
	salary := addCmd.Int("salary", 0, "Salary")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'add' or 'list' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])

		if *name == "" || *age == 0 || *position == "" {
			fmt.Println("Please provide --id, --name, --age, and --position flags")
			os.Exit(1)
		}

		newID := manager.NextID()

		newEmp := employee.Employee{
			ID:         newID,
			Name:       *name,
			Age:        *age,
			Position:   *position,
			Department: *department,
			Salary:     float64(*salary),
		}

		manager.AddEmployee(newEmp)
		err := utils.SaveToFile("data/employees.json", manager.Employees)
		if err != nil {
			fmt.Println("Error saving:", err)
			os.Exit(1)
		}
		fmt.Println("Employee added successfully!")

	case "list":
		fmt.Println("Employee List:")
		listCmd.Parse(os.Args[2:])
		if len(manager.ListEmployees()) == 0 {
			fmt.Println("No employees found.")
			os.Exit(0)
		}
		for _, emp := range manager.ListEmployees() {
			fmt.Printf("ID: %d, Name: %s, Age: %d, Position: %s, Department: %s, Salary: %.2f\n",
				emp.ID, emp.Name, emp.Age, emp.Position, emp.Department, emp.Salary)
		}

	default:
		fmt.Println("Unknown subcommand:", os.Args[1])
		os.Exit(1)
	}
}