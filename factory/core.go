package factory

import "fmt"

/*
	factory employs people, manages a cargo bay, an inventory of carparts with the
	end goal of producing Model T's.
*/
type factory struct {
	state  *factoryState
	events []interface{}
}

/*
	FromEvents reconstructs a Factory from an array of events.
*/
func FromEvents(events []interface{}) *factory {
	return &factory{
		StateFromEvents(events),
		make([]interface{}, 0),
	}
}

/*
	recordThat an event has happened inside of a Factory.
*/
func (f *factory) recordThat(event interface{}) {
	f.events = append(f.events, event)
	f.state.Mutate(event)
}

type Quantity int
type CarPart string

type Command interface {
	fmt.Stringer
	Execute(*factory) error
}

type Event interface {
	fmt.Stringer
}

/*
	factoryState represents the internal state of a factory that is required to guard invariants.
*/
type factoryState struct {
	employeesCurrentlyInFactory            map[string]bool
	shipmentsInCargoBay                    []Shipment
	employeesThatInventorizedCargoBayToday map[string]bool
	carPartsInInventory                    map[CarPart]Quantity
}

/*
	StateFromEvents reconstructs factoryState from an array of events.
*/
func StateFromEvents(events []interface{}) *factoryState {
	f := &factoryState{
		make(map[string]bool),
		make([]Shipment, 0),
		make(map[string]bool),
		make(map[CarPart]Quantity),
	}

	for _, event := range events {
		f.Mutate(event)
	}

	return f
}

/*
	Mutate is the only way to change factoryState.
	In an event sourced system, it is important that events are the single source
	of truth. In order to ensure this, it is only possible to mutate state through
	events.
*/
func (state *factoryState) Mutate(event interface{}) {
	switch t := event.(type) {
	case EmployeeWasAssignedToFactory:
		state.applyEmployeeWasAssignedToFactory(t)
	case ShipmentArrivedInCargoBay:
		state.applyShipmentArrivedInCargoBay(t)
	case CargoBayWasInventorized:
		state.applyCargoBayWasInventorized(t)
	}
}

/*
	applyEmployeeWasAssignedToFactory tracks current employees in factory.
*/
func (state *factoryState) applyEmployeeWasAssignedToFactory(e EmployeeWasAssignedToFactory) {
	state.employeesCurrentlyInFactory[e.EmployeeName] = true
}

/*
	applyShipmentArrivedInCargoBay tracks current shipments in factory cargo bay.
*/
func (state *factoryState) applyShipmentArrivedInCargoBay(e ShipmentArrivedInCargoBay) {
	state.shipmentsInCargoBay = append(state.shipmentsInCargoBay, e.S)
}

/*
	applyCargoBayWasInventorized manages the inventory.
*/
func (state *factoryState) applyCargoBayWasInventorized(e CargoBayWasInventorized) {
	state.employeesThatInventorizedCargoBayToday[e.EmployeeName] = true

	for _, shipment := range state.shipmentsInCargoBay {
		for part, qty := range shipment {
			state.carPartsInInventory[part] += qty
		}
	}

	state.shipmentsInCargoBay = make([]Shipment, 0)
}
