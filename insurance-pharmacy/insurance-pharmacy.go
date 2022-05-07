package insurancePharmacy

import (
	"dispense/drug"
	"fmt"
)

type PharmacyInsurance struct {
	inventory map[drug.DrugId]int
}

func (phc PharmacyInsurance) CheckInventory(drug drug.DrugId) (int, error) {
	fmt.Printf("Check Inventory for %s\n", drug)

	if quantity, ok := phc.inventory[drug]; ok {
		return quantity, nil
	} else {
		return -1, fmt.Errorf("%s not found in inventory", drug)
	}
}

func (phc PharmacyInsurance) DispenseDrug(drug drug.DrugId) error {
	fmt.Printf("Dispense drug for %s\n", drug)
	amount, err := phc.CheckInventory(drug)

	if err != nil {
		return err
	}

	if amount > 0 {
		phc.inventory[drug] -= 1
	}

	return nil
}

func (phc PharmacyInsurance) LoadInventory(drug drug.DrugId, quantity int) {
	_, err := phc.CheckInventory(drug)

	if err != nil {
		phc.inventory[drug] = 0
	}

	phc.inventory[drug] += quantity
}

func (phc PharmacyInsurance) DisplayFullInventory() {
	for drug, quantity := range phc.Inventory {
		fmt.Printf("%s Quantity: %d\n", drug, quantity)
	}
}