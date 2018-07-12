package data

import "database/sql"

// CoffeeMachine represents a single Coffee Machine product.
type CoffeeMachine struct {
	CoffeeMachineID int    `json:"coffee_machine_id,omitempty"`
	SizeID          int    `json:"size_id,omitempty"`
	SizeName        string `json:"size_name,omitempty"`
	SKU             string `json:"sku,omitempty"`
	ModelID         int    `json:"model_id,omitempty"`
	ModelName       string `json:"model_name,omitempty"`
	WaterLine       bool   `json:"water_line,omitempty"`
}

// Machines returns all machines.
func (db *DB) Machines(sizeID int) ([]CoffeeMachine, error) {
	var machines []CoffeeMachine
	var rows *sql.Rows
	var err error

	// Get the corresponding machine.
	if sizeID == 0 {
		rows, err = db.Query(coffeeMachineQ)
	} else {
		rows, err = db.Query(coffeeMachineQ+" WHERE CoffeeMachine.SizeID = ?", sizeID)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		m := new(CoffeeMachine)
		err = rows.Scan(&m.CoffeeMachineID, &m.SizeID, &m.SizeName,
			&m.SKU, &m.ModelID, &m.ModelName, &m.WaterLine)
		if err != nil {
			return nil, err
		}
		machines = append(machines, *m)
	}

	return machines, nil
}

// CrossSellMachines returns pods to be cross sold on the machine pages.
func (db *DB) CrossSellMachines(id int) ([]Pod, error) {
	var machine CoffeeMachine
	var pods []Pod

	// Get the corresponding machine.
	err := db.QueryRow(coffeeMachineQ+" where CoffeeMachineId = ?", id).
		Scan(&machine.CoffeeMachineID, &machine.SizeID, &machine.SizeName,
			&machine.SKU, &machine.ModelID, &machine.ModelName,
			&machine.WaterLine)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(machineCrossQ, machine.SizeID)
	defer rows.Close()

	for rows.Next() {
		pod := new(Pod)
		err = rows.Scan(&pod.PodID, &pod.FlavorID, &pod.FlavorName,
			&pod.SizeID, &pod.SizeName, &pod.SKU, &pod.Quantity)
		if err != nil {
			return nil, err
		}
		pods = append(pods, *pod)
	}

	return pods, nil
}
