package factory

import "errors"

func (f *factory) AssignEmployee(employeeName string) error {
	if employeeName == "Bender" {
		return errors.New("Employees named \"Bender\" are trouble")
	}

	if _, exists := f.state.employeesCurrentlyInFactory[employeeName]; exists {
		return errors.New("An employee named \"" + employeeName + "\" was already assigned to the factory")
	}

	f.recordThat(EmployeeWasAssignedToFactory{
		employeeName,
	})

	return nil
}
