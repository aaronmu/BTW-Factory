package factory

import "fmt"

type EmployeeWasAssignedToFactory struct {
	EmployeeName string
}

func (e EmployeeWasAssignedToFactory) String() string {
	return fmt.Sprint(e.EmployeeName, " was assigned to factory")
}

type ShipmentArrivedInCargoBay struct {
	S Shipment
}

func (e ShipmentArrivedInCargoBay) String() string {
	return fmt.Sprint("Shipment containing ", e.S, " arrived in cargo bay")
}

type CargoBayWasInventorized struct {
	EmployeeName string
}

func (e CargoBayWasInventorized) String() string {
	return fmt.Sprint(e.EmployeeName, " inventorized the cargo bay")
}

type CarWasProduced struct {
	EmployeeName string
	Schematic    Schematic
}

func (e CarWasProduced) String() string {
	return fmt.Sprint(e.EmployeeName, " produced a car of model ", e.Schematic.Name)
}

type AssignEmployee struct {
	EmployeeName string
}

func (cmd AssignEmployee) String() string {
	return "Assign an employee named \"" + cmd.EmployeeName + "\" to factory"
}

func (cmd AssignEmployee) Execute(f *factory) error {
	return f.AssignEmployee(cmd.EmployeeName)
}

type SendShipmentToCargoBay struct {
	Shipment Shipment
}

func (cmd SendShipmentToCargoBay) String() string {
	return fmt.Sprint("Send a shipment containing ", cmd.Shipment, " to cargo bay")
}

func (cmd SendShipmentToCargoBay) Execute(f *factory) error {
	return f.SendShipmentToCargoBay(cmd.Shipment)
}

type InventorizeShipmentsInCargoBay struct {
	EmployeeName string
}

func (cmd InventorizeShipmentsInCargoBay) String() string {
	return cmd.EmployeeName + " inventorize the items currently in cargo bay"
}

func (cmd InventorizeShipmentsInCargoBay) Execute(f *factory) error {
	return f.InventorizeShipmentsInCargoBay(cmd.EmployeeName)
}

type ProduceCar struct {
	EmployeeName string
	Schema       Schematic
}

func (cmd ProduceCar) String() string {
	return cmd.EmployeeName + " produce a car named " + cmd.Schema.Name
}

func (cmd ProduceCar) Execute(f *factory) error {
	return f.ProduceCar(cmd.EmployeeName, cmd.Schema)
}
