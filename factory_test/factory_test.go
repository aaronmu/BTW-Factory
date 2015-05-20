package factory_test

import (
	"testing"

	"github.com/aaronmu/BTW-Factory/factory"
)

func TestAssigningAnEmployeeToAFactory(t *testing.T) {
	s := factory.NewFactorySpecification("Assigning an employee to a factory", t)

	s.When(factory.AssignEmployee{"Jeff"})

	s.Then(factory.EmployeeWasAssignedToFactory{"Jeff"})

	s.Run(t)
}

func TestAssigningMultipleEmployeesToAFactory(t *testing.T) {
	s := factory.NewFactorySpecification("Assigning multiple employees to a factory", t)

	s.When(factory.AssignEmployee{"John"})
	s.When(factory.AssignEmployee{"Jeff"})

	s.Then(factory.EmployeeWasAssignedToFactory{"Jeff"})

	s.Run(t)
}

func TestAssigningTwoEmployeesThatHaveTheSameNameToAFactory(t *testing.T) {
	s := factory.NewFactorySpecification("Assigning two employees to a factory that have the same name", t)

	s.When(factory.AssignEmployee{"John"})
	s.When(factory.AssignEmployee{"Jeff"})
	s.When(factory.AssignEmployee{"John"})

	s.ThenError("An employee named \"John\" was already assigned to the factory")

	s.Run(t)
}

func TestEmployeeNamedBenderIsNotAllowedToBeAssigned(t *testing.T) {
	s := factory.NewFactorySpecification("Assigning an employee named \"Bender\" to a factory is not allowed.", t)

	s.When(factory.AssignEmployee{"Bender"})

	s.ThenError("Employees named \"Bender\" are trouble")

	s.Run(t)
}

func TestSendingAShipmentToAnEmptyFactory(t *testing.T) {
	s := factory.NewFactorySpecification("Sending a shipment to an empty factory", t)

	s.When(factory.SendShipmentToCargoBay{factory.Shipment{"Engine": 1}})

	s.ThenError("There has to be somebody at the factory in order to accept the shipment")

	s.Run(t)
}

func TestEmptyShipmentsAreNotAllowed(t *testing.T) {
	s := factory.NewFactorySpecification("Factory denies empty shipments", t)

	s.Given(factory.EmployeeWasAssignedToFactory{"Aaron"})

	s.When(factory.SendShipmentToCargoBay{factory.Shipment{}})

	s.ThenError("Empty shipments are not accepted")

	s.Run(t)
}

func TestShipmentArrivesInCargoBay(t *testing.T) {
	s := factory.NewFactorySpecification("factory.Shipment arrives in cargo bay", t)

	s.Given(factory.EmployeeWasAssignedToFactory{"Aaron"})

	s.When(factory.SendShipmentToCargoBay{factory.Shipment{"engine": 1}})

	s.Then(factory.ShipmentArrivedInCargoBay{
		factory.Shipment{"engine": 1},
	})

	s.Run(t)
}

func TestCargoBayHasAMaximumCapacityOfTwoShipments(t *testing.T) {
	s := factory.NewFactorySpecification("A factory's cargo bay has a maximum capacity of two shipments", t)

	s.Given(
		factory.EmployeeWasAssignedToFactory{"Aaron"},
		factory.ShipmentArrivedInCargoBay{
			factory.Shipment{"Engine": 1},
		},
		factory.ShipmentArrivedInCargoBay{
			factory.Shipment{"Rear view mirror": 1},
		},
	)

	s.When(factory.SendShipmentToCargoBay{factory.Shipment{"Engine": 1}})

	s.ThenError("Cargo bay has a maximum capacity of two shipments")

	s.Run(t)
}

func TestShipmentsInCargoBayCanBeInventorized(t *testing.T) {
	s := factory.NewFactorySpecification("An employee can inventorize a cargo bay containing shipments", t)

	s.Given(
		factory.EmployeeWasAssignedToFactory{"Jeff"},
		factory.ShipmentArrivedInCargoBay{
			factory.Shipment{"Engine": 1},
		},
	)

	s.When(factory.InventorizeShipmentsInCargoBay{"Jeff"})

	s.Then(factory.CargoBayWasInventorized{"Jeff"})

	s.Run(t)
}

func TestEmployeeCanNotInventorizeAnEmptyCargoBay(t *testing.T) {
	s := factory.NewFactorySpecification("An employee cannot inventorize an empty cargo bay", t)

	s.Given(factory.EmployeeWasAssignedToFactory{"Jeff"})

	s.When(factory.InventorizeShipmentsInCargoBay{"Jeff"})

	s.ThenError("There are no shipments to inventorize in the cargo bay")

	s.Run(t)
}

func TestEmployeeCanInventorizeTheCargoBayOnceADay(t *testing.T) {
	s := factory.NewFactorySpecification("An employee can only inventorize the cargo bay once a day", t)

	s.Given(
		factory.ShipmentArrivedInCargoBay{
			factory.Shipment{"Engine": 1},
		},
		factory.CargoBayWasInventorized{"Jeff"},
		factory.ShipmentArrivedInCargoBay{
			factory.Shipment{"Wheel": 7},
		},
	)

	s.When(factory.InventorizeShipmentsInCargoBay{"Jeff"})

	s.ThenError("Jeff has already inventorized the cargo bay today")

	s.Run(t)
}

func TestOnlyModelTCanBeProduced(t *testing.T) {
	s := factory.NewFactorySpecification("Only Model T's can be produced by the factor", t)

	s.When(factory.ProduceCar{
		"Jeff",
		factory.Schematic{
			"Opel Corsa",
			map[factory.CarPart]factory.Quantity{
				"Engine": 1,
				"Wheel":  2,
			},
		},
	})

	s.ThenError("Only Model T's can be produced")

	s.Run(t)
}

func TestEmployeeCanProduceACar(t *testing.T) {
	s := factory.NewFactorySpecification("An employee can produce a car", t)

	s.Given(
		factory.ShipmentArrivedInCargoBay{
			factory.Shipment{
				"Wheel":  3,
				"Engine": 1,
			},
		},
		factory.ShipmentArrivedInCargoBay{
			factory.Shipment{
				"Wheel":           3,
				"Bits and pieces": 2,
			},
		},
		factory.CargoBayWasInventorized{"Jeff"},
	)

	s.When(factory.ProduceCar{
		"Jeff",
		factory.ModelTSchematic(),
	})

	s.Then(factory.CarWasProduced{
		"Jeff",
		factory.ModelTSchematic(),
	})

	s.Run(t)
}

func TestAllCarPartsInASchematicNeedToBeInTheInventoryBeforeCarCanBeProduced(t *testing.T) {
	s := factory.NewFactorySpecification("In order to produce a car, all carparts in a schematic need to be inventorized", t)

	s.When(factory.ProduceCar{
		"Jeff",
		factory.ModelTSchematic(),
	})

	s.ThenError("The required carparts for building a Model T are not in the inventory")

	s.Run(t)
}
