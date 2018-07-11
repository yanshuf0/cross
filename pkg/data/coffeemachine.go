package data

import "github.com/zls0/sh-cross-sell/pkg/models"

// CoffeeMachine represents a single Coffee Machine product.
type CoffeeMachine struct {
	CoffeeMachineID int    `json:"coffee_machine_id,omitempty"`
	SizeID          int    `json:"size_id,omitempty"`
	SizeName        string `json:"size_name,omitempty"`
	SKU             string `json:"sku,omitempty"`
	Description     string `json:"description,omitempty"`
	Waterline       bool   `json:"waterline,omitempty"`
}

// CrossSellMachines returns pods to be cross sold on the machine pages.
func (db *DB) CrossSellMachines(id int) ([]models.Pod, error) {
	var pods []models.Pod

	return pods, nil
}
