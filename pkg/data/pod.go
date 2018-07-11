package data

import "github.com/zls0/sh-cross-sell/pkg/models"

// Pod represents a single coffee pod product.
type Pod struct {
	PodID       int    `json:"pod_id,omitempty"`
	SizeID      int    `json:"size_id,omitempty"`
	SizeName    string `json:"size_name,omitempty"`
	FlavorID    int    `json:"flavor_id,omitempty"`
	FlavorName  string `json:"flavor_name,omitempty"`
	SKU         string `json:"sku,omitempty"`
	Description string `json:"description,omitempty"`
}

// CrossSellPods returns different pack sizes for each flavor
// (within a machine size).
func (db *DB) CrossSellPods(id int) ([]models.Pod, error) {
	var pods []models.Pod

	return pods, nil
}
