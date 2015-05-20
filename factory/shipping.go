package factory

import (
	"errors"
	"fmt"
	"strings"
)

type Shipment map[CarPart]Quantity

func (s Shipment) String() string {
	list := make([]string, 0)
	for part, qty := range s {
		list = append(list, fmt.Sprint(qty, part))
	}

	return strings.Join(list, ",")
}

func (f *factory) SendShipmentToCargoBay(s Shipment) error {
	if len(f.state.employeesCurrentlyInFactory) == 0 {
		return errors.New("There has to be somebody at the factory in order to accept the shipment")
	}

	if len(s) == 0 {
		return errors.New("Empty shipments are not accepted")
	}

	if len(f.state.shipmentsInCargoBay) == 2 {
		return errors.New("Cargo bay has a maximum capacity of two shipments")
	}

	f.recordThat(ShipmentArrivedInCargoBay{
		s,
	})

	return nil
}

func (f *factory) InventorizeShipmentsInCargoBay(employeeName string) error {
	if len(f.state.shipmentsInCargoBay) == 0 {
		return errors.New("There are no shipments to inventorize in the cargo bay")
	}

	if _, ok := f.state.employeesThatInventorizedCargoBayToday[employeeName]; ok {
		return errors.New(employeeName + " has already inventorized the cargo bay today")
	}

	f.recordThat(CargoBayWasInventorized{
		employeeName,
	})

	return nil
}
