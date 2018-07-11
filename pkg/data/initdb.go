package data

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// InitDB will create the necessary records and tables for the the sqlite
// database. This will ensure the app is easy to run from others computers.
func (db *DB) initDB() error {
	machineTable, err := db.Prepare(coffeeQ)
	if err != nil {
		return err
	}
	flavorTable, err := db.Prepare(flavorQ)
	if err != nil {
		return err
	}
	sizeTable, err := db.Prepare(sizeQ)
	if err != nil {
		return err
	}
	modelTable, err := db.Prepare(modelQ)
	if err != nil {
		return err
	}
	podTable, err := db.Prepare(podQ)
	if err != nil {
		return err
	}
	if _, err = flavorTable.Exec(); err != nil {
		return err
	}
	if _, err = sizeTable.Exec(); err != nil {
		return err
	}
	if _, err = modelTable.Exec(); err != nil {
		return err
	}
	if _, err = machineTable.Exec(); err != nil {
		return err
	}
	if _, err = podTable.Exec(); err != nil {
		return err
	}
	// Populates size table
	sizes := []string{"espresso", "small", "large"}
	for i, v := range sizes {
		stmt, err := db.Prepare(`INSERT INTO Size (SizeID, SizeName) VALUES (?, ?)`)
		if err != nil {
			return err
		}
		if _, err = stmt.Exec(i, v); err != nil {
			return err
		}
	}
	// Populates flavor table
	flavors := []string{"vanilla", "caramel", "psl", "mocha", "hazelnut"}
	for i, v := range flavors {
		stmt, err := db.Prepare(`INSERT INTO Flavor (FlavorID, FlavorName) VALUES (?, ?)`)
		if err != nil {
			return err
		}
		if _, err = stmt.Exec(i, v); err != nil {
			return err
		}
	}
	// Populates model table
	models := []string{"base model", "premium model", "deluxe model"}
	for i, v := range models {
		stmt, err := db.Prepare(`INSERT INTO Model (ModelID, ModelName) VALUES (?, ?)`)
		if err != nil {
			return err
		}
		if _, err = stmt.Exec(i, v); err != nil {
			return err
		}
	}
	err = db.parseProducts()
	if err != nil {
		log.Fatalf("unable to parse products, err: %v", err)
	}
	return nil
}

// parseProducts parses the list of products to facilitate adding to the db.
func (db *DB) parseProducts() error {
	f, err := os.Open(assetsDir + "/products.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var err error
		product := sc.Text()
		id := string(product[:2])
		switch {
		case id == "CM" || id == "EM":
			err = db.insertMachine(product)
		case id == "CP" || id == "EP":
			err = db.insertPod(product)
		}
		if err != nil {
			return err
		}
	}
	if err := sc.Err(); err != nil {
		return err
	}
	return nil
}

// insertMachine inserts a machine record to the db using the string
// provided
func (db *DB) insertMachine(m string) error {
	sku := string(m[:5])
	var water bool
	var sizeID int
	var modelID int
	if strings.Contains(m, "water line compatible") {
		water = true
	} else {
		water = false
	}
	sizeName := strings.Split(m, " ")[2]
	rows, err := db.Query("SELECT SizeID FROM Size WHERE SizeName = ?", sizeName)
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&sizeID)
	}
	modelName := strings.Trim(strings.Split(m, ",")[1], " ")
	rows, err = db.Query("SELECT ModelID FROM Model WHERE ModelName = ?", modelName)
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&modelID)
	}

	stmt, err := db.Prepare("INSERT INTO CoffeeMachine (SKU, ModelID, WaterLine, SizeID) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(sku, modelID, water, sizeID); err != nil {
		return err
	}
	return nil
}

// insertMachine inserts a machine record to the db using the string
// provided
func (db *DB) insertPod(p string) error {
	var sizeID int
	var flavorID int
	sku := string(p[:5])
	parts := strings.Split(p, ",")
	dozens, err := strconv.Atoi(string(strings.Trim(parts[1], " ")[0]))
	if err != nil {
		return err
	}
	flavorName := strings.Trim(parts[len(parts)-1], " ")
	rows, err := db.Query("SELECT FlavorID FROM Flavor WHERE FlavorName = ?", flavorName)
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&flavorID)
	}
	quantity := dozens * 12
	sizeName := strings.Split(p, " ")[2]
	rows, err = db.Query("SELECT SizeID FROM Size WHERE SizeName = ?", sizeName)
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&sizeID)
	}
	stmt, err := db.Prepare("INSERT INTO Pod (SKU, FlavorID, Quantity, SizeID) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(sku, flavorID, quantity, sizeID); err != nil {
		return err

	}
	return nil
}
