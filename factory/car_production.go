package factory

import "errors"

func (f *factory) ProduceCar(employeeName string, schematic Schematic) error {
	if schematic.Name != "Model T" {
		return errors.New("Only Model T's can be produced")
	}

	inventory := make(map[CarPart]Quantity)
	for part, qty := range f.state.carPartsInInventory {
		inventory[part] = qty
	}

	for part, qty := range schematic.Parts {
		inventory[part] -= qty

		if inventory[part] < 0 {
			return errors.New("The required carparts for building a " + schematic.Name + " are not in the inventory")
		}
	}

	f.recordThat(CarWasProduced{
		employeeName,
		schematic,
	})

	return nil
}

type Schematic struct {
	Name  string
	Parts map[CarPart]Quantity
}

func ModelTSchematic() Schematic {
	return Schematic{
		"Model T",
		map[CarPart]Quantity{
			"Wheel":           6,
			"Engine":          1,
			"Bits and pieces": 2,
		},
	}
}
