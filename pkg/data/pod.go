package data

import (
	"database/sql"
)

// Pod represents a single coffee pod product.
type Pod struct {
	PodID      int    `json:"pod_id,omitempty"`
	SizeID     int    `json:"size_id,omitempty"`
	SizeName   string `json:"size_name,omitempty"`
	FlavorID   int    `json:"flavor_id,omitempty"`
	FlavorName string `json:"flavor_name,omitempty"`
	SKU        string `json:"sku,omitempty"`
	Quantity   int    `json:"quantity,omitempty"`
}

// Pods returns a slice of pods, it accepts flavorID and or sizeID to filter the
// pods.
func (db *DB) Pods(flavorID int, sizeID int) ([]Pod, error) {
	var pods []Pod
	var rows *sql.Rows
	var err error

	switch {
	case flavorID != 0 && sizeID != 0:
		query := podQ + " WHERE Pod.FlavorID = ? AND Pod.SizeID = ?"
		rows, err = db.Query(query, flavorID, sizeID)
		if err != nil {
			return nil, err
		}
	case flavorID != 0:
		query := podQ + " WHERE Pod.FlavorID = ?"
		rows, err = db.Query(query, flavorID)
		if err != nil {
			return nil, err
		}
	case sizeID != 0:
		query := podQ + " WHERE Pod.SizeID = ?"
		rows, err = db.Query(query, sizeID)
		if err != nil {
			return nil, err
		}
	default:
		rows, err = db.Query(podQ)
		if err != nil {
			return nil, err
		}
	}
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

// CrossSellPods returns different pack sizes for each flavor
// (within a machine size).
func (db *DB) CrossSellPods(id int) ([]Pod, error) {
	var pods []Pod
	var pod Pod

	err := db.QueryRow(podQ+" WHERE PodID = ?", id).Scan(&pod.PodID,
		&pod.FlavorID, &pod.FlavorName, &pod.SizeID, &pod.SizeName,
		&pod.SKU, &pod.Quantity)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(podQ+" WHERE Pod.FlavorID = ? AND Pod.SizeID = ? AND PodID != ?",
		pod.FlavorID, pod.SizeID, pod.PodID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pod := new(Pod)
		if err = rows.Scan(&pod.PodID, &pod.FlavorID, &pod.FlavorName,
			&pod.SizeID, &pod.SizeName, &pod.SKU, &pod.Quantity); err != nil {
			return nil, err
		}
		pods = append(pods, *pod)
	}

	return pods, nil
}
