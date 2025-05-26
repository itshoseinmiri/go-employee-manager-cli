package employee
import (
	"fmt"
	"errors"
)

type Manager struct {
	Employees []Employee
}

func (m *Manager) NextID() int {
	maxID := 0
	for _, e := range m.Employees {
		if e.ID > maxID {
			maxID = e.ID
		}
	}
	return maxID + 1
}

func (manager *Manager) AddEmployee(employee Employee) {
	manager.Employees = append(manager.Employees, employee)
}

func (manager *Manager) ListEmployees() []Employee {
	return manager.Employees
}

func (manager *Manager) DeleteEmoloyeeByID(id int) error {
	for i, employee := range manager.Employees {
		if employee.ID == id {
			manager.Employees = append(manager.Employees[:i], manager.Employees[i+1:]...)
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Employee with ID %d not found", id))
}